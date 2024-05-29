package strategy

import "gostock/model"

type TradeDay struct {
	Dates []string
}

var TD *TradeDay

func (td *TradeDay) getKey(date string) int {
	for k, d := range td.Dates {
		if d == date {
			return k
		}
	}
	return -1
}

func (td *TradeDay) IsTradeDay(date string) bool {
	k := td.getKey(date)
	if k >= 0 {
		return true
	}
	return false
}

func (td *TradeDay) PreTradeDate(date string, num int) string {
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

func (td *TradeDay) NextTradeDate(date string, num int) string {
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

func (td *TradeDay) InitTradeDay() {
	klines, _ := new(model.KlineModel).GetByCode("1A0001")
	for _, kline := range klines {
		td.Dates = append(td.Dates, kline.Date)
	}
}

func init() {
	TD = new(TradeDay)
}
