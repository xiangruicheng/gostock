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
	upLineRate := (kline.High - max(kline.Open, kline.Close)) / (kline.High - kline.Low)
	entityRate := math.Abs(kline.Open-kline.Close) / (kline.High - kline.Low)
	downLineRate := (min(kline.Open, kline.Close) - kline.Low) / (kline.High - kline.Low)
	if upLineRate > 0.45 && downLineRate > 0.45 && entityRate < 0.01 {
		return true
	}
	return false
}

func (f *FeatureStruct) IsT(code string, date string) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	upLineRate := (kline.High - max(kline.Open, kline.Close)) / (kline.High - kline.Low)
	entityRate := math.Abs(kline.Open-kline.Close) / (kline.High - kline.Low)
	downLineRate := (min(kline.Open, kline.Close) - kline.Low) / (kline.High - kline.Low)
	if upLineRate < 0.1 && downLineRate > 0.9 && entityRate < 0.01 {
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

func (f *FeatureStruct) PercentRange(code string, date string, min float64, max float64) bool {
	kline, _ := new(model.KlineModel).GetByCodeAndDate(code, date)
	if kline.Percent > min && kline.Percent < max {
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

func (f *FeatureStruct) IsCyb(code string) bool {
	info, _ := new(model.StockInfoModel).GetByCode(code)
	if info.Cyb == 1 {
		return true
	}
	return false
}

func (f *FeatureStruct) IsHs300(code string) bool {
	info, _ := new(model.StockInfoModel).GetByCode(code)
	if info.Hs300 == 1 {
		return true
	}
	return false
}

func (f *FeatureStruct) IsKcb(code string) bool {
	if code[0:3] == "688" {
		return true
	}
	return false
}

// VolumeRateRange 量比在min,max范围内
func (f *FeatureStruct) VolumeRateRange(code string, date string, min float64, max float64) bool {
	if !TradeDay.IsTradeDay(date) {
		return false
	}

	// get data
	startDate := TradeDay.PreTradeDate(date, 5)
	klines, _ := new(model.KlineModel).GetByCodeERangeDate(code, startDate, date)
	if len(klines) < 6 {
		return false
	}

	// volumeRate
	currVolume := klines[5].Volume
	var totalVolume, total float64
	for k, kline := range klines {
		if k < 5 {
			totalVolume += kline.Volume
			total += 1
		}

	}
	volumeRate := currVolume / (totalVolume / total)

	if volumeRate >= min && volumeRate <= max {
		return true
	}
	return false
}

// IsLastXDaysCloseMin 是否为最近X天的最低价
func (f *FeatureStruct) IsLastXDaysCloseMin(code string, date string, x int) bool {
	if !TradeDay.IsTradeDay(date) {
		return false
	}
	startDate := TradeDay.PreTradeDate(date, x)

	klines, _ := new(model.KlineModel).GetByCodeERangeDate(code, startDate, date)
	if len(klines) < (x + 1) {
		return false
	}
	currKline := klines[len(klines)-1]
	for _, kline := range klines {
		if kline.Close < currKline.Close {
			return false
		}
	}
	return true
}

// IsLastXDaysMacdMin 是否为最近X天macd最小值
func (f *FeatureStruct) IsLastXDaysMacdMin(code string, date string, x int) bool {
	if !TradeDay.IsTradeDay(date) {
		return false
	}
	startDate := TradeDay.PreTradeDate(date, x)

	macds, _ := new(model.MacdModel).GetByCodeERangeDate(code, startDate, date)
	if len(macds) < (x + 1) {
		return false
	}
	currMacd := macds[len(macds)-1]
	for _, macd := range macds {
		if macd.Macd < currMacd.Macd {
			return false
		}
	}
	return true
}

func init() {
	Feature = new(FeatureStruct)
}
