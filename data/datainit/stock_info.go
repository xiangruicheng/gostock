package datainit

import (
	"fmt"
	"gostock/config"
	"gostock/data/datasource/eastmoney"
	"gostock/model"
	"gostock/server"
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
	server.Log.Info(fmt.Sprintf("initIndexType succ"))
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
	server.Log.Info(fmt.Sprintf("initEtfType succ"))

}

func initStockType() {
	for _, market := range []string{"SH", "SZ"} {
		stockAll, err := eastmoney.StockAll(market)
		if err != nil {
			server.Log.Error(fmt.Sprintf("initStockType %s fail,%s", market, err.Error()))
			return
		}
		for _, item := range stockAll.Data.Diff {
			stockInfoRecord := new(model.StockInfoRecord)
			stockInfoRecord.Code = item.Code
			stockInfoRecord.Name = item.Name
			stockInfoRecord.Market = market
			stockInfoRecord.CTime = time.Now().Unix()
			stockInfoRecord.UTime = time.Now().Unix()
			new(model.StockInfoModel).Insert(stockInfoRecord)
		}
		server.Log.Info(fmt.Sprintf("initStockType %s succ", market))
	}
}
