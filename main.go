package main

import (
	"gostock/config"
	"gostock/route"
	"gostock/server"
	"gostock/strategy"
)

func main() {
	config.InitConfig()
	server.InitRedis()
	server.InitMysql()
	strategy.TradeDay.InitTradeDay()

	// exec command
	isCommand := route.CommandInit()
	if isCommand {
		return
	}

	// start http
	r := route.RouteInit()
	r.Run("127.0.0.1:9217")
}
