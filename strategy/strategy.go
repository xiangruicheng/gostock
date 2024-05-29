package strategy

import (
	"fmt"
	"gostock/model"
)

type Strategy interface {
	Run()
}

// hs300:  filter: -4 to -6  buy:close sell:2 day close
// cyb: filter:>-2 buy:close sell:2day open

func Test() {
	all, _ := new(model.StockInfoModel).GetAll()
	for _, item := range all {
		if Feature.IsT(item.Code, "20240529") {
			fmt.Println(item.Code)
		}
	}
}
