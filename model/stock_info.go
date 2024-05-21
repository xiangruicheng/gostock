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
	Cyb    int64  `json:"cyb"`
	Hs300  int64  `json:"hs300"`
	Kcb    int64  `json:"kcb"`
	CTime  int64  `json:"c_time"`
	UTime  int64  `json:"u_time"`
}

type StockInfoModel struct {
}

func (model *StockInfoModel) GetByCode(code string) (*StockInfoRecord, error) {
	sql := "SELECT id,code,name,market,cyb,hs300,kcb,c_time,u_time FROM stock_info where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	stockInfoRecord := new(StockInfoRecord)
	if rows.Next() {
		err = rows.Scan(&stockInfoRecord.Id,
			&stockInfoRecord.Code,
			&stockInfoRecord.Name,
			&stockInfoRecord.Market,
			&stockInfoRecord.Cyb,
			&stockInfoRecord.Hs300,
			&stockInfoRecord.Kcb,
			&stockInfoRecord.CTime,
			&stockInfoRecord.UTime)
		if err != nil {
			return nil, err
		}
	}
	return stockInfoRecord, nil
}

func (model *StockInfoModel) GetAll() ([]*StockInfoRecord, error) {
	sql := "SELECT id,code,name,market,cyb,hs300,kcb,c_time,u_time FROM stock_info "
	rows, err := server.MysqlClient.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	stockInfoRecords := []*StockInfoRecord{}
	for rows.Next() {
		stockInfoRecord := new(StockInfoRecord)
		err = rows.Scan(&stockInfoRecord.Id,
			&stockInfoRecord.Code,
			&stockInfoRecord.Name,
			&stockInfoRecord.Market,
			&stockInfoRecord.Cyb,
			&stockInfoRecord.Hs300,
			&stockInfoRecord.Kcb,
			&stockInfoRecord.CTime,
			&stockInfoRecord.UTime)
		if err != nil {
			return nil, err
		}
		stockInfoRecords = append(stockInfoRecords, stockInfoRecord)
	}
	return stockInfoRecords, nil
}

// Insert insert
func (model *StockInfoModel) Insert(record *StockInfoRecord) (int64, error) {
	var id int64 = 0

	sql := "INSERT IGNORE INTO `stock_info` (`code`, `name`, `market`, `cyb`, `hs300`, `kcb`,  `c_time`, `u_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	res, err := server.MysqlClient.Exec(sql, record.Code, record.Name, record.Market, record.Cyb, record.Hs300, record.Kcb, record.CTime, record.UTime)
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
	sql := "INSERT IGNORE INTO `stock_info` (`code`, `name`, `market`, `cyb`, `hs300`, `kcb`,  `c_time`, `u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" (%s, %s, %s, %d, %d, %d, %d, %d, %d),", record.Code, record.Name, record.Market, record.Cyb, record.Hs300, record.Kcb, record.CTime, record.UTime)
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

func (model *StockInfoModel) UpdateHs300(codeArr []string) (int64, error) {
	var affected int64 = 0
	sql := "UPDATE `stock_info` SET `hs300`=0"
	res, err := server.MysqlClient.Exec(sql)
	if err != nil {
		return affected, err
	}

	codeStr := ""
	for _, code := range codeArr {
		codeStr += "'" + code + "',"
	}
	codeStr = strings.TrimRight(codeStr, ",")
	sql = "UPDATE `stock_info` SET `hs300`=1 WHERE code in(" + codeStr + ");"
	res, err = server.MysqlClient.Exec(sql)
	if err != nil {
		return affected, err
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return affected, err
	}
	return affected, nil
}

func (model *StockInfoModel) UpdateCyb(codeArr []string) (int64, error) {
	var affected int64 = 0
	sql := "UPDATE `stock_info` SET `cyb`=0"
	res, err := server.MysqlClient.Exec(sql)
	if err != nil {
		return affected, err
	}

	codeStr := ""
	for _, code := range codeArr {
		codeStr += "'" + code + "',"
	}
	codeStr = strings.TrimRight(codeStr, ",")
	sql = "UPDATE `stock_info` SET `cyb`=1 WHERE code in(" + codeStr + ");"
	res, err = server.MysqlClient.Exec(sql)
	if err != nil {
		return affected, err
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return affected, err
	}
	return affected, nil
}
