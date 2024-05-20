package model

import "gostock/server"

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
	sql := "INSERT INTO `gp_base_info` ( `code`, `name`, `pe_forecast`, `pe_ttm`, `pe_lyr`, `pb`, `total_shares`, `float_shares`, `float_market_capital`, `market_capital`, `amount`, `volume`, `turnover_rate`, `amplitude`, `navps`, `eps`, `volume_ratio`, `pankou_ratio`, `high`, `low`, `open`, `current`, `dividend`, `dividend_yield`, `date`, `c_time`, `u_time`) VALUES ( ?, ?, ?, ?,?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
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
	sql := "UPDATE `gp_base_info` set `code`=?, `name`=?, `pe_forecast`=?, `pe_ttm`=?, `pe_lyr`=?, `pb`=?, `total_shares`=?, `float_shares`=?, `float_market_capital`=?, `market_capital`=?, `amount`=?, `volume`=?, `turnover_rate`=?, `amplitude`=?, `navps`=?, `eps`=?, `volume_ratio`=?, `pankou_ratio`=?, `high`=?, `low`=?, `open`=?, `current`=?, `dividend`=?, `dividend_yield`=?, `date`=?, `u_time`=? where code=?"
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
	sql := "SELECT * FROM gp_base_info where code=?"
	rows, _ := server.MysqlClient.Query(sql, code)
	if rows.Next() { //如果有值为真就是有数据
		rows.Close()
		return true, nil
	}
	return false, nil
}
