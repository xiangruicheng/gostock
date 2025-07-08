package main

import (
	"gostock/config"
	"gostock/route"
	"gostock/server"
	"gostock/strategy"
)

func main() {
	// 现金流折现计算
	//strategy.Dcf.Demo()
	//return

	config.InitConfig()
	server.InitRedis()
	server.InitMysql()
	strategy.TradeDay.InitTradeDay()

	// exec command
	if route.CommandInit() {
		return
	}

	//new(report.PeopleReport).Run()

	new(strategy.Threeban).Run()

}
