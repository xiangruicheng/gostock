package datainit

import (
	"fmt"
	"gostock/data/datasource/eastmoney"
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"time"
)

func BatchUpdateStockPeople() {
	stockCNList, err := xueqiu.StockAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	count := len(stockCNList)
	for k, stockCN := range stockCNList {
		UpdateStockQuote(stockCN.Code)
		server.Log.Info(fmt.Sprintf("update quote %d/%d %s", k, count, stockCN.Code))
	}
}

func UpdateStockPeople(code string) {

	peopleRes, err := eastmoney.People(code)
	if err != nil {
		server.Log.Error(fmt.Sprintf("UpdateStockPeople fail:%s,%s ", code, err.Error()))
		return
	}
	if !peopleRes.Success {
		server.Log.Error(fmt.Sprintf("UpdateStockPeople fail:%s,%s ", code, peopleRes.Message))
		return
	}

	stockPeopleRecords := []*model.StockPeopleRecord{}
	for _, item := range peopleRes.Result.Data {
		stockPeopleRecord := new(model.StockPeopleRecord)
		stockPeopleRecord.Code = code
		stockPeopleRecord.HolderNum = item.HOLDERNUM
		stockPeopleRecord.AvgMarket = item.AVGMARKETCAP
		stockPeopleRecord.AvgHoldNum = item.AVGHOLDNUM
		stockPeopleRecord.CTime = time.Now().Unix()
		stockPeopleRecord.UTime = time.Now().Unix()

		t, _ := time.Parse("2006-01-02 15:04:05", item.ENDDATE)
		stockPeopleRecord.Date = t.Format("20060102")

		stockPeopleRecords = append(stockPeopleRecords, stockPeopleRecord)
	}
	new(model.StockPeopleModel).BatchInsert(stockPeopleRecords)
}
