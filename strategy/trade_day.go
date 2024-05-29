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
	k := td.getKey(date)
	if k < 0 {
		return ""
	}
	if k-num < 0 {
		return ""
	}
	return td.Dates[k-num]
}

func (td *TradeDay) NextTradeDate(date string, num int) string {
	k := td.getKey(date)
	if k < 0 {
		return ""
	}
	if k+num > len(td.Dates) {
		return ""
	}
	return td.Dates[k+num]
}

func (td *TradeDay) InitTradeDay() {
	klines, _ := new(model.KlineModel).Query("code='1A0001' order by date asc")
	for _, kline := range klines {
		td.Dates = append(td.Dates, kline.Date)
	}
}

func init() {
	TD = new(TradeDay)
}
