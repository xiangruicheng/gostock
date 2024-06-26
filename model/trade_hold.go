package model

import (
	"database/sql"
	"fmt"
	"gostock/server"
)

type TradeHoldRecord struct {
	Id     int64   `json:"id"`
	Uid    int64   `json:"uid"`
	Code   string  `json:"code"`
	Price  float64 `json:"price"`
	Number int64   `json:"number"`
	CTime  int64   `json:"c_time"`
	UTime  int64   `json:"u_time"`
}

type TradeHoldModel struct {
}

func (model *TradeHoldModel) Insert(record *TradeHoldRecord) (int64, error) {
	var lastInsertId int64 = 0
	fmt.Println(record)
	sql := "INSERT IGNORE INTO `trade_hold` (`uid`,`code`,`price`,`number`,`c_time`,`u_time`) VALUES (?,?,?,?,?,?) "
	res, err := server.MysqlClient.Exec(sql, record.Uid, record.Code, record.Price, record.Number, record.CTime, record.UTime)
	if err != nil {
		return lastInsertId, err
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}

func (model *TradeHoldModel) Update(record *TradeHoldRecord) (int64, error) {
	var lastInsertId int64 = 0

	sql := "UPDATE `trade_hold` set `uid`=?,`code`=?,`price`=?,`number`=?,`u_time`=? where id=? "
	res, err := server.MysqlClient.Exec(sql, record.Uid, record.Code, record.Price, record.Number, record.UTime, record.Id)
	if err != nil {
		return lastInsertId, err
	}

	lastInsertId, err = res.RowsAffected()
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}

func (model *TradeHoldModel) GetByUidAndCode(uid int64, code string) (*TradeHoldRecord, error) {
	sql := "SELECT `id`,`uid`,`code`,`price`,`number`,`c_time`,`u_time` FROM trade_hold where uid=? and code=?"
	rows, err := server.MysqlClient.Query(sql, uid, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.queryOne(rows)
}

func (model *TradeHoldModel) queryOne(rows *sql.Rows) (*TradeHoldRecord, error) {
	record := new(TradeHoldRecord)
	if rows.Next() {
		err := rows.Scan(&record.Id,
			&record.Uid,
			&record.Code,
			&record.Price,
			&record.Number,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
	}
	return record, nil

}
