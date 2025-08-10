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
	server.InitMysqlEngine()
	strategy.TradeDay.InitTradeDay()

	// exec command
	if route.CommandInit() {
		return
	}
	/*	datainit.InitKline("512480", "SH", config.Data.Xueqiu.InitNum)
		return*/
	//new(report.PeopleReport).Run()

}
