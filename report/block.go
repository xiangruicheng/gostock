package report

import (
	"fmt"
	"gostock/model"
)

type BlockReportStruct struct {
}

var BlockReport *BlockReportStruct

func init() {
	BlockReport = new(BlockReportStruct)
}

func (r *BlockReportStruct) Run() {
	bkCode := "BK0475"
	records, _ := new(model.StockBlockCodeModel).GetByBkCode(bkCode)
	codeArr := []string{}
	for _, record := range records {
		codeArr = append(codeArr, record.Code)
	}
	quoteRecords, _ := new(model.StockQuoteModel).GetByCodes(codeArr, "order by pb asc")
	for _, quote := range quoteRecords {
		fmt.Printf("%s|%f\n", quote.Name, quote.Pb)
	}
}
