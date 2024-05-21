package datainit

import (
	"fmt"
	"gostock/config"
	"gostock/data/datasource/eastmoney"
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"strings"
	"time"
)

func UpdateHs300() {
	hs300Res, err := eastmoney.Hs300()
	if err != nil {
		server.Log.Error(fmt.Sprintf("UpdateHs300 fail %s", err.Error()))
		return
	}
	if !hs300Res.Success {
		server.Log.Error(fmt.Sprintf("UpdateHs300 fail %s ", hs300Res.Message))
		return
	}

	codeArr := []string{}
	for _, item := range hs300Res.Result.Data {
		codeArr = append(codeArr, item.SECURITYCODE)
	}
	affected, err := new(model.StockInfoModel).UpdateHs300(codeArr)
	if err != nil {
		server.Log.Error(fmt.Sprintf("UpdateHs300 fail %s ", err.Error()))
		return
	}
	server.Log.Error(fmt.Sprintf("UpdateHs300 succ %d ", affected))
	return
}

func InitStockInfo() {
	initStockType()
	initIndexType()
	initEtfType()
}

func initIndexType() {
	for _, etf := range config.Data.Index {
		market := etf.Market
		stockInfoRecord := new(model.StockInfoRecord)
		stockInfoRecord.Code = etf.Code
		stockInfoRecord.Name = etf.Name
		stockInfoRecord.Market = market
		stockInfoRecord.CTime = time.Now().Unix()
		stockInfoRecord.UTime = time.Now().Unix()
		new(model.StockInfoModel).Insert(stockInfoRecord)
	}
}

func initEtfType() {
	for _, etf := range config.Data.Etf {
		market := etf.Market
		stockInfoRecord := new(model.StockInfoRecord)
		stockInfoRecord.Code = etf.Code
		stockInfoRecord.Name = etf.Name
		stockInfoRecord.Market = market
		stockInfoRecord.CTime = time.Now().Unix()
		stockInfoRecord.UTime = time.Now().Unix()
		new(model.StockInfoModel).Insert(stockInfoRecord)
	}
}

func initStockType() {
	stockCNList, err := xueqiu.StockAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	for _, stockCN := range stockCNList {
		market := stockCN.Code[0:2]
		stockInfoRecord := new(model.StockInfoRecord)
		stockInfoRecord.Code = strings.TrimLeft(stockCN.Code, market)
		stockInfoRecord.Name = stockCN.Name
		stockInfoRecord.Market = market
		stockInfoRecord.CTime = time.Now().Unix()
		stockInfoRecord.UTime = time.Now().Unix()
		new(model.StockInfoModel).Insert(stockInfoRecord)
	}
}
