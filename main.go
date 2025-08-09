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

	var s strategy.StrategyInterface
	s = &strategy.Turtle{
		Code:       "002317",
		TotalMoney: 1000000.00,
		Risk:       0.01,
	}
	s.Help()
	s.Run()

}
