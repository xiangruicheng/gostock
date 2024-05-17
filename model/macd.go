package model

import (
	"fmt"
	"gostock/server"
	"strings"
)

type MacdRecord struct {
	Id    int64   `json:"id"`
	Code  string  `json:"code"`
	Date  string  `json:"date"`
	Close float64 `json:"close"`
	Ema12 float64 `json:"ema12"`
	Ema26 float64 `json:"ema26"`
	Diff  float64 `json:"diff"`
	Dea   float64 `json:"dea"`
	Macd  float64 `json:"macd"`
	CTime int64   `json:"c_time"`
	UTime int64   `json:"u_time"`
}

type MacdModel struct {
}

func (model *MacdModel) BatchInsert(records []*MacdRecord) (int64, error) {
	var rowsAffected int64 = 0

	sql := "INSERT IGNORE INTO `macd` (`code`, `date`,`close`, `ema12`, `ema26`, `diff`, `dea`, `macd`, `c_time`, `u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" ('%s', '%s', %f, %f, %f, %f, %f, %f,  %d, %d),", record.Code, record.Date, record.Close, record.Ema12, record.Ema26, record.Diff, record.Dea, record.Macd, record.CTime, record.UTime)
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

func (model *MacdModel) GetByCode(code string) ([]*MacdRecord, error) {
	sql := "SELECT id,code,date,close,ema12,ema26,diff,dea,macd,c_time,u_time FROM macd where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	macdRecords := []*MacdRecord{}

	for rows.Next() {
		macdRecord := new(MacdRecord)
		err = rows.Scan(&macdRecord.Id,
			&macdRecord.Code,
			&macdRecord.Date,
			&macdRecord.Close,
			&macdRecord.Ema12,
			&macdRecord.Ema26,
			&macdRecord.Diff,
			&macdRecord.Dea,
			&macdRecord.Macd,
			&macdRecord.CTime,
			&macdRecord.UTime)
		if err != nil {
			return nil, err
		}
		macdRecords = append(macdRecords, macdRecord)
	}
	return macdRecords, nil
}
