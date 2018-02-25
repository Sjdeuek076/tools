package buildinfunc

import "bytes"

func Formatdate(indate string, informat string, insplit string) string {
	//fmt.Println(string(informat[0]))
	var day, mon, year, rval bytes.Buffer
	var leng, start int
	leng = 0
	start = 0
	for i := 0; i < 3; i++ {
		switch string(informat[i]) {
		case "D":
			leng = 2 + start - 1
			day.WriteString(indate[start : leng+1])
			//fmt.Println(leng, start, indate[start:leng+1])
			start = leng + 2
		case "M":
			leng = 2 + start - 1
			mon.WriteString(indate[start : leng+1])
			//fmt.Println(leng, start, indate[start:leng+1])
			start = leng + 2
		case "Y":
			leng = 4 + start - 1
			year.WriteString(indate[start : leng+1])
			//fmt.Println(leng, start, indate[start:leng+1])
			start = leng + 2
		}
	}
	/*fmt.Println(day.String())
	fmt.Println(mon.String())
	fmt.Println(year.String())*/
	rval.WriteString(year.String())
	rval.WriteString(insplit)
	rval.WriteString(mon.String())
	rval.WriteString(insplit)
	rval.WriteString(day.String())
	return (rval.String())
	//fmt.Println(rval.String())
}
