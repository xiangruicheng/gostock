package main

import (
	"gostock/config"
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/report"
	"gostock/server"
	"gostock/strategy"
	"os"
	"reflect"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	// exec command
	isCommand := CommandInit()
	if isCommand {
		return
	}

	//new(report.PeopleReport).Run()
	new(strategy.TailStrategy).RunLine()

}

var commandConfig = map[string]any{
	"migrate:db":     ddl.Create,
	"migrate:stock":  datainit.InitStockInfo,
	"migrate:kline":  datainit.BatchInitKline,
	"migrate:people": datainit.BatchUpdateStockPeople,

	"report:real": report.Real,
	"report:day":  report.Day,

	"daily:kline": datainit.BatchIncrKline,
	"daily:macd":  datainit.BatchUpdateMacd,
	"daily:quote": datainit.BatchUpdateStockQuote,
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
