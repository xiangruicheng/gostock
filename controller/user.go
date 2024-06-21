package controller

import (
	"github.com/gin-gonic/gin"
	"gostock/model"
	"strconv"
	"time"
)

type UserControllerStruct struct {
}

var UserController *UserControllerStruct

func init() {
	UserController = new(UserControllerStruct)
}

func (c *UserControllerStruct) InitUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	money, _ := strconv.ParseFloat(ctx.PostForm("money"), 64)

	record := new(model.TradeUserRecord)
	record.Name = name
	record.RechargeMoney = money
	record.UsableMoney = money
	record.CTime = time.Now().Unix()
	record.UTime = time.Now().Unix()

	id, err := new(model.TradeUserModel).Insert(record)
	if err != nil {
		ReturnError(ctx, 499, err.Error())
	} else {
		ReturnSucc(ctx, map[string]any{"id": id})
	}
}
