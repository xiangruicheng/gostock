package model

import (
	"gostock/server"
)

type Kline struct {
	Id      int64   `json:"id" xorm:"id"`
	Code    string  `json:"code" xorm:"code"`
	Date    string  `json:"date" xorm:"date"`
	Amount  float64 `json:"amount" xorm:"amount"`
	Volume  float64 `json:"volume" xorm:"volume"`
	Open    float64 `json:"open" xorm:"open"`
	High    float64 `json:"high" xorm:"high"`
	Low     float64 `json:"low" xorm:"low"`
	Close   float64 `json:"close" xorm:"close"`
	Chg     float64 `json:"chg" xorm:"chg"`
	Percent float64 `json:"percent" xorm:"percent"`
	CTime   int64   `json:"c_time" xorm:"c_time"`
	UTime   int64   `json:"u_time" xorm:"u_time"`
}

type KlineModel struct {
}

// GetByCode Get by code
func (model *KlineModel) GetByCode(code string) ([]*Kline, error) {
	klines := []*Kline{}
	err := server.MysqlEngine.Where("code=?", code).Find(&klines)
	return klines, err
}

// GetByCodeAndDate
func (model *KlineModel) GetByCodeAndDate(code string, date string) (*Kline, error) {
	kline := new(Kline)
	_, err := server.MysqlEngine.Where("code=? and date=?", code, date).Get(kline)
	return kline, err
}

// Insert insert
func (model *KlineModel) Insert(record *Kline) (int64, error) {
	return server.MysqlEngine.Insert(record)
}

// BatchInsert batch insert
func (model *KlineModel) BatchInsert(records []*Kline) (int64, error) {
	return server.MysqlEngine.Insert(records)
}

func (model *KlineModel) GetByCodeRangeDate(code string, min string, max string) ([]*Kline, error) {
	klines := []*Kline{}
	err := server.MysqlEngine.Where("code=? and date>? and date<?", code, min, max).Find(&klines)
	return klines, err
}

func (model *KlineModel) GetByCodeERangeDate(code string, min string, max string) ([]*Kline, error) {
	klines := []*Kline{}
	err := server.MysqlEngine.Where("code=? and date>=? and date<=?", code, min, max).OrderBy("date asc").Find(&klines)
	return klines, err
}

func (model *KlineModel) GetByCodeGtDate(code string, min string) ([]*Kline, error) {
	klines := []*Kline{}
	err := server.MysqlEngine.Where("code=? and date>?", code, min).Find(&klines)
	return klines, err
}
