package model

import (
	"gostock/server"
)

type TradeOrderRecord struct {
	Id     int64   `json:"id"`
	Uid    int64   `json:"uid"`
	Type   int64   `json:"type"`
	Code   string  `json:"code"`
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Number int64   `json:"number"`
	Fee    float64 `json:"fee"`
	Money  float64 `json:"money"`
	CTime  int64   `json:"c_time"`
	UTime  int64   `json:"u_time"`
}

type TradeOrderModel struct {
}

func (model *TradeOrderModel) Insert(record *TradeOrderRecord) (int64, error) {
	var lastInsertId int64 = 0

	sql := "INSERT IGNORE INTO `trade_order` (`uid`,`type`,`code`, `date`,`price`,`number`,`fee`,`money`,`c_time`,`u_time`) VALUES (?,?,?,?,?,?,?,?,?,?) "
	res, err := server.MysqlClient.Exec(sql, record.Uid, record.Type, record.Code, record.Date, record.Price, record.Number, record.Fee, record.Money, record.CTime, record.UTime)
	if err != nil {
		return lastInsertId, err
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}
