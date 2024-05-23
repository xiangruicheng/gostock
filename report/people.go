package report

import (
	"fmt"
	"gostock/model"
)

type PeopleReport struct {
}

func (r *PeopleReport) Run() {
	all, _ := new(model.StockInfoModel).GetAllByTag("hs300")
	for _, item := range all {
		stockPeopleRecords, _ := new(model.StockPeopleModel).GetByCode(item.Code)
		if r.isReduce(stockPeopleRecords) {
			fmt.Println(item.Code)
		}
	}

}

// 连续减少
func (r *PeopleReport) isReduce(list []*model.StockPeopleRecord) bool {
	if len(list) < 3 {
		return false
	}
	curRecord := list[0]
	preRecord := list[1]
	pre2Record := list[2]

	rate := -0.02
	rate1 := (float64(curRecord.HolderNum) - float64(preRecord.HolderNum)) / float64(preRecord.HolderNum)
	rate2 := (float64(preRecord.HolderNum) - float64(pre2Record.HolderNum)) / float64(pre2Record.HolderNum)
	if rate1 < rate && rate2 < rate {
		//fmt.Printf("%d %d %d\n", pre2Record.HolderNum, preRecord.HolderNum, curRecord.HolderNum)
		return true
	}
	return false
}
