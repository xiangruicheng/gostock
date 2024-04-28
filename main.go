package main

import (
	"gostock/config"
	"gostock/data/datainit"
	"gostock/model"
	"gostock/server"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	datainit.InitKline(model.StockInfoModel_TypeEtf, "159915", "SZ", config.Data.Xueqiu.InitNum)

}
