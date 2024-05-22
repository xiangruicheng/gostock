package datainit

import (
	"fmt"
	"gostock/config"
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"gostock/util"
	"time"
)

func InitKline(code string, market string, days int64) int64 {
	symbol := market + code
	klines, err := xueqiu.KlineList(symbol, days)
	if err != nil {
		server.Log.Error(err.Error())
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
		server.Log.Error(err.Error())
		return 0
	}
	return affectedRows
}

func BatchInitKline() {
	stockInfoRecords, err := new(model.StockInfoModel).GetAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	for key, stockInfoRecord := range stockInfoRecords {
		InitKline(stockInfoRecord.Code, stockInfoRecord.Market, config.Data.Xueqiu.InitNum)
		server.Log.Info(fmt.Sprintf("init kline key=%d code=%s", key, stockInfoRecord.Code))
	}
}

func BatchIncrKline() {
	stockInfoRecords, err := new(model.StockInfoModel).GetAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	for key, stockInfoRecord := range stockInfoRecords {
		InitKline(stockInfoRecord.Code, stockInfoRecord.Market, config.Data.Xueqiu.IncrNum)
		server.Log.Info(fmt.Sprintf("incr kline key=%d code=%s", key, stockInfoRecord.Code))
	}
}
