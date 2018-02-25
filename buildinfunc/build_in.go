package buildinfunc

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type INdex struct {
	XMLName      xml.Name `xml:"index"`
	NO           string   `xml:"no,attr"`
	Code         string   `xml:"code,attr"`
	Sector       string   `xml:"sector,attr"`
	Status       string   `xml:"status,attr"`
	Datetime     string   `xml:"datetime,attr"`
	Current      string   `xml:"current,attr"`
	High         string   `xml:"high,attr"`
	Low          string   `xml:"low,attr"`
	Change       string   `xml:"Change,attr"`
	Precent      string   `xml:"precent,attr"`
	Name         string   `xml:"name"`
	CName        string   `xml:"cname"`
	Constituents string   `xml:"constituents"`
}

type AllIndex struct {
	XMLName xml.Name `xml:"allindex"`
	INdexes []INdex  `xml:"index"`
}

type ConfigItem struct {
	XMLName           xml.Name `xml:"configitem"`
	Title             string   `xml:"title"`
	Curl              string   `xml:"curl"`
	Filternode        string   `xml:"filternode"`
	Filterkey         string   `xml:"filterkey"`
	Filterval         string   `xml:"filterval"`
	Fcharset          string   `xml:"fcharset"`
	Filterbylengthup  int      `xml:"filterbylengthup"`
	Filterbylengthlow int      `xml:"filterbylengthlow"`
	Corder            int      `xml:"corder"`
}

type SysConfigItem struct {
	XMLName          xml.Name `xml:"sysconfigitem"`
	FileOutPath      string   `xml:"fileoutpath"`
	FileBkPath       string   `xml:"filebkpath"`
	FileNamePrex     string   `xml:"filenameprex"`
	FileShortPutPrex string   `xml:"fileshortputprex"`
}

type OwnConfig struct {
	XMLName        xml.Name        `xml:"ownconfig"`
	ConfigItems    []ConfigItem    `xml:"configitem"`
	SysConfigItems []SysConfigItem `xml:"sysconfigitem"`
}

type INdex_sub2 struct {
	XMLName      xml.Name `xml:"index"`
	NO           string   `xml:"no,attr"`
	Code         string   `xml:"code,attr"`
	Sector       string   `xml:"sector,attr"`
	Status       string   `xml:"status,attr"`
	Datetime     string   `xml:"datetime,attr"`
	Current      string   `xml:"current,attr"`
	High         string   `xml:"high,attr"`
	Low          string   `xml:"low,attr"`
	Change       string   `xml:"change,attr"`
	Precent      string   `xml:"precent,attr"`
	Name         string   `xml:"name"`
	CName        string   `xml:"cname"`
	Constituents string   `xml:"constituents"`
}

type Subindex_sub struct {
	XMLName      xml.Name     `xml:"sub-indexes"`
	INdex_sub2es []INdex_sub2 `xml:"index"`
}

type INdex_sub struct {
	XMLName        xml.Name       `xml:"index"`
	NO             string         `xml:"no,attr"`
	Code           string         `xml:"code,attr"`
	Sector         string         `xml:"sector,attr"`
	Status         string         `xml:"status,attr"`
	Datetime       string         `xml:"datetime,attr"`
	Current        string         `xml:"current,attr"`
	High           string         `xml:"high,attr"`
	Low            string         `xml:"low,attr"`
	Change         string         `xml:"change,attr"`
	Precent        string         `xml:"precent,attr"`
	Name           string         `xml:"name"`
	CName          string         `xml:"cname"`
	Constituents   string         `xml:"constituents"`
	Subindexes_sub []Subindex_sub `xml:"sub-indexes"`
}

type AllIndex_sub struct {
	float64
	XMLName     xml.Name    `xml:"allindex"`
	INdexes_sub []INdex_sub `xml:"index"`
}

type ShortPut struct {
	StockNo    string
	Company    string
	Qty        string
	Amount     string
	Unit_price string
}

func GetExchangeplugin() string {
	const Explugin = `<!DOCTYPE html>
	<html>
	<head>
	<link rel="stylesheet" href="../css/final_out.css" type="text/css" />
	<meta charset="utf-8"/>
	</head>
	<body>
	<!-- CURRENCY.ME.UK CURRENCY RATES TABLE START -->
<div style="width:178px;margin:0;padding:0;border:1px solid #2D6AB4;background:#F0F0F0;">
<div style="width:178px;text-align:center;padding:2px 0px;background:#2D6AB4;font-family:arial;font-size:11px;color:#FFFFFF;font-weight:bold;vertical-align:middle;">
<img src="http://www.exchangerates.org.uk/images/flags/us.gif" style="padding-right:5px;">
<a style="color:#FFFFFF;text-decoration:none;text-transform:uppercase;" href="http://www.currency.me.uk/rates/usd-us-dollar" target="_new" title="US Dollar exchange rate">US Dollar exchange rate</a>
</div>
<div style="padding:5px;text-align:center;">
<script type="text/javascript">
var nb = '10';
var mc = 'USD';
var tz = '8';
var c1 = 'USD';
var c2 = 'EUR';
var c3 = 'AUD';
var c4 = 'JPY';
var c5 = 'INR';
var c6 = 'CAD';
var c7 = 'ZAR';
var c8 = 'SGD';
var c9 = 'HKD';
var c10 = 'CNY';
var mcol = '2D6AB4';
var mbg = 'F0F0F0';
var f = 'arial';
var fc = '000000';
var tc = 'FFFFFF';

</script>
<script type="text/javascript" src="http://www.currency.me.uk/remote/CUK-TABLE2-1.php"></script>
</div></div>
<!-- CURRENCY.ME.UK CURRENCY RATES TABLE END -->`


	var buffer bytes.Buffer
	buffer.WriteString(Explugin)
	//buffer.WriteString(BDIplugin)

	return buffer.String()
}

func Getfilename(foption int, inconfig string) (genfilename string) { //foption 0 is filename, 1 is bk_filename
	xmlFile, err := os.Open(inconfig)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	XMLdata, _ := ioutil.ReadAll(xmlFile)
	var rc OwnConfig
	xml.Unmarshal(XMLdata, &rc)
	switch foption {
	case 0:
		genfilename := rc.SysConfigItems[0].FileOutPath + rc.SysConfigItems[0].FileNamePrex + ".htm"
		return genfilename
	case 1:
		genfilename := rc.SysConfigItems[0].FileBkPath + rc.SysConfigItems[0].FileNamePrex + strconv.FormatInt(time.Now().Unix(), 10) + ".htm"
		return genfilename
	}
	return genfilename
}


