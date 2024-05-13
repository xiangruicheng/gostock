package main

import (
	"gostock/config"
	"gostock/server"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	//datainit.BatchIncrKline()
	//strategy.Cyb()

	//report.Real()

	//datainit.InitKline(model.StockInfoModel_TypeEtf, "515790", "SH", config.Data.Xueqiu.InitNum)

}
