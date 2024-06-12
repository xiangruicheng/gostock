package route

import (
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/report"
	"gostock/util"
	"os"
	"reflect"
)

var commandConfig = map[string]any{
	"make:db":     ddl.Create,
	"make:stock":  datainit.InitStockInfo,
	"make:kline":  datainit.BatchInitKline,
	"make:people": datainit.BatchUpdateStockPeople,
	"make:block":  datainit.InitBlock,

	"report:real": report.Real,
	"report:day":  report.Day,

	"daily:kline": datainit.BatchIncrKline,
	"daily:macd":  datainit.BatchUpdateMacd,
	"daily:quote": datainit.BatchUpdateStockQuote,
	"-h":          help,
}

func help() {
	util.PrintCommand("make:db", "Create DB and Create Table")
	util.PrintCommand("make:stock", "Init stock_info Table")
	util.PrintCommand("make:kline", "Init kline Table")
	util.PrintCommand("make:people", "Init stock_people Table")
	util.PrintCommand("make:block", "Init stock_block&stock_block_code Table")
}

func CommandInit() bool {
	params := os.Args
	if len(params) > 1 {
		serverType := params[1]
		method := commandConfig[serverType]
		if method != nil {
			m := reflect.ValueOf(method)
			m.Call(nil)
			return true
		}
	}
	return false
}
