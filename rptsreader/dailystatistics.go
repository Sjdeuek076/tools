package rptsreader

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func F1_hsi(n *html.Node, intbl string) {
	buffer := bytes.NewBuffer(nil)
	if n.Type == html.ElementNode && n.Data == "td" {
		err := html.Render(io.Writer(buffer), n)
		if err == nil {
			//tester(buffer)
			//HsiGetContent(buffer)
			if buffer.Len() > 100 {
				ntr, _ := html.Parse(strings.NewReader(buffer.String()))
				hsi(ntr, intbl)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		F1_hsi(c, intbl)
	}
}

func hsi(n *html.Node, intbl string) {
	var token_go bool
	buffer := bytes.NewBuffer(nil)
	token_go = false
	if n.Type == html.ElementNode && n.Data == "tr" {
		for _, a := range n.Attr {
			if a.Key != "BGCOLOR" && a.Val != "#AAB6CA" {
				token_go = true
				break
			}
		}
		if token_go {
			err := html.Render(io.Writer(buffer), n)
			if err == nil {
				hget(buffer, intbl)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hsi(c, intbl)
	}
}

func hget(inbuffer *bytes.Buffer, intbl string) {
	//var buffer bytes.Buffer
	var outline string
	inscanner := bufio.NewScanner(inbuffer)
	for inscanner.Scan() {
		line := inscanner.Text()
		if len(line) != 0 && strings.Contains(line, `<td`) {
			//fmt.Println(line)
			outline += remove(line) + ","
			//buffer.WriteString(outline)
			//buffer.WriteString(line)
			//buffer.WriteString(",")
		}
	}
	outline = rmlastcomma(outline)
	//fmt.Println("hello", intbl)
	In2db(outline, intbl)
	/*buffer.WriteString(outline)
	buffer.WriteString("\n")
	//fmt.Println(buffer.String())
	f, _ := os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, 0600)
	n1, _ := f.WriteString(buffer.String())
	fmt.Printf(" Wrote %d bytes\n", n1)
	defer f.Close()*/
}

func remove(inline string) string {
	var outline string
	//this pattern mathces any html tag
	re0, _ := regexp.Compile(`\,`)
	outline = re0.ReplaceAllString(inline, "")
	re1, _ := regexp.Compile(`<[^>]*>`)
	outline = re1.ReplaceAllString(outline, "")
	re2, _ := regexp.Compile(`\t`) // deal with tabs after remove all html nodes
	outline = re2.ReplaceAllString(outline, "")
	re3, _ := regexp.Compile(`\s`) // deal with date
	outline = re3.ReplaceAllString(outline, "-")
	//fmt.Println(outline)
	return outline
}

func rmlastcomma(ins string) string {
	ins = strings.TrimSuffix(ins, ",")
	//fmt.Println(ins)
	return ins
}
