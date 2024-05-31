package datainit

import (
	"fmt"
	"gostock/data/datasource/eastmoney"
	"gostock/model"
	"gostock/server"
	"time"
)

func InitBlock() {
	initStockBlock()
	initStockBlockCode()
}

func initStockBlock() {
	for _, blockType := range []int64{1, 2} {
		blockAll, err := eastmoney.Block(blockType)
		if err != nil {
			server.Log.Error(fmt.Sprintf("initStockType %s fail,%s", err.Error()))
			return
		}
		records := []*model.StockBlockRecord{}
		for _, item := range blockAll.Data.Diff {
			record := new(model.StockBlockRecord)
			record.Type = blockType
			record.Code = item.Code
			record.Name = item.Name
			record.CTime = time.Now().Unix()
			record.UTime = time.Now().Unix()
			records = append(records, record)
		}
		new(model.StockBlockModel).BatchInsert(records)
		server.Log.Info(fmt.Sprintf("InitStockBlock %s succ", blockType))
	}
}

func initStockBlockCode() {
	blocks, _ := new(model.StockBlockModel).GetAll()
	for _, block := range blocks {
		blockCodeAll, err := eastmoney.BlockCode(block.Code)
		if err != nil {
			server.Log.Error(fmt.Sprintf("InitStockBlockCode %s fail,%s", err.Error()))
			return
		}
		records := []*model.StockBlockCodeRecord{}
		for _, item := range blockCodeAll.Data.Diff {
			record := new(model.StockBlockCodeRecord)
			record.BkCode = block.Code
			record.Code = item.Code
			record.CTime = time.Now().Unix()
			record.UTime = time.Now().Unix()
			records = append(records, record)
		}
		new(model.StockBlockCodeModel).BatchInsert(records)
		server.Log.Info(fmt.Sprintf("InitStockBlockCode %s succ", block.Code))
	}
}
