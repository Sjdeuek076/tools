package rptsreader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FormatCSV(inline string) string {
	var outline string
	re0, _ := regexp.Compile(`\,`)
	outline = re0.ReplaceAllString(inline, "")
	re1, _ := regexp.Compile(`</font></pre><pre><font size='1'>`)
	outline = re1.ReplaceAllString(outline, "")
	re2, _ := regexp.Compile(`\*|\%`)
	outline = re2.ReplaceAllString(outline, "")

	return outline
}

func custom_split(input string, rectype string) (string, bool) {
	//const input = "1234 5678 1234567901234567890"
	var linend bool
	linend = false
	scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner := bufio.NewScanner(f)
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF) ///bufio.ScanRunes(data, atEOF)
		if err == nil && token != nil {
			//_, err = strconv.ParseInt(string(token), 10, 32)
			if advance >= 36 && rectype == "fdouble" {
				//fmt.Println(advance, string(token))
				linend = true
			} else if rectype == "fsingle" {
				linend = true
			}
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	var csvline string
	var i int
	i = 0
	var istart bool
	istart = false
	for scanner.Scan() {
		//fmt.Printf("%s\n", scanner.Text())
		if rectype == "fdouble" && linend == false {
			if i == 0 {
				csvline += scanner.Text() + ","
				i++
			}
			if istart {
				csvline += scanner.Text() + ","
			}
			switch scanner.Text() {
			case "HKD", "CNY", "USD", "GBP", "CAD", "JPY", "SGD", "AUD", "EUR":
				istart = true
			}
		} else if rectype == "fdouble" && linend == true {
			csvline += scanner.Text() + ","
		} else if rectype == "fsingle" {
			//_, erri := strconv.ParseInt(scanner.Text(), 12, 64)
			if _, err := strconv.Atoi(scanner.Text()); err == nil {
				csvline += scanner.Text() + ","
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	csvline = fmtcsvline(csvline)
	if rectype == "fsingle" {
		outline := fmtcsvline_shortsell(csvline)
		return outline, linend
	}
	return csvline, linend
}

func fmtcsvline(csvline string) string {
	re0, _ := regexp.Compile(`\-`)
	csvline = re0.ReplaceAllString(csvline, "0")
	re1, _ := regexp.Compile(`N\/A`)
	csvline = re1.ReplaceAllString(csvline, "0")
	re2, _ := regexp.Compile(`\#[a-zA-Z0-9]{1,}`)
	csvline = re2.ReplaceAllString(csvline, "")
	re3, _ := regexp.Compile(`HKD\,`)
	csvline = re3.ReplaceAllString(csvline, "")
	re4, _ := regexp.Compile(`([a-zA-Z]{1,}[0-9]{1,}\,)|([a-zA-Z]{1,}\,)`)
	csvline = re4.ReplaceAllString(csvline, "")
	/*re5, _ := regexp.Compile(`0{78}\,*`)
	csvline = re5.ReplaceAllString(csvline, "")*/

	//fmt.Println(csvline)
	return csvline
}

func fmtcsvline_shortsell(csvline string) string {
	var outline string
	linesplit := strings.Split(csvline, ",")
	//fmt.Println(len(linesplit), "~~", csvline)
	i := len(linesplit)
	if i >= 5 {
		outline = linesplit[0] + "," + linesplit[i-5] + "," + linesplit[i-4] + "," + linesplit[i-3] + "," + linesplit[i-2] + ","
		//fmt.Println(outline)
	}
	return outline
}

func GetContent(infile io.Reader, inrpttag string, tocsv bool, indate string) { //rectype: gs = general stock; ss = short sell record
	var buffer bytes.Buffer
	//var bluechips bytes.Buffer
	var token_go, linend bool //, secondrow bool
	var rpttag []string
	rpttag = getrptctl(inrpttag)
	token_go = false
	//secondrow = false
	inscanner := bufio.NewScanner(infile)
	for inscanner.Scan() {
		line := inscanner.Text()
		re3, _ := regexp.Compile(Trading_Suspended) // ignored the Trading Suspended records
		if re3.MatchString(line) {
			continue
		}
		re4, _ := regexp.Compile(Trading_Halted) // ignored the Trading Halted records
		if re4.MatchString(line) {
			continue
		}
		//if strings.TrimSpace(line) == Fmtstr(rpttag[0]) { //inbegin
		if strings.Contains(strings.TrimSpace(line), Fmtstr(rpttag[0])) {
			token_go = true
			continue
		}
		//if strings.TrimSpace(line) == Fmtstr(rpttag[1]) { //inend
		if strings.Contains(strings.TrimSpace(line), Fmtstr(rpttag[1])) {
			token_go = false
			break
		}
		if token_go == true {
			// all stock to csv file
			tline := strings.TrimSpace(line)
			if tline == Quotations_header1 || tline == Quotations_header2 { //remove Quotations header
				continue
			}

			if tline == Short_Selling_header1 || tline == Short_Selling_header2 { // remove Short selling --daily rpt header
				continue
			}

			if strings.Contains(tline, Rpt_tail) || strings.Contains(tline, Short_Selling_end) { //remove tail and stop scanning
				break
				token_go = false
			}

			if tocsv {
				line = FormatCSV(line)
				line, linend = custom_split(line, Fmtstr(rpttag[2]))
				buffer.WriteString(line)
				if linend {
					buffer.WriteString(indate)
					buffer.WriteString("\n")
				}
				// on screen (for testing only)
			} else {
				buffer.WriteString(line)
				buffer.WriteString(indate)
				buffer.WriteString("\n")
			}
			// blue chips to csv
			/*
				if CheckBlueChips(line) {
					bluechips.WriteString(line)
					secondrow = true
				}
				if linend && secondrow {
					if Fmtstr(rpttag[2]) == "fdouble" {
						bluechips.WriteString(line)
					}
					bluechips.WriteString(indate)
					bluechips.WriteString("\n")
					secondrow = false
				}
			*/
		}
	}
	if tocsv {
		outf, err := os.Create(Fmtstr(rpttag[3]))
		Check(err)
		defer outf.Close()
		w := bufio.NewWriter(outf)
		n, err := w.WriteString(buffer.String())
		fmt.Printf("wrote %d bytes\n", n)
		w.Flush()

		/*
			// blue chip files
			bcf, err := os.Create(Fmtstr(rpttag[4]))
			Check(err)
			defer bcf.Close()
			bcw := bufio.NewWriter(bcf)
			bcn, err := bcw.WriteString(bluechips.String())
			fmt.Printf("wrote %d bytes\n", bcn)
			bcw.Flush()
		*/

	} else {
		fmt.Println(buffer.String())
	}
}

func CheckBlueChips(inline string) bool {
	var isbluechip bool
	isbluechip = false
	bluechips := strings.Split(Blue_Chips, ",")
	incode := strings.Split(inline, ",")
	for i := 0; i < 50; i++ {
		if bluechips[i] == incode[0] {
			isbluechip = true
		}
	}
	return isbluechip
}

//1. With bufio.Scanner
/*func WithScanner(input io.ReadSeeker, start int64) error {
	fmt.Println("--SCANNER, start:", start)
	if _, err := input.Seek(start, 0); err != nil { //input.Seek(offset, whence)
		return err
	}
	scanner := bufio.NewScanner(input)
	pos := start
	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		pos += int64(advance)
		return
	}
	scanner.Split(scanLines)
	for scanner.Scan() {
		fmt.Printf("Pos: %d, Scanned: %s\n", pos, scanner.Text())
	}
	return scanner.Err()
}
[a-zA-Z]\.[a-zA-Z]*
//2. With bufio.Reader
func WithReader(input io.ReadSeeker, start int64) error {
	fmt.Println("--READER, start:", start)
	if _, err := input.Seek(start, 0); err != nil {
		return err
	}
	r := bufio.NewReader(input)
	pos := start
	for {
		data, err := r.ReadBytes('\n')
		pos += int64(len(data))
		if err == nil || err == io.EOF {
			if len(data) > 0 && data[len(data)-1] == '\n' {
				data = data[:len(data)-1]
			}
			if len(data) > 0 && data[len(data)-1] == '\r' {
				data = data[:len(data)-1]
			}
			fmt.Printf("Pos: %d, Read: %s\n", pos, data)
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

func Tester() {
	const content = "first\r\nsecond\nthird\nfourth"
	if err := WithScanner(strings.NewReader(content), 0); err != nil {
		fmt.Println("Scanner error:", err)
	}
	if err := WithReader(strings.NewReader(content), 0); err != nil {
		fmt.Println("Reader error:", err)
	}
	if err := WithScanner(strings.NewReader(content), 14); err != nil {
		fmt.Println("Scanner error:", err)
	}
	if err := WithReader(strings.NewReader(content), 14); err != nil {
		fmt.Println("Reader error:", err)
	}
}*/
