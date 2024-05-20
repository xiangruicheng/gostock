package datainit

import (
	"gostock/data/datasource/xueqiu"
	"gostock/model"
	"gostock/server"
	"strings"
	"time"
)

func BatchUpdateStockBaseInfo() {
	stockCNList, err := xueqiu.StockAll()
	if err != nil {
		server.Log.Error(err.Error())
		return
	}
	for _, stockCN := range stockCNList {
		market := stockCN.Code[0:2]
		code := strings.TrimLeft(stockCN.Code, market)
		quoteRes, _ := xueqiu.Quote(stockCN.Code)
		quote := quoteRes.Data.Quote

		stockQuoteRecord := new(model.StockQuoteRecord)
		stockQuoteRecord.Code = code
		stockQuoteRecord.Name = quote.Name
		//stockQuoteRecord.PeForecast = quote.PeForecast
		//stockQuoteRecord.PeTtm = quote.PeTtm
		//stockQuoteRecord.PeLyr = quote.PeLyr
		//stockQuoteRecord.Pb = quote.Pb
		stockQuoteRecord.TotalShares = quote.TotalShares
		stockQuoteRecord.FloatShares = quote.FloatShares
		stockQuoteRecord.FloatMarketCapital = quote.FloatMarketCapital
		stockQuoteRecord.MarketCapital = quote.MarketCapital
		stockQuoteRecord.Amount = quote.Amount
		stockQuoteRecord.Volume = quote.Volume
		//stockQuoteRecord.TurnoverRate = quote.TurnoverRate
		stockQuoteRecord.Amplitude = quote.Amplitude
		//stockQuoteRecord.Navps = quote.Navps
		//stockQuoteRecord.Eps = quote.Eps
		//stockQuoteRecord.VolumeRatio = quote.VolumeRatio
		//stockQuoteRecord.PankouRatio = quote.PankouRatio
		stockQuoteRecord.High = quote.High
		stockQuoteRecord.Low = quote.Low
		stockQuoteRecord.Open = quote.Open
		stockQuoteRecord.Current = quote.Current
		//stockQuoteRecord.Dividend = quote.Dividend
		//stockQuoteRecord.DividendYield = quote.DividendYield
		//stockQuoteRecord.Date = quote.Date
		stockQuoteRecord.CTime = time.Now().Unix()
		stockQuoteRecord.UTime = time.Now().Unix()

	}
}
