package datainit

import (
	"gostock/config"
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"strings"
	"time"
)

func InitStockInfo() {
	initIndexType()
	initEtfType()
	initStockType()
}

func initIndexType() {
	for _, etf := range config.Data.Index {
		market := etf.Market
		stockInfoRecord := new(model.StockInfoRecord)
		stockInfoRecord.Code = etf.Code
		stockInfoRecord.Name = etf.Name
		stockInfoRecord.Market = market
		stockInfoRecord.Type = model.StockInfoModel_TypeIndex
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
		stockInfoRecord.Type = model.StockInfoModel_TypeEtf
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
		stockInfoRecord.Type = model.StockInfoModel_TypeStock
		stockInfoRecord.CTime = time.Now().Unix()
		stockInfoRecord.UTime = time.Now().Unix()
		new(model.StockInfoModel).Insert(stockInfoRecord)
	}
}
