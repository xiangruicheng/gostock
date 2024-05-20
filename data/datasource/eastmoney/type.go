package eastmoney

type PeopleResponse struct {
	Version string `json:"version"`
	Result  struct {
		Pages int `json:"pages"`
		Data  []struct {
			ENDDATE        string  `json:"END_DATE"`
			TOTALASHARES   int64   `json:"TOTAL_A_SHARES"`
			AVGHOLDNUM     float64 `json:"AVG_HOLD_NUM"`
			TOTALMARKETCAP float64 `json:"TOTAL_MARKET_CAP"`
			AVGMARKETCAP   float64 `json:"AVG_MARKET_CAP"`
			HOLDERNUM      int64   `json:"HOLDER_NUM"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
