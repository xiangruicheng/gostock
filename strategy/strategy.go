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
	date := "20240531"
	all, _ := new(model.StockInfoModel).GetAll()
	for _, item := range all {
		if item.Code[0:1] == "3" || item.Code[0:3] == "688" {
			continue
		}
		if !Feature.IsLastXDaysMin(item.Code, date, 10) {
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
		if date < "20230101" || date > "20240101" {
			continue
		}
		for _, item := range all {
			if item.Code[0:1] == "3" || item.Code[0:3] == "688" {
				continue
			}
			if !Feature.IsMacdGold(item.Code, date) {
				continue
			}
			if !Feature.IsT(item.Code, date) {
				continue
			}
			buyKline, _ := new(model.KlineModel).GetByCodeAndDate(item.Code, date)
			if buyKline.Percent < 9 {
				continue
			}

			buy := buyKline.Close
			nextDay := TradeDay.NextTradeDate(date, 1)
			if nextDay == "" {
				continue
			}
			sellKline, _ := new(model.KlineModel).GetByCodeAndDate(item.Code, nextDay)
			if sellKline.Date == "" {
				continue
			}
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

func Test4() {

	for _, date := range TradeDay.Dates {
		if date < "20230101" || date > "20250101" {
			continue
		}
		startDate := TradeDay.PreTradeDate(date, 11)
		endDate := TradeDay.NextTradeDate(date, 11)
		if startDate == "" || endDate == "" {
			continue
		}
		currKline, _ := new(model.KlineModel).GetByCodeAndDate("000001", date)
		klines, _ := new(model.KlineModel).GetByCodeRangeDate("000001", startDate, endDate)
		tag := false
		for _, kline := range klines {
			if max(kline.Close, kline.Open) > max(currKline.Open, currKline.Close) {
				tag = true
				continue
			}
		}
		if tag == false {
			fmt.Printf("%s %f\n", date, max(currKline.Close, currKline.Open))
		}
	}

}
