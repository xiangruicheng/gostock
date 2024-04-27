package model

import (
	"fmt"
	"gostock/server"
	"strings"
)

type StockInfoRecord struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Market string `json:"market"`
	Type   int64  `json:"type"`
	Cyb    int64  `json:"cyb"`
	Hs300  int64  `json:"hs300"`
	Kcb    int64  `json:"kcb"`
	CTime  int64  `json:"c_time"`
	UTime  int64  `json:"u_time"`
}

const (
	StockInfoModel_TypeStock = 1
	StockInfoModel_TypeIndex = 2
	StockInfoModel_TypeEtf   = 3
)

type StockInfoModel struct {
}

// Insert insert
func (model *StockInfoModel) Insert(record *StockInfoRecord) (int64, error) {
	var id int64 = 0

	sql := "INSERT IGNORE INTO `stock_info` (`code`, `name`, `market`, `type`, `cyb`, `hs300`, `kcb`,  `c_time`, `u_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	res, err := server.MysqlClient.Exec(sql, record.Code, record.Name, record.Market, record.Type, record.Cyb, record.Hs300, record.Kcb, record.CTime, record.UTime)
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

// BatchInsert batch insert
func (model *StockInfoModel) BatchInsert(records []*StockInfoRecord) (int64, error) {
	var rowsAffected int64 = 0
	sql := "INSERT IGNORE INTO `stock_info` (`code`, `name`, `market`, `type`, `cyb`, `hs300`, `kcb`,  `c_time`, `u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" (%s, %s, %s, %d, %d, %d, %d, %d, %d),", record.Code, record.Name, record.Market, record.Type, record.Cyb, record.Hs300, record.Kcb, record.CTime, record.UTime)
	}
	sql = strings.TrimRight(sql, ",")
	fmt.Println(sql)
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
