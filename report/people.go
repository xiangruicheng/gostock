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
		if r.isReduce(stockPeopleRecords, 3, -5) {
			fmt.Println(item.Code, item.Name)
		}
	}

}

// isReduce
// example isReduce(list,3,-2)
func (r *PeopleReport) isReduce(list []*model.StockPeopleRecord, num int64, rate float64) bool {
	if len(list) < int(num) {
		return false
	}
	var tag bool = true
	var counter int64 = 0
	for k, item := range list {
		counter += 1
		if counter > num {
			break
		}
		preItem := list[k+1]
		realRate := ((float64(item.HolderNum) - float64(preItem.HolderNum)) / float64(preItem.HolderNum)) * 100
		if realRate > rate {
			tag = false
			break
		}
	}
	return tag
}
