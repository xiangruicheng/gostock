package main

import (
	"gostock/config"
	"gostock/server"
	"gostock/strategy"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	strategy.Cyb()
	//datainit.InitKline(model.StockInfoModel_TypeEtf, "515790", "SH", config.Data.Xueqiu.InitNum)

}
