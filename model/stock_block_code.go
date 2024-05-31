package model

import (
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
