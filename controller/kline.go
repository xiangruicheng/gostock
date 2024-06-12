package controller

import (
	"github.com/gin-gonic/gin"
	"gostock/model"
)

func Kline(ctx *gin.Context) {
	code := ctx.Query("code")
	klines, _ := new(model.KlineModel).GetByCode(code)
	ctx.JSON(200, klines)
}
