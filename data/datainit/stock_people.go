package datainit

import (
	"fmt"
	"gostock/data/datasource/eastmoney"
	"gostock/model"
	"gostock/server"
	"time"
)

func BatchUpdateStockPeople() {
	list, err := new(model.StockInfoModel).GetAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	count := len(list)
	for k, item := range list {
		UpdateStockPeople(item.Code)
		server.Log.Info(fmt.Sprintf("update pepple %d/%d %s", k, count, item.Code))
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
