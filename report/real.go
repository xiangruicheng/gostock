package report

import (
	"fmt"
	"github.com/fatih/color"
	"gostock/data/datasource/xueqiu"
)

var RealPeportConfig = [][]string{
	{"GF", "SH515790"},
	{"CYB", "SZ159915"},
	{"JG", "SH512660"},
}

func Real() {
	fmt.Printf("%10s|%10s|%10s|%10s|\n", "TAG", "PRICE", "PERCENT", "CHG")
	fmt.Printf("%10s|%10s|%10s|%10s|\n", "----------", "----------", "----------", "----------")
	for _, item := range RealPeportConfig {
		tag := item[0]
		symbol := item[1]
		quote, _ := xueqiu.Quote(symbol)
		currentColor := ""
		if quote.Data.Quote.Chg > 0 {
			currentColor = color.RedString(fmt.Sprintf("%10f", quote.Data.Quote.Current))
		} else if quote.Data.Quote.Chg < 0 {
			currentColor = color.GreenString(fmt.Sprintf("%10f", quote.Data.Quote.Current))
		} else {
			currentColor = fmt.Sprintf("%10f", quote.Data.Quote.Current)
		}
		fmt.Printf("%10s|%s|%10f|%10f|\n", tag, currentColor, quote.Data.Quote.Percent, quote.Data.Quote.Chg)
	}
}
