package strategy

import "fmt"

type DcfStruct struct {
}

var Dcf *DcfStruct

func init() {
	Dcf = new(DcfStruct)
}

func (s *DcfStruct) Demo() {
	var profit float64 = 1.4
	var gr float64 = 0.05
	var tenAfterGr float64 = 0.02
	var r float64 = 0.1
	s.Dcf(profit, gr, tenAfterGr, r)
}

// 现金流量折现法(Discounted Cash Flow Technique，简称DCF
// profit 利润
// gr 增长率
// tenAfterGr 十年后增长率(永续增长率)
// r 折现率
func (s *DcfStruct) Dcf(profit float64, gr float64, tenAfterGr float64, r float64) {
	// 总和
	var totalValue float64 = 0
	for n := 1; n <= 11; n++ {
		var moneyFlowR float64
		if n <= 10 {
			powGr := s.pow(gr, n)
			powR := s.pow(r, n)
			//现金流
			moneyFlow := profit * powGr
			//现金流折现
			moneyFlowR = moneyFlow / powR
			fmt.Printf("n=%d %f\n", n, moneyFlowR)
		}
		if n == 11 {
			powGr := s.pow(gr, 10)
			powR := s.pow(r, 11)
			//现金流
			moneyFlow := profit * powGr * (1 + tenAfterGr) / (r - tenAfterGr)
			//现金流折现
			moneyFlowR = moneyFlow / powR
			fmt.Printf("n=%d %f\n", n, moneyFlowR)
		}
		totalValue += moneyFlowR
	}
	fmt.Printf("taotalValue %f\n", totalValue)
}

func (s *DcfStruct) pow(v float64, n int) float64 {
	var result float64 = 1
	for i := 1; i <= n; i++ {
		result = result * (1 + v)
	}
	return result
}
