package strategy

import (
	"fmt"
	"gostock/model"
)

type MacdStrategy struct {
}

func (s *MacdStrategy) gold(code string) []*model.MacdRecord {
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

func (s *MacdStrategy) dbl(code string) []*model.MacdRecord {
	list := []*model.MacdRecord{}
	macds := s.gold(code)
	for k, curMacd := range macds {
		if k-1 < 0 {
			continue
		}

		preMacd := macds[k-1]
		percent := ((curMacd.Close - preMacd.Close) / preMacd.Close) * 100
		if preMacd.Diff < curMacd.Diff &&
			preMacd.Dea < curMacd.Dea &&
			percent < -1.0 {
			list = append(list, curMacd)
		}
	}
	return list
}

func (s *MacdStrategy) gold2buy(code string) []*model.MacdRecord {
	list := []*model.MacdRecord{}
	macds := s.gold(code)
	for k, curMacd := range macds {
		if k-2 < 0 {
			continue
		}
		preMacd := macds[k-1]
		pre2Macd := macds[k-2]
		if preMacd.Diff < 0 && preMacd.Dea < 0 &&
			pre2Macd.Diff < 0 && pre2Macd.Dea < 0 &&
			curMacd.Diff > 0 && curMacd.Dea > 0 {
			list = append(list, curMacd)
		}
	}
	return list
}

type StrategyResult struct {
	UpNum        int
	DownNum      int
	TotalNum     int
	TotalChg     float64
	TotalPercent float64
	MaxPercent   float64
	MinPercent   float64
}

func (s *MacdStrategy) printStrategyResult(r *StrategyResult) {
	fmt.Printf("UpNum=%d    ", r.UpNum)
	fmt.Printf("DownNum=%d  ", r.DownNum)
	fmt.Printf("TotalNum=%d ", r.TotalNum)
	fmt.Printf("TotalChg=%f ", r.TotalChg)
	fmt.Printf("TotalPercent=%.2f ", r.TotalPercent)
	fmt.Printf("MaxPercent=%.2f ", r.MaxPercent)
	fmt.Printf("MinPercent=%.2f\n", r.MinPercent)

}

func (s *MacdStrategy) runOne(code string) *StrategyResult {
	sr := new(StrategyResult)

	klines, _ := new(model.KlineModel).GetByCode(code)
	goldList := s.dbl(code)

	for _, startKline := range goldList {
		endKline := afterXKline(klines, startKline.Date, 5)
		if endKline == nil {
			continue
		}
		chg := endKline.Close - startKline.Close
		percent := (chg / startKline.Close) * 100

		if chg > 0 {
			sr.UpNum += 1
		}
		if chg < 0 {
			sr.DownNum += 1
		}
		if percent > sr.MaxPercent {
			sr.MaxPercent = percent
		}
		if percent < sr.MinPercent {
			sr.MinPercent = percent
		}
		sr.TotalChg += chg
		sr.TotalPercent += percent
		sr.TotalNum += 1

		//fmt.Printf("%s %f\n", startKline.Date, percent)
	}
	return sr
}

func (s *MacdStrategy) Run() {

	list, _ := new(model.StockInfoModel).GetAllByTag("hs300")
	TotalResult := new(StrategyResult)
	for k, item := range list {
		oneResult := s.runOne(item.Code)
		if oneResult.TotalNum < 2 {
			continue
		}
		fmt.Printf("%d:%s  %d/%d\n", k, item.Code, oneResult.UpNum, oneResult.TotalNum)
		s.printStrategyResult(oneResult)

		TotalResult.UpNum += oneResult.UpNum
		TotalResult.DownNum += oneResult.DownNum
		TotalResult.TotalChg += oneResult.TotalChg
		TotalResult.TotalPercent += oneResult.TotalPercent
		TotalResult.TotalNum += oneResult.TotalNum
	}

	s.printStrategyResult(TotalResult)
}

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
