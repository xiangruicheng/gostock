package strategy

import "gostock/model"

type TradeDayStruct struct {
	Dates []string
}

var TradeDay *TradeDayStruct

func (td *TradeDayStruct) getKey(date string) int {
	for k, d := range td.Dates {
		if d == date {
			return k
		}
	}
	return -1
}

func (td *TradeDayStruct) IsTradeDay(date string) bool {
	k := td.getKey(date)
	if k >= 0 {
		return true
	}
	return false
}

func (td *TradeDayStruct) PreTradeDate(date string, num int) string {
	var index int = -1
	for k, d := range td.Dates {
		if d >= date {
			index = k
			break
		}
	}
	if index-num >= 0 {
		return td.Dates[index-num]
	}
	return ""
}

func (td *TradeDayStruct) NextTradeDate(date string, num int) string {
	var tagNum int = 0
	for _, d := range td.Dates {
		if d > date {
			tagNum += 1
			if tagNum == num {
				return d
			}
		}
	}
	return ""
}

func (td *TradeDayStruct) InitTradeDay() {
	klines, _ := new(model.KlineModel).GetByCode("1A0001")
	for _, kline := range klines {
		td.Dates = append(td.Dates, kline.Date)
	}
}

func init() {
	TradeDay = new(TradeDayStruct)
}
