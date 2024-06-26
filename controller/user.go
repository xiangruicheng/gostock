package controller

import (
	"github.com/gin-gonic/gin"
	"gostock/model"
	"gostock/strategy"
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

func (c *UserControllerStruct) Hold(ctx *gin.Context) {
	date := ctx.PostForm("date")
	if date == "" {
		ReturnError(ctx, 400, "date is null")
		return
	}
	if !strategy.TradeDay.IsTradeDay(date) {
		ReturnError(ctx, 400, "date not is trade day")
		return
	}

	type Hold struct {
		Code      string  `json:"code"`
		CostPrice float64 `json:"cost_price"`
		CurrPrice float64 `json:"curr_price"`
		Number    int64   `json:"number"`
		Cost      float64 `json:"cost"`
		Percent   float64 `json:"percent"`
		Money     float64 `json:"money"`
	}

	var uid int64 = 1
	var holdMoney float64
	holdArr := []*Hold{}
	holdRecordArr, _ := new(model.TradeHoldModel).GetByUid(uid)
	for _, holdRecord := range holdRecordArr {
		kline, _ := new(model.KlineModel).GetByCodeAndDate(holdRecord.Code, date)
		hold := new(Hold)
		hold.CostPrice = holdRecord.Price
		hold.CurrPrice = kline.Close
		hold.Number = holdRecord.Number
		hold.Code = holdRecord.Code
		hold.Cost = hold.CostPrice * float64(holdRecord.Number)
		hold.Money = hold.CurrPrice * float64(holdRecord.Number)
		hold.Percent = (hold.Money - hold.Cost) / hold.Cost
		holdArr = append(holdArr, hold)

		holdMoney += hold.Money
	}

	userRecord, _ := new(model.TradeUserModel).GetById(uid)
	userRecord.HoldMoney = holdMoney

	result := map[string]any{
		"user": userRecord,
		"hold": holdArr,
	}
	ReturnSucc(ctx, result)
}
