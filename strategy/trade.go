package strategy

import "fmt"

type Trade struct {
	Code    string         //股票电视
	IsHold  bool           //是否有持仓
	Records []*TradeRecord //交易记录
}

const (
	TradeTypeBuy = iota + 1
	TradeTypeAdd
	TradeTypeSell
	TradeTypeStop
)

type TradeRecord struct {
	Code  string  //股票代码
	Date  string  //交易日期
	Type  int     //交易类型
	Num   int64   //数量
	Price float64 //平均价格
	Other string  //其他参数
}

// 建仓
func (t *Trade) Buy(date string, price float64, num int64, other string) {
	t.Records = append(t.Records, &TradeRecord{
		Code:  t.Code,
		Date:  date,
		Type:  TradeTypeBuy,
		Num:   num,
		Price: price,
		Other: other,
	})
	if t.GetTotalNum() > 0 {
		t.IsHold = true
	} else {
		t.IsHold = false
	}
}

// 加仓
func (t *Trade) Add(date string, price float64, num int64, other string) {
	t.Records = append(t.Records, &TradeRecord{
		Code:  t.Code,
		Date:  date,
		Type:  TradeTypeAdd,
		Num:   num,
		Price: price,
		Other: other,
	})
	if t.GetTotalNum() > 0 {
		t.IsHold = true
	} else {
		t.IsHold = false
	}
}

// 卖出
func (t *Trade) Sell(date string, price float64, num int64, other string) {
	t.Records = append(t.Records, &TradeRecord{
		Code:  t.Code,
		Date:  date,
		Type:  TradeTypeSell,
		Num:   num,
		Price: price,
		Other: other,
	})
	if t.GetTotalNum() > 0 {
		t.IsHold = true
	} else {
		t.IsHold = false
	}
}

// 止损
func (t *Trade) Stop(date string, price float64, num int64, other string) {
	t.Records = append(t.Records, &TradeRecord{
		Code:  t.Code,
		Date:  date,
		Type:  TradeTypeStop,
		Num:   num,
		Price: price,
		Other: other,
	})
	if t.GetTotalNum() > 0 {
		t.IsHold = true
	} else {
		t.IsHold = false
	}
}

func (t *Trade) Report() {
	tradeTypeMap := map[int]string{
		1: "买入",
		2: "加仓",
		3: "卖出",
		4: "止损",
	}
	var totalNum int64
	var totalEarn float64
	for k, record := range t.Records {
		if record.Type == 1 || record.Type == 2 {
			totalNum += record.Num
		} else {
			totalNum -= record.Num
		}
		var earn float64
		if totalNum == 0 {
			earn = float64(record.Num) * (record.Price - t.Records[k-1].Price)
		}

		totalEarn += earn
		fmt.Printf("date=%s type=%s price=%f num=%d other=%s\n", record.Date, tradeTypeMap[record.Type], record.Price, record.Num, record.Other)
		if earn != 0 {
			fmt.Printf("收益=%f\n\n", earn)
		}
	}
	fmt.Printf("总收益=%f\n\n", totalEarn)

}

func (t *Trade) GetTotalNum() int64 {
	var num int64
	for _, r := range t.Records {
		if r.Type == TradeTypeBuy || r.Type == TradeTypeAdd {
			num += r.Num
		}
		if r.Type == TradeTypeSell || r.Type == TradeTypeStop {
			num -= r.Num
		}
	}
	return num
}

func (t *Trade) GetLastRecord() *TradeRecord {
	len := len(t.Records)
	if len <= 0 {
		return nil
	}
	return t.Records[len-1]
}
