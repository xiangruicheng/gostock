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

	strategy.TradeDay.InitTradeDay()

	// exec command
	isCommand := CommandInit()
	if isCommand {
		return
	}

	//strategy.Feature.VolumeRateRange("000001", "20240531", 0, 10)
	strategy.Test()

	//new(report.PeopleReport).Run()

}

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
	"-h":          server.Help,
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
