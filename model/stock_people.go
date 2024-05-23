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

func (model *StockPeopleModel) GetByCode(code string) ([]*StockPeopleRecord, error) {
	sql := "SELECT id,code,date,holder_num,avg_market,avg_hold_num,c_time,u_time FROM stock_people where code=? order by date desc"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	stockPeopleRecords := []*StockPeopleRecord{}

	for rows.Next() {
		stockPeopleRecord := new(StockPeopleRecord)
		err = rows.Scan(&stockPeopleRecord.Id,
			&stockPeopleRecord.Code,
			&stockPeopleRecord.Date,
			&stockPeopleRecord.HolderNum,
			&stockPeopleRecord.AvgHoldNum,
			&stockPeopleRecord.AvgHoldNum,
			&stockPeopleRecord.CTime,
			&stockPeopleRecord.UTime)
		if err != nil {
			return nil, err
		}
		stockPeopleRecords = append(stockPeopleRecords, stockPeopleRecord)
	}
	return stockPeopleRecords, nil
}
