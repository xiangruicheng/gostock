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

func (f *FeatureStruct) IsCrossStar(code string, date string) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	if math.Abs(kline.Open-kline.Close)/(kline.High-kline.Low) < 0.0001 {
		return true
	}
	return false
}

func init() {
	Feature = new(FeatureStruct)
}
