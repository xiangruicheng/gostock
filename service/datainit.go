package datainit

import (
	"fmt"
	"gostock/config"
	"gostock/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"gostock/util"
	"strings"
	"time"
)

func InitKline(code string, market string) int64 {
	symbol := market + code
	klines, err := xueqiu.KlineList(symbol, config.Data.Xueqiu.InitNum)
	if err != nil {
		server.Log(server.LogLevelError, err.Error())
		return 0
	}
	klineRecords := []*model.KlineRecord{}
	for _, kline := range klines {
		klineRecord := new(model.KlineRecord)
		klineRecord.Code = code
		klineRecord.Date = util.FormatDate(kline.Time)
		klineRecord.Volume = kline.Volume
		klineRecord.Amount = kline.Amount
		klineRecord.Open = kline.Open
		klineRecord.High = kline.High
		klineRecord.Low = kline.Low
		klineRecord.Close = kline.Close
		klineRecord.Chg = kline.Chg
		klineRecord.Percent = kline.Percent
		klineRecord.CTime = time.Now().Unix()
		klineRecord.UTime = time.Now().Unix()
		klineRecords = append(klineRecords, klineRecord)
	}
	affectedRows, err := new(model.KlineModel).BatchInsert(klineRecords)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return affectedRows
}

func BatchInitKline() {
	stockInfoRecords, err := new(model.StockInfoModel).GetAll()
	if err != nil {
		server.Log(server.LogLevelError, err.Error())
		return
	}
	for _, stockInfoRecord := range stockInfoRecords {
		InitKline(stockInfoRecord.Code, stockInfoRecord.Market)
		server.Log(server.LogLevelInfo, fmt.Sprintf("init kline code=%s", stockInfoRecord.Code))
	}
}

func InitStockInfo() int64 {
	stockCNList, err := xueqiu.StockAll()
	if err != nil {
		fmt.Println(err)
		return 0
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
	return 0
}
