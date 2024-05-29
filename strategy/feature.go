package strategy

import (
	"gostock/model"
	"math"
)

type FeatureStruct struct {
}

var Feature *FeatureStruct

func (f *FeatureStruct) IsMacdGold(code string, date string) bool {
	currMacd, _ := new(model.MacdModel).GetByCodeAndDate(code, date)
	preTradeDay := TradeDay.PreTradeDate(date, 1)
	if preTradeDay == "" {
		return false
	}
	preMacd, _ := new(model.MacdModel).GetByCodeAndDate(code, preTradeDay)

	if preMacd.Macd < 0 && currMacd.Macd > 0 {
		return true
	}
	return false
}

func (f *FeatureStruct) IsMacdDie(code string, date string) bool {
	currMacd, _ := new(model.MacdModel).GetByCodeAndDate(code, date)
	nextTradeDay := TradeDay.NextTradeDate(date, 1)
	if nextTradeDay == "" {
		return false
	}
	nextMacd, _ := new(model.MacdModel).GetByCodeAndDate(code, nextTradeDay)

	if nextMacd.Macd < 0 && currMacd.Macd > 0 {
		return true
	}
	return false
}

func (f *FeatureStruct) IsCrossStar(code string, date string) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	if math.Abs(kline.Open-kline.Close)/(kline.High-kline.Low) < 0.0001 {
		return true
	}
	return false
}

func (f *FeatureStruct) PercentGtX(code string, date string, x float64) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	if kline.Percent > x {
		return true
	}
	return false
}

func (f *FeatureStruct) PercentLtX(code string, date string, x float64) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	if kline.Percent < x {
		return true
	}
	return false
}

func init() {
	Feature = new(FeatureStruct)
}
