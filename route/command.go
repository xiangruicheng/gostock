package route

import (
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/report"
	"gostock/util"
	"os"
	"reflect"
)

type CommandStruct struct {
	Key      string
	Func     any
	Descript string
}

var Commands []*CommandStruct

func init() {
	Commands = []*CommandStruct{
		{"-h", help, "help doc"},
		{"start", startGin, "start gin server"},

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
	}
}

func help() {
	for _, command := range Commands {
		util.PrintCommand(command.Key, command.Descript)
	}
}

func startGin() {
	// start http
	r := RouteInit()
	r.Run("127.0.0.1:9217")
}

func CommandInit() bool {
	params := os.Args
	if len(params) > 1 {
		key := params[1]
		var method any
		for _, command := range Commands {
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
