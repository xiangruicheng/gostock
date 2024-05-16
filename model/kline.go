package model

import (
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
	klineRecords := []*KlineRecord{}

	for rows.Next() {
		klineRecord := new(KlineRecord)
		err = rows.Scan(&klineRecord.Id,
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

func (model *KlineModel) GetByTypeCodeDate(code string, min string, max string) ([]*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? and date>? and date<?"
	rows, err := server.MysqlClient.Query(sql, code, min, max)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	klineRecords := []*KlineRecord{}

	for rows.Next() {
		klineRecord := new(KlineRecord)
		err = rows.Scan(&klineRecord.Id,
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

func (model *KlineModel) GetIndexLast(code string) (*KlineRecord, error) {
	sql := "SELECT id,code,date,volume,open,high,low,close,chg,percent,c_time,u_time FROM kline where code=? order by `date` desc limit 1"
	rows, err := server.MysqlClient.Query(sql, code)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	klineRecord := new(KlineRecord)

	if rows.Next() {
		err = rows.Scan(&klineRecord.Id,
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
