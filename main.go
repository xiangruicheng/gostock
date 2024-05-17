package main

import (
	"fmt"
	"gostock/config"
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/data/indicator"
	"gostock/report"
	"gostock/server"
	"os"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	params := os.Args
	if len(params) > 1 {
		serverType := params[1]
		route(serverType)
		return
	}

}

func route(serverType string) {
	switch serverType {

	case "migrate:db":
		ddl.Create()
	case "migrate:stock":
		datainit.InitStockInfo()
	case "migrate:kline":
		datainit.BatchInitKline()

		// report
	case "report:real":
		report.Real()
	case "report:day":
		report.Day()

		//daily task
	case "daily:kline":
		datainit.BatchIncrKline()
	case "daily:macd":
		indicator.MacdBatchRun()
	default:
		fmt.Println("go stock")
	}
}
