package datainit

import (
	"fmt"
	"gostock/config"
	"gostock/datasource/xueqiu"
	"gostock/model"
	"gostock/util"
	"time"
)

func Kline(code string) int64 {
	symbol := "SZ" + code
	klines, err := xueqiu.KlineList(symbol, config.Data.Xueqiu.InitNum)
	if err != nil {
		fmt.Printf("%s初始化错误%s\n", symbol, err.Error())
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
