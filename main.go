package main

import (
	"gostock/config"
	"gostock/server"
	datainit "gostock/service"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()

	//datainit.Kline("159915")

	datainit.StockInfo()

}
