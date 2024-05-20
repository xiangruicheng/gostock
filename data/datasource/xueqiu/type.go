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
			CurrentExt               interface{} `json:"current_ext"`
			Symbol                   string      `json:"symbol"`
			VolumeExt                interface{} `json:"volume_ext"`
			High52W                  float64     `json:"high52w"`
			Delayed                  int         `json:"delayed"`
			Type                     int         `json:"type"`
			TickSize                 float64     `json:"tick_size"`
			FloatShares              float64     `json:"float_shares"`
			LimitDown                float64     `json:"limit_down"`
			NoProfit                 string      `json:"no_profit"`
			High                     float64     `json:"high"`
			FloatMarketCapital       float64     `json:"float_market_capital"`
			TimestampExt             interface{} `json:"timestamp_ext"`
			LotSize                  int         `json:"lot_size"`
			LockSet                  interface{} `json:"lock_set"`
			WeightedVotingRights     string      `json:"weighted_voting_rights"`
			Chg                      float64     `json:"chg"`
			Eps                      float64     `json:"eps"`
			LastClose                float64     `json:"last_close"`
			ProfitFour               float64     `json:"profit_four"`
			Volume                   float64     `json:"volume"`
			VolumeRatio              float64     `json:"volume_ratio"`
			ProfitForecast           float64     `json:"profit_forecast"`
			TurnoverRate             float64     `json:"turnover_rate"`
			Low52W                   float64     `json:"low52w"`
			Name                     string      `json:"name"`
			Exchange                 string      `json:"exchange"`
			PeForecast               float64     `json:"pe_forecast"`
			TotalShares              float64     `json:"total_shares"`
			Status                   int         `json:"status"`
			IsVieDesc                string      `json:"is_vie_desc"`
			SecurityStatus           interface{} `json:"security_status"`
			Code                     string      `json:"code"`
			GoodwillInNetAssets      float64     `json:"goodwill_in_net_assets"`
			AvgPrice                 float64     `json:"avg_price"`
			Percent                  float64     `json:"percent"`
			WeightedVotingRightsDesc string      `json:"weighted_voting_rights_desc"`
			Amplitude                float64     `json:"amplitude"`
			Current                  float64     `json:"current"`
			IsVie                    string      `json:"is_vie"`
			CurrentYearPercent       float64     `json:"current_year_percent"`
			IssueDate                int64       `json:"issue_date"`
			SubType                  string      `json:"sub_type"`
			Low                      float64     `json:"low"`
			IsRegistrationDesc       string      `json:"is_registration_desc"`
			NoProfitDesc             string      `json:"no_profit_desc"`
			MarketCapital            float64     `json:"market_capital"`
			Dividend                 float64     `json:"dividend"`
			DividendYield            float64     `json:"dividend_yield"`
			Currency                 string      `json:"currency"`
			Navps                    float64     `json:"navps"`
			Profit                   float64     `json:"profit"`
			Timestamp                int64       `json:"timestamp"`
			PeLyr                    float64     `json:"pe_lyr"`
			Amount                   float64     `json:"amount"`
			PledgeRatio              float64     `json:"pledge_ratio"`
			TradedAmountExt          interface{} `json:"traded_amount_ext"`
			IsRegistration           string      `json:"is_registration"`
			Pb                       float64     `json:"pb"`
			LimitUp                  float64     `json:"limit_up"`
			PeTtm                    float64     `json:"pe_ttm"`
			Time                     int64       `json:"time"`
			Open                     float64     `json:"open"`
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
