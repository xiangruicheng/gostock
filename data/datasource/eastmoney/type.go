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

type Hs300Response struct {
	Version string `json:"version"`
	Result  struct {
		Pages int `json:"pages"`
		Data  []struct {
			SECUCODE         string  `json:"SECUCODE"`
			SECURITYCODE     string  `json:"SECURITY_CODE"`
			TYPE             string  `json:"TYPE"`
			SECURITYNAMEABBR string  `json:"SECURITY_NAME_ABBR"`
			CLOSEPRICE       float64 `json:"CLOSE_PRICE"`
			INDUSTRY         string  `json:"INDUSTRY"`
			REGION           string  `json:"REGION"`
			WEIGHT           float64 `json:"WEIGHT"`
			EPS              float64 `json:"EPS"`
			BPS              float64 `json:"BPS"`
			ROE              float64 `json:"ROE"`
			TOTALSHARES      float64 `json:"TOTAL_SHARES"`
			FREESHARES       float64 `json:"FREE_SHARES"`
			FREECAP          float64 `json:"FREE_CAP"`
			F2               float64 `json:"f2"`
			F3               float64 `json:"f3"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type StockAllResponse struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int64  `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Total int `json:"total"`
		Diff  []struct {
			Code string `json:"f12"`
			Name string `json:"f14"`
		} `json:"diff"`
	} `json:"data"`
}
