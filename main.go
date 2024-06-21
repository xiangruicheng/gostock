package main

import (
	"gostock/config"
	"gostock/route"
	"gostock/server"
	"gostock/strategy"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()
	strategy.TradeDay.InitTradeDay()

	// exec command
	if route.CommandInit() {
		return
	}

	//new(report.PeopleReport).Run()

}
