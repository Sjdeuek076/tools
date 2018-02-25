package rptsreader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func Hsiwalk(n *html.Node) {
	buffer := bytes.NewBuffer(nil)
	if n.Type == html.ElementNode && n.Data == "table" {
		err := html.Render(io.Writer(buffer), n)
		if err == nil {
			//tester(buffer)
			//HsiGetContent(buffer)
			ntr, _ := html.Parse(strings.NewReader(buffer.String()))
			Hsitrwalk(ntr)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Hsiwalk(c)
	}
	//tester(buffer)
}

/*func tester(inbuffer *bytes.Buffer) { //rectype: gs = general stock; ss = short sell record
	f, _ := os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString(inbuffer.String())
	f.WriteString("Hello called WriteString\n")
	defer f.Close()
}*/

func Hsitrwalk(n *html.Node) {
	buffer := bytes.NewBuffer(nil)
	if n.Type == html.ElementNode && n.Data == "tr" {
		err := html.Render(io.Writer(buffer), n)
		if err == nil {
			hsigetcontent(buffer)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Hsitrwalk(c)
	}
}

func hsigetcontent(inbuffer *bytes.Buffer) {
	var buffer bytes.Buffer
	inscanner := bufio.NewScanner(inbuffer)
	for inscanner.Scan() {
		line := inscanner.Text()
		fmt.Println(len(line))
		if len(line) != 0 {
			outline := removenode(line)
			buffer.WriteString(outline)
			//buffer.WriteString(line)
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("\n")
	//fmt.Println(buffer.String())
	f, _ := os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, 0600)
	n1, _ := f.WriteString(buffer.String())
	fmt.Printf(" Wrote %d bytes\n", n1)
	defer f.Close()
}

func removenode(inline string) string {
	var outline string
	//this pattern mathces any html tag
	re1, _ := regexp.Compile(`<[^>]*>`)
	outline = re1.ReplaceAllString(inline, "")
	return outline
}
