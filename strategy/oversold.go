package strategy

import (
	"fmt"
	"gostock/model"
)

type Oversold struct {
	Code    string  // 股票代码
	Percent float64 //超跌幅度
	MinDate string  //开始日期
	MaxDate string  //结束日期

}

func (s *Oversold) Help() {
	fmt.Println("抄底策略")
	fmt.Println("当天跌幅超过Percent% ")

}

func (s *Oversold) Run() {
	s.runRangeTime(s.MinDate, s.MaxDate)
}

func (s *Oversold) runRangeTime(min string, max string) {
	klines, _ := new(model.KlineModel).GetByCodeRangeDate(s.Code, min, max)
	var downDyas int
	var nextDayPercentTotal float64

	for k, kline := range klines {
		if kline.Percent < -2 {
			if k+1 > len(klines)-1 {
				continue
			}
			nextKline := klines[k+1]
			percent := (nextKline.Open - kline.Close) / kline.Close
			if percent > 0 {
				nextDayPercentTotal += nextKline.Percent
				//fmt.Printf("------%s %f\n", kline.Date, nextKline.Percent)
			} else {
				nextDayPercentTotal += percent
				//fmt.Printf("------%s %f\n", kline.Date, percent)

			}
			downDyas += 1
		}
	}
	fmt.Printf("%s-%s %d/%d %f\n", min, max, downDyas, len(klines), nextDayPercentTotal)
}
