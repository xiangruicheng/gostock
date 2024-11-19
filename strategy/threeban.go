package strategy

import (
	"fmt"
	"gostock/model"
)

type Threeban struct {
}

func (s *Threeban) Run() {
	percent := 9.98
	day := 2
	all, _ := new(model.StockInfoModel).GetAll()
	for _, item := range all {
		klines, _ := new(model.KlineModel).GetByCode(item.Code)
		for k, kline := range klines {
			if k < 2 {
				continue
			}
			if kline.Percent >= percent &&
				klines[k-1].Percent >= percent &&
				klines[k-2].Percent >= percent {
				if (k + day) < len(klines) {
					fmt.Printf("%s %s %f %f\n", kline.Code, kline.Date, kline.Close, klines[k+day].Close)
				}
			}
		}
	}
}
