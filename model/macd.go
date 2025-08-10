package model

import (
	"gostock/server"
)

type Macd struct {
	Id    int64   `json:"id" xorm:"id"`
	Code  string  `json:"code" xorm:"code"`
	Date  string  `json:"date" xorm:"date"`
	Close float64 `json:"close" xorm:"close"`
	Ema12 float64 `json:"ema12" xorm:"ema12"`
	Ema26 float64 `json:"ema26" xorm:"ema26"`
	Diff  float64 `json:"diff" xorm:"diff"`
	Dea   float64 `json:"dea" xorm:"dea"`
	Macd  float64 `json:"macd" xorm:"macd"`
	CTime int64   `json:"c_time" xorm:"c_time"`
	UTime int64   `json:"u_time" xorm:"u_time"`
}

type MacdModel struct {
}

func (model *MacdModel) BatchInsert(records []*Macd) (int64, error) {
	return server.MysqlEngine.Insert(records)
}

func (model *MacdModel) GetByCode(code string) ([]*Macd, error) {
	macds := []*Macd{}
	err := server.MysqlEngine.Where("code=?", code).Find(&macds)
	return macds, err
}

func (model *MacdModel) GetByCodeAndDate(code string, date string) (*Macd, error) {
	macd := new(Macd)
	_, err := server.MysqlEngine.Where(" code=? and date=?", code, date).Get(&macd)
	return macd, err
}

func (model *MacdModel) GetByCodeERangeDate(code string, min string, max string) ([]*Macd, error) {
	macds := []*Macd{}
	err := server.MysqlEngine.Where("code=? and date>=? and date<=?", min, max).OrderBy("date asc").Find(&macds)
	return macds, err
}
