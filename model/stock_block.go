package model

import (
	"fmt"
	"gostock/server"
	"strings"
)

type StockBlockRecord struct {
	Id    int32
	Type  int64
	Code  string
	Name  string
	CTime int64
	UTime int64
}

type StockBlockModel struct {
}

// BatchInsert batch insert
func (model *StockBlockModel) BatchInsert(records []*StockBlockRecord) (int64, error) {
	var rowsAffected int64 = 0

	sql := "INSERT IGNORE INTO `stock_block` (`type`,`code`, `name`,`c_time`,`u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" (%d,'%s', '%s', %d, %d),", record.Type, record.Code, record.Name, record.CTime, record.UTime)
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

func (model *StockBlockModel) GetAll() ([]*StockBlockRecord, error) {
	sql := "SELECT id,type,code,name,c_time,u_time FROM stock_block "
	rows, err := server.MysqlClient.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	stockblcokRecords := []*StockBlockRecord{}
	for rows.Next() {
		stockblcokRecord := new(StockBlockRecord)
		err = rows.Scan(&stockblcokRecord.Id,
			&stockblcokRecord.Type,
			&stockblcokRecord.Code,
			&stockblcokRecord.Name,
			&stockblcokRecord.CTime,
			&stockblcokRecord.UTime)
		if err != nil {
			return nil, err
		}
		stockblcokRecords = append(stockblcokRecords, stockblcokRecord)
	}
	return stockblcokRecords, nil
}
