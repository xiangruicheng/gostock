package route

import (
	"github.com/gin-gonic/gin"
	"gostock/controller"
)

func RouteInit() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/kline", controller.Kline)
	r.POST("/user/init", controller.UserController.InitUser)
	r.POST("/user/hold", controller.UserController.Hold)

	r.POST("/order/buy", controller.OrderController.Buy)
	r.POST("/order/sell", controller.OrderController.Sell)

	return r
}

func GinStart() {
	r := RouteInit()
	r.Run("127.0.0.1:9217")
}
