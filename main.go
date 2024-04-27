package main

import (
	"gostock/config"
	"gostock/data/datainit"
	"gostock/server"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	datainit.BatchInitKline()
	//datainit.InitStockInfo()
	//datainit.InitKline("515790", "SH", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("159915", "SZ", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("000001", "SH", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("399001", "SZ", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("399006", "SZ", config.Data.Xueqiu.InitNum)

}
