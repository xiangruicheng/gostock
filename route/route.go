package route

import (
	"github.com/gin-gonic/gin"
	"gostock/controller"
)

func RouteInit() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/kline", controller.Kline)
	return r
}
