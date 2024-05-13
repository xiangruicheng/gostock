package report

import (
	"fmt"
	"gostock/data/datasource/xueqiu"
)

var RealPeportConfig = [][]string{
	{"gf", "SH515790"},
	{"cyb", "SZ159915"},
}

func Real() {
	for _, item := range RealPeportConfig {
		tag := item[0]
		symbol := item[1]
		quote, _ := xueqiu.Quote(symbol)
		fmt.Printf("\t%s\t|\t%f\t|\t%f\t|\t%f\t\n", tag, quote.Data.Quote.Current, quote.Data.Quote.Percent, quote.Data.Quote.Chg)
	}
}
