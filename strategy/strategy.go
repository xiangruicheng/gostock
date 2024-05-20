package strategy

import "gostock/model"

// afterXKline Date After X day Kline
func afterXKline(klines []*model.KlineRecord, date string, x int) *model.KlineRecord {
	xTag := 0
	for _, kline := range klines {
		if kline.Date > date {
			xTag += 1
			if xTag == x {
				return kline
			}
		}
	}
	return nil
}
