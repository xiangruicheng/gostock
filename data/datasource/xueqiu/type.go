package xueqiu

type KlineListResponse struct {
	Data struct {
		Symbol string      `json:"symbol"`
		Column []string    `json:"column"`
		Item   [][]float64 `json:"item"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type StockAllResponse struct {
	Data struct {
		Count int `json:"count"`
		List  []struct {
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
		} `json:"list"`
	} `json:"data"`
}

type QuoteResponse struct {
	Data struct {
		Market struct {
			StatusId     int         `json:"status_id"`
			Region       string      `json:"region"`
			Status       string      `json:"status"`
			TimeZone     string      `json:"time_zone"`
			TimeZoneDesc interface{} `json:"time_zone_desc"`
			DelayTag     int         `json:"delay_tag"`
		} `json:"market"`
		Quote struct {
			Bol                string      `json:"bol"`
			Code               string      `json:"code"`
			AccUnitNav         float64     `json:"acc_unit_nav"`
			High52W            float64     `json:"high52w"`
			NavDate            int64       `json:"nav_date"`
			AvgPrice           float64     `json:"avg_price"`
			Delayed            int         `json:"delayed"`
			Type               int         `json:"type"`
			ExpirationDate     interface{} `json:"expiration_date"`
			Percent            float64     `json:"percent"`
			TickSize           float64     `json:"tick_size"`
			FloatShares        interface{} `json:"float_shares"`
			LimitDown          float64     `json:"limit_down"`
			Amplitude          float64     `json:"amplitude"`
			Current            float64     `json:"current"`
			High               float64     `json:"high"`
			CurrentYearPercent float64     `json:"current_year_percent"`
			FloatMarketCapital interface{} `json:"float_market_capital"`
			IssueDate          int64       `json:"issue_date"`
			Low                float64     `json:"low"`
			SubType            string      `json:"sub_type"`
			MarketCapital      float64     `json:"market_capital"`
			Currency           string      `json:"currency"`
			LotSize            int         `json:"lot_size"`
			LockSet            interface{} `json:"lock_set"`
			Iopv               float64     `json:"iopv"`
			Timestamp          int64       `json:"timestamp"`
			FoundDate          int64       `json:"found_date"`
			Amount             float64     `json:"amount"`
			Chg                float64     `json:"chg"`
			LastClose          float64     `json:"last_close"`
			Volume             int         `json:"volume"`
			VolumeRatio        interface{} `json:"volume_ratio"`
			LimitUp            float64     `json:"limit_up"`
			TurnoverRate       interface{} `json:"turnover_rate"`
			Low52W             float64     `json:"low52w"`
			Name               string      `json:"name"`
			PremmRate          float64     `json:"premm_rate"`
			Exchange           string      `json:"exchange"`
			UnitNav            float64     `json:"unit_nav"`
			Time               int64       `json:"time"`
			TotalShares        int64       `json:"total_shares"`
			Open               float64     `json:"open"`
			Status             int         `json:"status"`
		} `json:"quote"`
		Others struct {
			PankouRatio float64 `json:"pankou_ratio"`
			CybSwitch   bool    `json:"cyb_switch"`
		} `json:"others"`
		Tags []struct {
			Description string `json:"description"`
			Value       int    `json:"value"`
		} `json:"tags"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type Kline struct {
	Time    float64 `json:"time"`
	Volume  float64 `json:"volume"`
	Amount  float64 `json:"amount"`
	Open    float64 `json:"open"`
	Close   float64 `json:"close"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Chg     float64 `json:"chg"`
	Percent float64 `json:"percent"`
}

type StockCN struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
