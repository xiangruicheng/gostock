package model

import (
	"database/sql"
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
	return model.query(rows)
}

func (model *StockBlockModel) GetByCode(code string) (*StockBlockRecord, error) {
	sql := "SELECT id,type,code,name,c_time,u_time FROM stock_block  where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.queryOne(rows)
}

func (model *StockBlockModel) query(rows *sql.Rows) ([]*StockBlockRecord, error) {
	records := []*StockBlockRecord{}
	for rows.Next() {
		record := new(StockBlockRecord)
		err := rows.Scan(&record.Id,
			&record.Type,
			&record.Code,
			&record.Name,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (model *StockBlockModel) queryOne(rows *sql.Rows) (*StockBlockRecord, error) {
	record := new(StockBlockRecord)
	if rows.Next() {
		err := rows.Scan(&record.Id,
			&record.Type,
			&record.Code,
			&record.Name,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
	}
	return record, nil
}
