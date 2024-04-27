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
