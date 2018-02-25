package buildinfunc

import (
	"bytes"
	"regexp"
)

func Filterhtmlsymbols(ib []byte) string {
	ob := bytes.Replace(ib, []byte("&nbsp;"), []byte(""), -1)
	src := string(ob)
	re1, _ := regexp.Compile("\\<iframe[\\S\\s]+?\\</iframe\\>")
	src = re1.ReplaceAllString(src, "")
	re2, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re2.ReplaceAllString(src, "")
	re3, _ := regexp.Compile("\\<h2[\\S\\s]+?\\</h2\\>")
	src = re3.ReplaceAllString(src, "")
	/*re4, _ := regexp.Compile("\\<ul[\\S\\s]+?\\</ul\\>")
	src = re4.ReplaceAllString(src, "")
	re5, _ := regexp.Compile("\\<li[\\S\\s]+?\\</li\\>")
	src = re5.ReplaceAllString(src, "")*/
	return src
}
