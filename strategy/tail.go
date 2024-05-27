package strategy

import (
	"fmt"
	"gostock/model"
	"strings"
	"time"
)

type TailStrategy struct {
}

type TailResult struct {
	UpNum     int
	DownNum   int
	UpPercent float64
}

func (s *TailStrategy) Run() {
	all, _ := new(model.StockInfoModel).GetAllByTag("hs300")
	for _, item := range all {
		code := item.Code
		s.RunOne(code)
	}
}

func (s *TailStrategy) RunLine() {

	result := new(TailResult)
	date := "20220101"
	end := "20220106"

	codeArr := []string{}
	all, _ := new(model.StockInfoModel).GetAllByTag("hs300")
	for _, item := range all {
		code := item.Code
		codeArr = append(codeArr, code)
	}
	codeStr := ""
	for _, code := range codeArr {
		codeStr += fmt.Sprintf("'%s',", code)
	}
	codeStr = strings.TrimRight(codeStr, ",")

	for {
		if date > end {
			return
		}
		fmt.Printf("%s begin\n", date)

		sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time from kline where code in(%s) and  date=%s and percent<%d"
		sql = fmt.Sprintf(sql, codeStr, date, -5)
		klines, _ := new(model.KlineModel).Query(sql)
		if len(klines) <= 0 {
			date = s.nextDay(date)
			continue
		}

		for _, kline := range klines {
			nextKlines, _ := new(model.KlineModel).GetByCodeGtDate(kline.Code, date)
			if len(nextKlines) < 0 {
				continue
			}
			nextKline := nextKlines[0]
			if nextKline.Open > kline.Close {
				result.UpNum += 1
			} else {
				result.DownNum += 1
			}
			percent := ((nextKline.Open - kline.Low) / kline.Close) * 100
			result.UpPercent += percent
			fmt.Printf("%s buy %s %f\n", date, kline.Code, percent)
		}
		date = s.nextDay(date)
		fmt.Println(result)
	}

}

func (s *TailStrategy) nextDay(date string) string {
	t, _ := time.Parse("20060102", date)
	return t.Add(time.Hour * 24).Format("20060102")
}

func (s *TailStrategy) RunOne(code string) {
	klines, _ := new(model.KlineModel).GetByCodeRangeDate(code, "20230101", "20240101")

	result := new(TailResult)

	total := len(klines)
	for k, kline := range klines {
		if k+1 > total-1 {
			continue
		}
		if kline.Percent < -5 {
			afterKline := klines[k+1]
			if afterKline.Open > kline.Low {
				result.UpNum += 1
			} else {
				result.DownNum += 1
			}
			percent := ((afterKline.Open - kline.Low) / kline.Close) * 100
			result.UpPercent += percent
		}
	}

	fmt.Printf("%s:upNum %d downNum %d percent:%f\n", code, result.UpNum, result.DownNum, result.UpPercent)
}
