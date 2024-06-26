package report

import (
	"fmt"
	"github.com/fatih/color"
	"gostock/data/datasource/eastmoney"
	"gostock/data/datasource/xueqiu"
	"strconv"
)

var RealPeportConfig = [][]string{
	//{"hsgy", "SH603260", "", ""},
	//{"htwy", "SH603288", "", ""},
	//{"ylgf", "SH600887", "", ""},
	{"300", "SH510330", "60000", "3.576"},
	{"BDT", "SH512480", "", ""},
	{"ZQ", "SH512880", "", ""},
	{"GF", "SH515790", "", ""},
	{"CYB", "SZ159915", "", ""},
	{"JG", "SH512660", "", ""},
}

// this is real
func Real() {
	fmt.Printf("%10s|%10s|%10s|%10s|%10s|%10s|%10s|\n", "TAG", "PRICE", "PERCENT", "High", "LOW", "COST", " P&L")
	fmt.Printf("%10s|%10s|%10s|%10s|%10s|%10s|%10s|\n", "----------", "----------", "----------", "----------", "----------", "----------", "----------")
	for _, item := range RealPeportConfig {
		tag := item[0]
		symbol := item[1]
		cost, _ := strconv.ParseFloat(item[3], 64)
		num, _ := strconv.ParseFloat(item[2], 64)

		quote, _ := xueqiu.Quote(symbol)
		pl := (quote.Data.Quote.Current - cost) * num

		currentColor := ""
		if quote.Data.Quote.Chg > 0 {
			currentColor = color.RedString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else if quote.Data.Quote.Chg < 0 {
			currentColor = color.GreenString(fmt.Sprintf("%10.3f", quote.Data.Quote.Current))
		} else {
			currentColor = fmt.Sprintf("%10.3f", quote.Data.Quote.Current)
		}
		fmt.Printf("%10s|%s|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|\n", tag, currentColor, quote.Data.Quote.Percent, quote.Data.Quote.High, quote.Data.Quote.Low, cost, pl)
	}
}

func Day() {
	config := []string{
		"SH1A0001", "SZ399001",
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

	//zdfb
	resp, _ := eastmoney.Zdfb()
	var z, f int
	for _, item := range resp.Data.Fenbu {
		for i := -20; i < 20; i++ {
			if v, ok := item[i]; ok {
				if i >= 0 {
					z += v
				} else {
					f += v
				}
			}
		}
	}
	fmt.Printf("UP:%d\nDOWN:%d\n", z, f)
}
