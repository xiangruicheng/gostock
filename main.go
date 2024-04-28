package main

import (
	"gostock/config"
	"gostock/data/ddl"
	"gostock/server"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	ddl.Create()

	//server.Log(server.LogLevelDebug, "this is error")

	//datainit.BatchInitKline()
	//datainit.InitStockInfo()
	//datainit.InitKline("515790", "SH", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("159915", "SZ", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("000001", "SH", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("399001", "SZ", config.Data.Xueqiu.InitNum)
	//datainit.InitKline("399006", "SZ", config.Data.Xueqiu.InitNum)

}
