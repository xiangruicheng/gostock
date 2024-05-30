package strategy

import (
	"fmt"
	"gostock/model"
)

type Strategy interface {
	Run()
}

// hs300:  filter: -4 to -6  buy:close sell:2 day close
// 策略2：cyb: filter:>-2 buy:close sell:2day open
// 策略3：MACD金叉 T型 第二天开盘上涨等到收盘卖，第二天下跌开盘就卖 胜率2：1

func Test() {
	date := "20240530"
	all, _ := new(model.StockInfoModel).GetAll()
	for _, item := range all {
		if !Feature.IsMacdGold(item.Code, date) {
			continue
		}
		if !Feature.IsT(item.Code, date) {
			continue
		}
		fmt.Println(item.Code)
	}
}

func Test3() {
	all, _ := new(model.StockInfoModel).GetAll()

	var upNum, downNum int
	var TotalPercent float64
	for _, date := range TradeDay.Dates {
		if date < "20240101" || date > "20250101" {
			continue
		}
		for _, item := range all {
			if !Feature.IsMacdGold(item.Code, date) {
				continue
			}
			if !Feature.IsT(item.Code, date) {
				continue
			}

			buyKline, _ := new(model.KlineModel).GetByCodeAndDate(item.Code, date)
			buy := buyKline.Close
			nextDay := TradeDay.NextTradeDate(date, 1)
			if nextDay == "" {
				continue
			}
			sellKline, _ := new(model.KlineModel).GetByCodeAndDate(item.Code, nextDay)
			sell := sellKline.Close
			if sellKline.Open > buyKline.Close {
				sell = sellKline.Open
			}
			percent := ((sell - buy) / buy) * 100

			fmt.Printf("%s %s %f %f %f\n", item.Code, date, buy, sell, percent)

			if percent > 0 {
				upNum += 1
			} else {
				downNum += 1
			}
			TotalPercent += percent
		}
	}

	fmt.Printf("%f %f %f\n", upNum, downNum, TotalPercent)

}
