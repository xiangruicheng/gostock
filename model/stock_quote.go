package model

import (
	"database/sql"
	"fmt"
	"gostock/server"
	"strings"
)

type StockQuoteRecord struct {
	Id                 int32
	Code               string
	Name               string
	PeForecast         float64
	PeTtm              float64
	PeLyr              float64
	Pb                 float64
	TotalShares        float64
	FloatShares        float64
	FloatMarketCapital float64
	MarketCapital      float64
	Amount             float64
	Volume             float64
	TurnoverRate       float64
	Amplitude          float64
	Navps              float64
	Eps                float64
	VolumeRatio        float64
	PankouRatio        float64
	High               float64
	Low                float64
	Open               float64
	Current            float64
	Dividend           float64
	DividendYield      float64
	Date               string
	CTime              int64
	UTime              int64
}

type StockQuoteModel struct {
}

func (dao *StockQuoteModel) Insert(model *StockQuoteRecord) (int64, error) {
	var id int64 = 0
	sql := "INSERT INTO `stock_quote` ( `code`, `name`, `pe_forecast`, `pe_ttm`, `pe_lyr`, `pb`, `total_shares`, `float_shares`, `float_market_capital`, `market_capital`, `amount`, `volume`, `turnover_rate`, `amplitude`, `navps`, `eps`, `volume_ratio`, `pankou_ratio`, `high`, `low`, `open`, `current`, `dividend`, `dividend_yield`, `date`, `c_time`, `u_time`) VALUES ( ?, ?, ?, ?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	res, err := server.MysqlClient.Exec(sql, model.Code, model.Name, model.PeForecast, model.PeTtm, model.PeLyr, model.Pb, model.TotalShares, model.FloatShares, model.FloatMarketCapital, model.MarketCapital, model.Amount, model.Volume, model.TurnoverRate, model.Amplitude, model.Navps, model.Eps, model.VolumeRatio, model.PankouRatio, model.High, model.Low, model.Open, model.Current, model.Dividend, model.DividendYield, model.Date, model.CTime, model.UTime)
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}

func (dao *StockQuoteModel) Update(model *StockQuoteRecord) (int64, error) {
	var id int64 = 0
	sql := "UPDATE `stock_quote` set `code`=?, `name`=?, `pe_forecast`=?, `pe_ttm`=?, `pe_lyr`=?, `pb`=?, `total_shares`=?, `float_shares`=?, `float_market_capital`=?, `market_capital`=?, `amount`=?, `volume`=?, `turnover_rate`=?, `amplitude`=?, `navps`=?, `eps`=?, `volume_ratio`=?, `pankou_ratio`=?, `high`=?, `low`=?, `open`=?, `current`=?, `dividend`=?, `dividend_yield`=?, `date`=?, `u_time`=? where code=?"
	res, err := server.MysqlClient.Exec(sql, model.Code, model.Name, model.PeForecast, model.PeTtm, model.PeLyr, model.Pb, model.TotalShares, model.FloatShares, model.FloatMarketCapital, model.MarketCapital, model.Amount, model.Volume, model.TurnoverRate, model.Amplitude, model.Navps, model.Eps, model.VolumeRatio, model.PankouRatio, model.High, model.Low, model.Open, model.Current, model.Dividend, model.DividendYield, model.Date, model.UTime, model.Code)

	if err != nil {
		return id, err
	}

	id, err = res.RowsAffected()
	if err != nil {
		return id, err
	}

	return id, nil
}

func (dao *StockQuoteModel) IsExist(code string) (bool, error) {
	sql := "SELECT * FROM stock_quote where code=?"
	rows, _ := server.MysqlClient.Query(sql, code)
	if rows.Next() { //如果有值为真就是有数据
		rows.Close()
		return true, nil
	}
	return false, nil
}

func (model *StockQuoteModel) GetByCodes(codeArr []string, other string) ([]*StockQuoteRecord, error) {
	codeStr := ""
	for _, code := range codeArr {
		codeStr += fmt.Sprintf("'%s',", code)
	}
	codeStr = strings.TrimRight(codeStr, ",")
	sql := fmt.Sprintf("SELECT id,code,name,pe_forecast,pe_ttm,pe_lyr,pb,total_shares,float_shares,float_market_capital ,market_capital,amount,volume,turnover_rate,amplitude,navps,eps,volume_ratio,pankou_ratio,high,low,open,current,dividend,dividend_yield,date,c_time,u_time FROM stock_quote where code in(%s) %s", codeStr, other)
	rows, err := server.MysqlClient.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.query(rows)
}

func (model *StockQuoteModel) GetByCode(code string) (*StockQuoteRecord, error) {
	sql := fmt.Sprintf("SELECT id,code,name,pe_forecast,pe_ttm,pe_lyr,pb,total_shares,float_shares,float_market_capital ,market_capital,amount,volume,turnover_rate,amplitude,navps,eps,volume_ratio,pankou_ratio,high,low,open,current,dividend,dividend_yield,date,c_time,u_time FROM stock_quote where code=%s", code)
	rows, err := server.MysqlClient.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	return model.queryOne(rows)
}

func (model *StockQuoteModel) query(rows *sql.Rows) ([]*StockQuoteRecord, error) {
	records := []*StockQuoteRecord{}
	for rows.Next() {
		record := new(StockQuoteRecord)
		err := rows.Scan(&record.Id,
			&record.Code,
			&record.Name,
			&record.PeForecast,
			&record.PeTtm,
			&record.PeLyr,
			&record.Pb,
			&record.TotalShares,
			&record.FloatShares,
			&record.FloatMarketCapital,
			&record.MarketCapital,
			&record.Amount,
			&record.Volume,
			&record.TurnoverRate,
			&record.Amplitude,
			&record.Navps,
			&record.Eps,
			&record.VolumeRatio,
			&record.PankouRatio,
			&record.High,
			&record.Low,
			&record.Open,
			&record.Current,
			&record.Dividend,
			&record.DividendYield,
			&record.Date,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (model *StockQuoteModel) queryOne(rows *sql.Rows) (*StockQuoteRecord, error) {
	record := new(StockQuoteRecord)
	if rows.Next() {
		err := rows.Scan(&record.Id,
			&record.Code,
			&record.Name,
			&record.PeForecast,
			&record.PeTtm,
			&record.PeLyr,
			&record.Pb,
			&record.TotalShares,
			&record.FloatShares,
			&record.FloatMarketCapital,
			&record.MarketCapital,
			&record.Amount,
			&record.Volume,
			&record.TurnoverRate,
			&record.Amplitude,
			&record.Navps,
			&record.Eps,
			&record.VolumeRatio,
			&record.PankouRatio,
			&record.High,
			&record.Low,
			&record.Open,
			&record.Current,
			&record.Dividend,
			&record.DividendYield,
			&record.Date,
			&record.CTime,
			&record.UTime)
		if err != nil {
			return nil, err
		}
	}
	return record, nil
}
