package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostock/model"
	"strconv"
	"time"
)

type OrderControllerStruct struct {
}

var OrderController *OrderControllerStruct

func init() {
	OrderController = new(OrderControllerStruct)
}

func (c *OrderControllerStruct) Buy(ctx *gin.Context) {
	code := ctx.PostForm("code")
	date := ctx.PostForm("date")
	price, _ := strconv.ParseFloat(ctx.PostForm("price"), 64)
	number, _ := strconv.ParseInt(ctx.PostForm("number"), 0, 64)

	// Check params
	if date == "" || code == "" || price <= 0 || number <= 0 {
		ReturnError(ctx, 400, "param error")
		return
	}

	var uid int64 = 1
	// Save User
	userRecord, err := new(model.TradeUserModel).GetById(uid)
	if err != nil {
		ReturnError(ctx, 400, err.Error())
		return
	}

	// Check money
	var fee, money float64
	money = price*float64(number) + fee
	if userRecord.UsableMoney < money {
		ReturnError(ctx, 401, "not usable money")
		return
	}

	// Save Order
	record := new(model.TradeOrderRecord)
	record.Uid = uid
	record.Type = 1
	record.Code = code
	record.Date = date
	record.Price = price
	record.Number = number
	record.Fee = fee
	record.Money = money
	record.CTime = time.Now().Unix()
	record.UTime = time.Now().Unix()
	id, err := new(model.TradeOrderModel).Insert(record)

	// Save Hold
	holdRecord, err := new(model.TradeHoldModel).GetByUidAndCode(uid, code)
	if err != nil {
		ReturnError(ctx, 400, err.Error())
		return
	}
	newNumber := holdRecord.Number + number
	newCost := holdRecord.Price*float64(holdRecord.Number) + money
	newPrice := newCost / float64(newNumber)
	holdRecord.Price = newPrice
	holdRecord.Code = code
	holdRecord.Number = newNumber
	holdRecord.UTime = time.Now().Unix()
	holdRecord.Uid = uid
	if holdRecord.Id == 0 {
		holdRecord.CTime = time.Now().Unix()
		_, err = new(model.TradeHoldModel).Insert(holdRecord)
	} else {
		_, err = new(model.TradeHoldModel).Update(holdRecord)
	}

	//Save user
	userRecord.UsableMoney -= money
	userRecord.UTime = time.Now().Unix()
	fmt.Println(userRecord)
	_, err = new(model.TradeUserModel).Update(userRecord)

	if err != nil {
		ReturnError(ctx, 400, err.Error())
	} else {
		ReturnSucc(ctx, map[string]any{"order_id": id})
	}
}
