package model

import (
	"database/sql"
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

func (model *TradeUserModel) Update(record *TradeUserRecord) (int64, error) {
	var lastInsertId int64 = 0

	sql := "UPDATE `trade_user` set `name`=?,`recharge_money`=?,`hold_money`=?,`usable_money`=?,`c_time`=?,`u_time`=? where id=? "
	res, err := server.MysqlClient.Exec(sql, record.Name, record.RechargeMoney, record.HoldMoney, record.UsableMoney, record.CTime, record.UTime, record.Id)
	if err != nil {
		return lastInsertId, err
	}

	lastInsertId, err = res.RowsAffected()
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}

func (model *TradeUserModel) GetById(id int64) (*TradeUserRecord, error) {
	sql := "SELECT `id`,`name`,`recharge_money`,`hold_money`,`usable_money`,`c_time`,`u_time` FROM trade_user where id=?"
	rows, err := server.MysqlClient.Query(sql, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.queryOne(rows)
}

func (model *TradeUserModel) queryOne(rows *sql.Rows) (*TradeUserRecord, error) {
	record := new(TradeUserRecord)
	if rows.Next() {
		err := rows.Scan(&record.Id,
			&record.Name,
			&record.RechargeMoney,
			&record.HoldMoney,
			&record.UsableMoney,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
	}
	return record, nil

}
