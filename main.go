package main

import (
	"fmt"
	"gostock/config"
	"gostock/data/datainit"
	"gostock/data/ddl"
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

	fmt.Printf("%s\t:\t%s\n", "report:real", "this is real report")
	fmt.Printf("%s\t:\t%s\n", "datainit:batchincr", "every day batch incr update kline")

}

func route(serverType string) {
	switch serverType {
	case "migrate:db":
		ddl.Create()
	case "migrate:stock":
		datainit.InitStockInfo()
	case "migrate:kline":
		datainit.BatchInitKline()

	case "report:real":
		report.Real()
	case "datainit:batchincr":
		datainit.BatchIncrKline()
	default:
		fmt.Println("go stock!!")
	}
}
