package report

import (
	"fmt"
	"gostock/model"
)

type stockReportStruct struct {
}

var StockReport *stockReportStruct

func init() {
	StockReport = new(stockReportStruct)
}

func (r stockReportStruct) Run(code string) {

	stockQuote, _ := new(model.StockQuoteModel).GetByCode(code)
	fmt.Printf("|%10s|%10s|%10s|\n", "NAME", "PE", "PB")
	fmt.Printf("|%10s|%10s|%10s|\n", "----------", "----------", "----------")
	fmt.Printf("|%-10s|%10f|%10f|\n", stockQuote.Name, stockQuote.PeTtm, stockQuote.Pb)

	blocks, _ := new(model.StockBlockCodeModel).GetByCode(code)
	for _, block := range blocks {
		stockBlock, _ := new(model.StockBlockModel).GetByCode(block.BkCode)
		if stockBlock.Type != 1 {
			continue
		}
		_, avgPe, avgPb := r.blockAvg(stockBlock.Code)
		fmt.Printf("|%-10s|%10f|%10f|\n", stockBlock.Name, avgPe, avgPb)

	}
}

func (r stockReportStruct) blockAvg(bkCode string) (int, float64, float64) {
	records, _ := new(model.StockBlockCodeModel).GetByBkCode(bkCode)
	codeArr := []string{}
	for _, record := range records {
		codeArr = append(codeArr, record.Code)
	}

	var num int
	var avgPe, avgPb, totalPe, totalPb float64
	quoteArr, _ := new(model.StockQuoteModel).GetByCodes(codeArr, "")
	for _, quote := range quoteArr {
		num += 1
		totalPe += quote.PeTtm
		totalPb += quote.Pb
	}
	avgPe = totalPe / float64(num)
	avgPb = totalPb / float64(num)
	return num, avgPe, avgPb
}
