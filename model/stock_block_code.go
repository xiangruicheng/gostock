package model

import (
	"database/sql"
	"fmt"
	"gostock/server"
	"strings"
)

type StockBlockCodeRecord struct {
	Id     int32
	BkCode string
	Code   string
	CTime  int64
	UTime  int64
}

type StockBlockCodeModel struct {
}

// BatchInsert batch insert
func (model *StockBlockCodeModel) BatchInsert(records []*StockBlockCodeRecord) (int64, error) {
	var rowsAffected int64 = 0

	sql := "INSERT IGNORE INTO `stock_block_code` (`bk_code`,`code`,`c_time`,`u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" ('%s', '%s', %d, %d),", record.BkCode, record.Code, record.CTime, record.UTime)
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

func (model *StockBlockCodeModel) GetByBkCode(bkCode string) ([]*StockBlockCodeRecord, error) {
	sql := "SELECT id,bk_code,code,c_time,u_time FROM stock_block_code where bk_code=?"
	rows, err := server.MysqlClient.Query(sql, bkCode)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *StockBlockCodeModel) GetByCode(code string) ([]*StockBlockCodeRecord, error) {
	sql := "SELECT id,bk_code,code,c_time,u_time FROM stock_block_code where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *StockBlockCodeModel) query(rows *sql.Rows) ([]*StockBlockCodeRecord, error) {
	records := []*StockBlockCodeRecord{}
	for rows.Next() {
		record := new(StockBlockCodeRecord)
		err := rows.Scan(&record.Id,
			&record.BkCode,
			&record.Code,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (model *StockBlockCodeModel) queryOne(rows *sql.Rows) (*StockBlockCodeRecord, error) {
	record := new(StockBlockCodeRecord)
	if rows.Next() {
		err := rows.Scan(&record.Id,
			&record.BkCode,
			&record.Code,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
	}
	return record, nil
}
