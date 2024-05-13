package report

import (
	"fmt"
	"github.com/fatih/color"
	"gostock/data/datasource/xueqiu"
)

var RealPeportConfig = [][]string{
	{"gf", "SH515790"},
	{"cyb", "SZ159915"},
}

func Real() {
	fmt.Printf("%10s|%10s|%10s|%10s|\n", "TAG", "PRICE", "PERCENT", "CHG")
	fmt.Printf("%10s|%10s|%10s|%10s|\n", "----------", "----------", "----------", "----------")
	for _, item := range RealPeportConfig {
		tag := item[0]
		symbol := item[1]
		quote, _ := xueqiu.Quote(symbol)
		if quote.Data.Quote.Chg > 0 {
			color.Red("%10s|%10f|%10f|%10f|\n", tag, quote.Data.Quote.Current, quote.Data.Quote.Percent, quote.Data.Quote.Chg)
		} else {
			color.Green("%10s|%10f|%10f|%10f|\n", tag, quote.Data.Quote.Current, quote.Data.Quote.Percent, quote.Data.Quote.Chg)
		}
	}
}
