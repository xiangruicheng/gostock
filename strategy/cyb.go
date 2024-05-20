package strategy

import (
	"fmt"
	"gostock/model"
)

func Cyb() {
	arr := [][]string{
		{"20100101", "20110101"},
		{"20110101", "20120101"},
		{"20120101", "20130101"},
		{"20130101", "20140101"},
		{"20140101", "20150101"},
		{"20150101", "20160101"},
		{"20160101", "20170101"},
		{"20170101", "20180101"},
		{"20180101", "20190101"},
		{"20190101", "20200101"},
		{"20200101", "20210101"},
		{"20210101", "20220101"},
		{"20220101", "20230101"},
		{"20230101", "20240101"},
		{"20240101", "20250101"},
	}

	for _, t := range arr {
		cybStrategy(t[0], t[1])
	}

}

// cybStrategy1
func cybStrategy(min string, max string) {
	klines, _ := new(model.KlineModel).GetByCodeRangeDate("159915", "20100101", "20250101")
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
