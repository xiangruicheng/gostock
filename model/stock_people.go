package model

import (
	"fmt"
	"gostock/server"
	"strings"
)

type StockPeopleRecord struct {
	Id         int32
	Code       string
	Date       string
	HolderNum  int64
	AvgMarket  float64
	AvgHoldNum float64
	CTime      int64
	UTime      int64
}

type StockPeopleModel struct {
}

// BatchInsert batch insert
func (model *StockPeopleModel) BatchInsert(records []*StockPeopleRecord) (int64, error) {
	var rowsAffected int64 = 0

	sql := "INSERT IGNORE INTO `stock_people` (`code`, `date`,`holder_num`,`avg_market`, `avg_hold_num`,`c_time`,`u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" ('%s', '%s', %d, '%f','%f', %d, %d),", record.Code, record.Date, record.HolderNum, record.AvgMarket, record.AvgHoldNum, record.CTime, record.UTime)
	}
	sql = strings.TrimRight(sql, ",")

	res, err := server.MysqlClient.Exec(sql)
	if err != nil {
		return rowsAffected, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}
	return rowsAffected, nil
}
