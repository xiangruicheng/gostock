package datainit

import (
	"fmt"
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"time"
)

func BatchUpdateStockQuote() {
	stockAll, err := new(model.StockInfoModel).GetAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	count := len(stockAll)
	for k, item := range stockAll {
		UpdateStockQuote(item.Market, item.Code)
		server.Log.Info(fmt.Sprintf("update quote %d/%d %s", k, count, item.Code))
	}
}

func UpdateStockQuote(market string, code string) {
	symbol := market + code
	quoteRes, _ := xueqiu.Quote(symbol)
	quote := quoteRes.Data.Quote

	stockQuoteRecord := new(model.StockQuoteRecord)
	stockQuoteRecord.Code = code
	stockQuoteRecord.Name = quote.Name
	stockQuoteRecord.PeForecast = quote.PeForecast
	stockQuoteRecord.PeTtm = quote.PeTtm
	stockQuoteRecord.PeLyr = quote.PeLyr
	stockQuoteRecord.Pb = quote.Pb
	stockQuoteRecord.TotalShares = quote.TotalShares
	stockQuoteRecord.FloatShares = quote.FloatShares
	stockQuoteRecord.FloatMarketCapital = quote.FloatMarketCapital
	stockQuoteRecord.MarketCapital = quote.MarketCapital
	stockQuoteRecord.Amount = quote.Amount
	stockQuoteRecord.Volume = quote.Volume
	stockQuoteRecord.TurnoverRate = quote.TurnoverRate
	stockQuoteRecord.Amplitude = quote.Amplitude
	stockQuoteRecord.Navps = quote.Navps
	stockQuoteRecord.Eps = quote.Eps
	stockQuoteRecord.VolumeRatio = quote.VolumeRatio
	stockQuoteRecord.PankouRatio = quoteRes.Data.Others.PankouRatio
	stockQuoteRecord.High = quote.High
	stockQuoteRecord.Low = quote.Low
	stockQuoteRecord.Open = quote.Open
	stockQuoteRecord.Current = quote.Current
	stockQuoteRecord.Dividend = quote.Dividend
	stockQuoteRecord.DividendYield = quote.DividendYield
	stockQuoteRecord.Date = time.Now().Format("20060102")
	stockQuoteRecord.CTime = time.Now().Unix()
	stockQuoteRecord.UTime = time.Now().Unix()

	model := new(model.StockQuoteModel)
	isExist, _ := model.IsExist(code)
	if isExist {
		model.Update(stockQuoteRecord)
	} else {
		model.Insert(stockQuoteRecord)
	}
}
