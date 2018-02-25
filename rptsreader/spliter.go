package rptsreader

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Split_file() {
	f, err := os.Open("")
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var part1, part2, part3 bool
	//var part1 bool
	part1 = true
	part2 = false
	part3 = false
	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		//part 1
		if strings.TrimSpace(line) != `<a name = "sales_all">SALES RECORDS FOR ALL STOCKS</a>` && part1 == true {
			//fmt.Print(intln(line)
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
		if strings.TrimSpace(line) == `<a name = "sales_all">SALES RECORDS FOR ALL STOCKS</a>` && part1 == true {
			outf, err := os.Create("testbufio1.txt")
			Check(err)
			defer outf.Close()
			w1 := bufio.NewWriter(outf)
			n1, err := w1.WriteString(buffer.String())
			fmt.Printf("wrote %d bytes\n", n1)
			w1.Flush()
			part1 = false
			part2 = true
			buffer.Reset()

		}
		//part2
		if strings.TrimSpace(line) != `<a name = "amendments">AMENDMENT RECORDS FOR TRADE</a>` && part2 == true {
			//fmt.Print(intln(line)
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
		if strings.TrimSpace(line) == `<a name = "amendments">AMENDMENT RECORDS FOR TRADE</a>` && part2 == true {
			outf2, err := os.Create("testbufio2.txt")
			Check(err)
			defer outf2.Close()
			w2 := bufio.NewWriter(outf2)
			n2, err := w2.WriteString(buffer.String())
			fmt.Printf("wrote %d bytes\n", n2)
			w2.Flush()
			part2 = false
			part3 = true
			buffer.Reset()
		}

		//part3
		if part1 == false && part2 == false && part3 == true {
			//fmt.Print(intln(line)
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
	}
	outf3, err := os.Create("testbufio3.txt")
	Check(err)
	defer outf3.Close()
	w3 := bufio.NewWriter(outf3)
	n3, err := w3.WriteString(buffer.String())
	fmt.Printf("wrote %d bytes\n", n3)
	w3.Flush()
	part3 = false
	//buffer.Reset()
}
