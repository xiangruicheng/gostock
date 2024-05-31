package model

import (
	"database/sql"
	"fmt"
	"gostock/server"
	"strings"
)

type KlineRecord struct {
	Id      int64   `json:"id"`
	Code    string  `json:"code"`
	Date    string  `json:"date"`
	Amount  float64 `json:"amount"`
	Volume  float64 `json:"volume"`
	Open    float64 `json:"open"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Close   float64 `json:"close"`
	Chg     float64 `json:"chg"`
	Percent float64 `json:"percent"`
	CTime   int64   `json:"c_time"`
	UTime   int64   `json:"u_time"`
}

type KlineModel struct {
}

// GetByCode Get by code
func (model *KlineModel) GetByCode(code string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

// GetByCodeAndDate
func (model *KlineModel) GetByCodeAndDate(code string, date string) (*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date=?"
	rows, err := server.MysqlClient.Query(sql, code, date)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.queryOne(rows)
}

// Insert insert
func (model *KlineModel) Insert(record *KlineRecord) (int64, error) {
	var id int64 = 0

	sql := "INSERT IGNORE INTO `kline` (`code`, `date`, `volume`, `amount`, `open`, `high`, `low`, `close`, `chg`, `percent`, `c_time`, `u_time`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	res, err := server.MysqlClient.Exec(sql, record.Code, record.Date, record.Volume, record.Amount, record.Open, record.High, record.Low, record.Close, record.Chg, record.Percent, record.CTime, record.UTime)
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
func (model *KlineModel) BatchInsert(records []*KlineRecord) (int64, error) {
	var rowsAffected int64 = 0

	sql := "INSERT IGNORE INTO `kline` (`code`, `date`, `volume`, `amount`, `open`, `high`, `low`, `close`, `chg`, `percent`, `c_time`, `u_time`) VALUES "
	for _, record := range records {
		sql += fmt.Sprintf(" ('%s', '%s', %f, %f, %f, %f, %f, %f, %f, %f, %d, %d),", record.Code, record.Date, record.Volume, record.Amount, record.Open, record.High, record.Low, record.Close, record.Chg, record.Percent, record.CTime, record.UTime)
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

func (model *KlineModel) GetByCodeRangeDate(code string, min string, max string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date>? and date<?"
	rows, err := server.MysqlClient.Query(sql, code, min, max)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *KlineModel) GetByCodeERangeDate(code string, min string, max string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date>=? and date<=? order by date asc"
	rows, err := server.MysqlClient.Query(sql, code, min, max)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *KlineModel) GetByCodeGtDate(code string, min string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date>?"
	rows, err := server.MysqlClient.Query(sql, code, min)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *KlineModel) GetByCodeLtDate(code string, min string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date>?"
	rows, err := server.MysqlClient.Query(sql, code, min)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *KlineModel) query(rows *sql.Rows) ([]*KlineRecord, error) {
	klineRecords := []*KlineRecord{}
	for rows.Next() {
		klineRecord := new(KlineRecord)
		err := rows.Scan(&klineRecord.Id,
			&klineRecord.Code,
			&klineRecord.Date,
			&klineRecord.Volume,
			&klineRecord.Open,
			&klineRecord.High,
			&klineRecord.Low,
			&klineRecord.Close,
			&klineRecord.Chg,
			&klineRecord.Percent,
			&klineRecord.CTime,
			&klineRecord.UTime)
		if err != nil {
			return nil, err
		}
		klineRecords = append(klineRecords, klineRecord)
	}
	return klineRecords, nil
}

func (model *KlineModel) queryOne(rows *sql.Rows) (*KlineRecord, error) {
	klineRecord := new(KlineRecord)
	if rows.Next() {
		err := rows.Scan(&klineRecord.Id,
			&klineRecord.Code,
			&klineRecord.Date,
			&klineRecord.Volume,
			&klineRecord.Open,
			&klineRecord.High,
			&klineRecord.Low,
			&klineRecord.Close,
			&klineRecord.Chg,
			&klineRecord.Percent,
			&klineRecord.CTime,
			&klineRecord.UTime)
		if err != nil {
			return nil, err
		}
	}
	return klineRecord, nil
}
