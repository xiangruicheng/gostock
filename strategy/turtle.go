package strategy

import (
	"fmt"
	"gostock/model"
	"math"
	"strconv"
)

// 海龟策略
type Turtle struct {
	Code       string  // 股票代码
	TotalMoney float64 // 总金额
	Risk       float64 // 承受的风险比例
}

// 海龟K线数据
type TurtleKline struct {
	Date              string
	Atr               float64
	Open              float64
	Close             float64
	High              float64
	Low               float64
	Tr                float64 // TR振幅
	High20            float64 // 当日的近20日最高价，收盘后才知道
	Low10             float64 // 当日的近10日最低价，收盘后才知道
	IsBreakoutUp      bool    // 是否向上突破，与上一日的high20比较，用于回测
	IsBreakoutFirstUp bool    // 是否首次向上突破
	IsBreakoutDown    bool    // 是否向下破位，与上一日的low10比较，用于回测
}

type TradeR struct {
	Atr       float64 // atr 平均振幅
	Num       int64   //持仓数量
	AvgPrice  float64 //平均价格
	BuyPrice  float64 //最后买入价格
	BuyDate   string  //最后买入日期
	SellPrice float64 //卖出价格
	SellDate  string  //卖出日期
	Earn      float64 //收益
}

func (s *Turtle) Help() {
	fmt.Println("海龟策略")
	// 使用方法
	/*s := &strategy.Turtle{
		Code:       "512480",
		TotalMoney: 1000000.00,
		Risk:       0.01,
	}
	s.Run()*/
}

func (s *Turtle) Run() {
	if s.Code == "" {
		fmt.Printf("缺少股票代码")
		return
	}
	klines, _ := new(model.KlineModel).GetByCode(s.Code)
	turtleKlines := s.initTurtleKline(klines)

	trade := new(Trade)
	for k, turtleKline := range turtleKlines {
		if !trade.IsHold {
			if turtleKline.IsBreakoutFirstUp {
				buyPrice := turtleKlines[k-1].High20
				atr, position := s.position(turtleKline)
				trade.Buy(turtleKline.Date, buyPrice, position, fmt.Sprintf("%f", atr))
			}
		}
		// sell
		if trade.IsHold {
			sellPrice := turtleKlines[k-1].Low10 //卖出价格
			lastTradeRecord := trade.GetLastRecord()
			lastAtr, _ := strconv.ParseFloat(lastTradeRecord.Other, 64)
			stopPrice := lastTradeRecord.Price - 2*lastAtr //止损价格

			if turtleKline.Low < stopPrice && turtleKline.IsBreakoutDown {
				trade.Stop(turtleKline.Date, max(sellPrice, stopPrice), trade.GetTotalNum(), "")
			} else if turtleKline.IsBreakoutDown {
				trade.Sell(turtleKline.Date, sellPrice, trade.GetTotalNum(), "")
			} else if turtleKline.Low < stopPrice {
				trade.Stop(turtleKline.Date, stopPrice, trade.GetTotalNum(), "")
			}

		}

	}
	trade.Report()
}

// 计算仓位
// - 风险金额 = 总资金 × 1% = 100万 × 0.01 = 1万元
// - 头寸规模 = 风险金额 / ATR = 10000 / 2.5 = 4000股
// - 买入4000股，成本51元，总投入20.4万元。
// 返回为股票数
func (s *Turtle) position(atrData *TurtleKline) (atr float64, position int64) {
	riskMoney := s.TotalMoney * s.Risk
	atr = atrData.Atr
	position = int64(riskMoney / atr)
	return atr, position
}

func (s *Turtle) getTurtleKlineByDate(atrs []*TurtleKline) (atr *TurtleKline) {

	return nil
}

// 计算ATR
// 公式： N=(19*PDN+TR)/20
// ATR的计算需要先计算每日的真实波幅（TR），然后取20日平均。
// TR = max(当日最高-当日最低, |当日最高-前日收盘|, |当日最低-前日收盘|)
// 然后ATR(20)就是过去20日TR的平均值。
func (s *Turtle) initTurtleKline(klines []*model.Kline) (turtleKlines []*TurtleKline) {
	for i, kline := range klines {
		tr := kline.High - kline.Low
		if i > 0 {
			tr = math.Max(math.Max(kline.High-kline.Low, math.Abs(kline.High-klines[i-1].Close)), math.Abs(kline.Low-klines[i-1].Close))
		}
		tr = math.Round(tr*10000) / 10000

		sumTr := 0.0
		sumNum := 0
		high20 := 0.0
		low10 := 10000.0
		for k := 0; k < 20; k++ {
			if i-k >= 0 {
				sumTr += tr
				sumNum += 1
				high20 = math.Max(high20, klines[i-k].High)
			}
		}
		for k := 0; k < 10; k++ {
			if i-k >= 0 {
				low10 = math.Min(low10, klines[i-k].Low)
			}
		}
		atr := sumTr / float64(sumNum)
		atr = math.Round(atr*10000) / 10000

		isBreakoutUp := false
		isBreakoutDown := false
		if i > 0 {
			if kline.High > turtleKlines[i-1].High20 {
				isBreakoutUp = true
			}
			if kline.Low < turtleKlines[i-1].Low10 {
				isBreakoutDown = true
			}
		}
		isBreakoutFirstUp := isBreakoutUp

		for k := 1; k < 20; k++ {
			if i-k >= 0 {
				if turtleKlines[i-k].IsBreakoutUp == true {
					isBreakoutFirstUp = false
				}
			}
		}

		turtleKlines = append(turtleKlines, &TurtleKline{
			Date:              kline.Date,
			Open:              kline.Open,
			Close:             kline.Close,
			High:              kline.High,
			Low:               kline.Low,
			Tr:                tr,
			Atr:               atr,
			High20:            high20,
			Low10:             low10,
			IsBreakoutUp:      isBreakoutUp,
			IsBreakoutFirstUp: isBreakoutFirstUp,
			IsBreakoutDown:    isBreakoutDown,
		})
	}
	return turtleKlines
}
