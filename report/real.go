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
			currentColor = color.RedString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else if quote.Data.Quote.Chg < 0 {
			currentColor = color.GreenString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else {
			currentColor = fmt.Sprintf("%10.3f", quote.Data.Quote.Current)
		}
		fmt.Printf("%10s|%s|%10.3f|%10.3f|\n", tag, currentColor, quote.Data.Quote.Percent, quote.Data.Quote.Chg)
	}
}

func Day() {
	config := []string{
		"SH000001", "SZ399001",
	}
	fmt.Printf("%10s|%10s|%10s|%10s|%10s|\n", "TAG", "PRICE", "PERCENT", "CHG", "AMOUNT")
	fmt.Printf("%10s|%10s|%10s|%10s|%10s|\n", "----------", "----------", "----------", "----------", "----------")
	for _, symbol := range config {
		quote, _ := xueqiu.Quote(symbol)
		currentColor := ""
		if quote.Data.Quote.Chg > 0 {
			currentColor = color.RedString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else if quote.Data.Quote.Chg < 0 {
			currentColor = color.GreenString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else {
			currentColor = fmt.Sprintf("%10.3f", quote.Data.Quote.Current)
		}
		amount := quote.Data.Quote.Amount / 100000000
		fmt.Printf("%10s|%s|%10.3f|%10.3f|%10.2f|\n", symbol, currentColor, quote.Data.Quote.Percent, quote.Data.Quote.Chg, amount)
	}
}
