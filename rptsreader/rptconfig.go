package rptsreader

import (
	"strings"
	"time"
)

const (
	Market_Highlights = `<a name = "market_highlights">MARKET HIGHLIGHTS</a>,
<a name = "quotations">QUOTATIONS</a>,fsingle,Market_Highlights.csv,Blue_Chips_Market_Highlights.csv`

	Quotations = `<a name = "quotations">QUOTATIONS</a>,<a name = "sales_all">SALES RECORDS FOR ALL STOCKS</a>,
fdouble,Quotations.csv,Blue_Chips_Quotations.csv`

	Quotations_header1 = `CODE  NAME OF STOCK    CUR PRV.CLO./    ASK/    HIGH/      SHARES TRADED/`
	Quotations_header2 = `CLOSING      BID     LOW        TURNOVER ($)`

	Rpt_tail = `-------------------------------------------------------------------------------`

	Short_Selling_header1 = `Total Short Selling Turnover         Total Turnover`
	Short_Selling_header2 = `CODE  NAME OF STOCK       (SH)           ($)           (SH)                ($)`

	Short_Selling_end = `Total Shares (SH):`

	Sales_Records_All_Stocks = `<a name = "sales_all">SALES RECORDS FOR ALL STOCKS</a>,
<a name = "sales_over">SALES RECORDS OVER $500,000</a>,fmulti,Sales_Records_All_Stocks.csv,Blue_Chips_Sales_Records_All_Stocks.csv`

	Sales_Records_Over_500000 = `<a name = "sales_over">SALES RECORDS OVER $500,000</a>,
<a name = "amendments">AMENDMENT RECORDS FOR TRADE</a>,fmulti,Sales_Records_Over_500000.csv,Blue_Chips_Sales_Records_Over_500000.csv`

	Amendment_Records_for_Trad = `<a name = "amendments">AMENDMENT RECORDS FOR TRADE</a>,
<a name = "adj_turnover">ADJUSTED TURNOVER</a>,fsingle,Amendment_Records_for_Trad.csv,Blue_Chips_Amendment_Records_for_Trad.csv`

	Adjusted_Turnover = `<a name = "adj_turnover">ADJUSTED TURNOVER</a>,
<a name = "short_selling">SHORT SELLING TURNOVER - DAILY REPORT</a>,fsingle,Adjusted_Turnover.csv,Blue_Chips_Adjusted_Turnover.csv`

	Short_Selling_Turnover = `<a name = "short_selling">SHORT SELLING TURNOVER - DAILY REPORT</a>,
<a name = "adj_short">PREVIOUS DAY'S ADJUSTED SHORT SELLING TURNOVER</a>,fsingle,Short_Selling_Turnover.csv,Blue_Chips_Short_Selling_Turnover.csv`

	Previous_Day_Adjusted_Short_Selling_Turnover = `<a name = "adj_short">PREVIOUS DAY'S ADJUSTED SHORT SELLING TURNOVER</a>,
</font></pre></body></html>,fsingle,Previous_Day_Adjusted_Short_Selling_Turnover.csv,Blue_Chips_Previous_Day_Adjusted_Short_Selling_Turnover.csv`

	Trading_Suspended = `TRADING SUSPENDED`

	Trading_Halted = `TRADING HALTED`

	Blue_Chips = `1,2,3,4,5,6,11,12,16,17,19,23,27,66,83,101,135,144,151,267,293,322,386,388,494,688,700,762,823,
836,857,883,939,941,992,1038,1044,1088,1109,1113,1299,1398,1880,1928,2318,2319,2388,2628,3328,3988`

	Daily_statistic_directory = "/home/adley/HKEX_reports/Daily_Reports/Daily_Statistics_derivatives/"

	Daily_statistic_filename = `DailyStatistics_F1_HSI@hsi_futures,DailyStatistics_F1_VHS@hsi_vix_futures,
DailyStatistics_F1_CUS@rmb_futures,DailyStatistics_F1_MCH@mini_hsi_futures,
DailyStatistics_FnO@futures_n_options,DailyStatistics_F2_21@stock_futures,
DailyStatistics_O_HSI@hsi_options,DailyStatistics_O_MHI@mini_hsi_options,
DailyStatistics_O_22@stock_options`
)


func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func getrptctl(intype string) (outctl []string) {
	//var begend []string
	switch intype {
	case "Market_Highlights":
		outctl = strings.Split(Market_Highlights, ",")
		return outctl
	case "Quotations":
		outctl = strings.Split(Quotations, ",")
		return outctl
	case "Short_Selling_Turnover":
		outctl = strings.Split(Short_Selling_Turnover, ",")
		return outctl
	default:
		return outctl
	}
}

func Fmtstr(in string) string {
	return strings.TrimSpace(strings.Replace(in, "\n", "", -1))
}

func Getdate() string {
	t := time.Now()
	formatedTime := t.Format(time.RFC3339)
	in := strings.Split(formatedTime, "-")
	inday := strings.Split(in[2], "T")
	finald := in[0] + "-" + in[1] + "-" + inday[0]
	//fmt.Println(finald)
	return finald
}

func GetHKEXrpt_date_suffix(indate string) string {
	rptdate := strings.Split(indate, "-")
	rptyear := rptdate[0]
	rptyear = rptyear[2:4]
	datesuffix := rptyear + rptdate[1] + rptdate[2]
	return datesuffix
}

