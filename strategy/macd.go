package strategy

import (
	"fmt"
	"gostock/model"
)

func macdGold(code string) []*model.MacdRecord {
	list := []*model.MacdRecord{}
	macds, _ := new(model.MacdModel).GetByCode(code)
	for k, curMacd := range macds {
		if k-1 < 0 {
			continue
		}
		preMacd := macds[k-1]
		if preMacd.Macd < 0 && curMacd.Macd > 0 {
			list = append(list, curMacd)
		}
	}
	return list
}

func macdBL(code string) []*model.MacdRecord {
	list := []*model.MacdRecord{}
	macds := macdGold(code)
	for k, curMacd := range macds {
		if k-1 < 0 {
			continue
		}
		preMacd := macds[k-1]
		if preMacd.Diff < curMacd.Diff &&
			preMacd.Dea < curMacd.Dea &&
			preMacd.Close > curMacd.Close {
			list = append(list, curMacd)
		}
	}
	return list
}

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

func MacdOne(code string) (int, int) {
	klines, _ := new(model.KlineModel).GetByCode(code)
	goldList := macdBL(code)
	var upNum int = 0
	for _, macd := range goldList {
		kline := afterXKline(klines, macd.Date, 15)
		if kline == nil {
			continue
		}

		chg := kline.Close - macd.Close
		//percent := chg / macd.Close
		if chg > 0 {
			upNum += 1
		}
		//fmt.Printf("%s %f %f\n", macd.Date, chg, percent)
	}

	return upNum, len(goldList)

}

func MacdStragegy() {
	var upNumTotal, numTotal int
	list, _ := new(model.StockInfoModel).GetAll()
	for k, item := range list {
		upNum, num := MacdOne(item.Code)
		fmt.Printf("%d:%s  %d/%d\n", k, item.Code, upNum, num)
		upNumTotal += upNum
		numTotal += num
	}

	fmt.Printf("%d/%d= %f\n", upNumTotal, numTotal, upNumTotal/numTotal)
}
