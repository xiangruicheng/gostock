package model

import (
	"fmt"
	"gostock/server"
	"strings"
)

type KlineRecord struct {
	Id      int     `json:"id"`
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

// GetById get by id
func (model *KlineModel) GetById(id int) (*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where id=?"
	row, err := server.MysqlClient.Query(sql, id)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	quotationRecord := new(KlineRecord)
	if row.Next() {
		err = row.Scan(&quotationRecord.Id,
			&quotationRecord.Code,
			&quotationRecord.Date,
			&quotationRecord.Volume,
			&quotationRecord.Open,
			&quotationRecord.High,
			&quotationRecord.Low,
			&quotationRecord.Close,
			&quotationRecord.Chg,
			&quotationRecord.Percent,
			&quotationRecord.CTime,
			&quotationRecord.UTime)
		if err != nil {
			return nil, err
		}
	}
	return quotationRecord, nil
}

// GetByCode Get by code
func (model *KlineModel) GetByCode(code string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=?"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	quotationRecords := []*KlineRecord{}

	for rows.Next() {
		quotationRecord := new(KlineRecord)
		err = rows.Scan(&quotationRecord.Id,
			&quotationRecord.Code,
			&quotationRecord.Date,
			&quotationRecord.Volume,
			&quotationRecord.Open,
			&quotationRecord.High,
			&quotationRecord.Low,
			&quotationRecord.Close,
			&quotationRecord.Chg,
			&quotationRecord.Percent,
			&quotationRecord.CTime,
			&quotationRecord.UTime)
		if err != nil {
			return nil, err
		}
		quotationRecords = append(quotationRecords, quotationRecord)
	}
	return quotationRecords, nil
}

// Insert insert
func (model *KlineModel) Insert(record *KlineRecord) (int64, error) {
	var id int64 = 0

	sql := "INSERT IGNORE INTO `kline` (`code`, `date`, `volume`, `amount`, `open`, `high`, `low`, `close`, `chg`, `percent`, `c_time`, `u_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
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
		sql += fmt.Sprintf(" (%s, %s, %f, %f, %f, %f, %f, %f, %f, %f, %d, %d),", record.Code, record.Date, record.Volume, record.Amount, record.Open, record.High, record.Low, record.Close, record.Chg, record.Percent, record.CTime, record.UTime)
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
