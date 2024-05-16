package main

import (
	"fmt"
	"gostock/config"
	"gostock/data/datainit"
	"gostock/data/ddl"
	"gostock/report"
	"gostock/server"
	"gostock/strategy"
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

	strategy.Cyb()
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
	default:
		fmt.Println("go stock")
	}
}
