package model

import (
	"fmt"
	"gostock/server"
)

type TradeUserRecord struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	RechargeMoney float64 `json:"recharge_money"`
	HoldMoney     float64 `json:"hold_money"`
	UsableMoney   float64 `json:"usable_money"`
	CTime         int64   `json:"c_time"`
	UTime         int64   `json:"u_time"`
}

type TradeUserModel struct {
}

func (model *TradeUserModel) Insert(record *TradeUserRecord) (int64, error) {
	var lastInsertId int64 = 0

	sql := "INSERT IGNORE INTO `trade_user` (`name`,`recharge_money`, `hold_money`,`usable_money`,`c_time`,`u_time`) VALUES "
	sql += fmt.Sprintf(" ('%s',%f,%f,%f, %d, %d)", record.Name, record.RechargeMoney, record.HoldMoney, record.UsableMoney, record.CTime, record.UTime)
	res, err := server.MysqlClient.Exec(sql)
	if err != nil {
		return lastInsertId, err
	}

	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}
