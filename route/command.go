package route

import (
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/report"
	"gostock/util"
	"os"
	"reflect"
)

type Command struct {
	Key      string
	Func     any
	Descript string
}

var CommandConfig []*Command

func init() {
	CommandConfig = []*Command{
		{"make:db", ddl.Create, "Create DB and Create Table"},
		{"make:stock", datainit.InitStockInfo, "Init stock_info Table"},
		{"make:kline", datainit.BatchInitKline, "Init kline Table"},
		{"make:people", datainit.BatchUpdateStockPeople, "Init stock_people Table"},
		{"make:block", datainit.InitBlock, "Init stock_block Table"},

		{"report:real", report.Real, "Report real"},
		{"report:day", report.Day, "Report Day"},

		{"daily:kline", datainit.BatchIncrKline, "Update K-line data daily"},
		{"daily:macd", datainit.BatchUpdateMacd, "Update MACD data daily"},
		{"daily:quote", datainit.BatchUpdateStockQuote, "Update quote data daily"},

		{"-h", help, "help doc"},
	}
}

func help() {
	for _, command := range CommandConfig {
		util.PrintCommand(command.Key, command.Descript)
	}
}

func CommandInit() bool {
	params := os.Args
	if len(params) > 1 {
		key := params[1]
		var method any
		for _, command := range CommandConfig {
			if command.Key == key {
				method = command.Func
			}
		}
		if method != nil {
			m := reflect.ValueOf(method)
			m.Call(nil)
			return true
		}
	}
	return false
}
