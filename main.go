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

	datainit.BatchInitKline()
	//datainit.InitStockInfo()

}
