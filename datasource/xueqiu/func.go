package xueqiu

import (
	"encoding/json"
	"fmt"
	"gostock/config"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// KlineList kline list
func KlineList(symbol string, days int64) ([]*Kline, error) {
	var klines []*Kline

	api := "https://stock.xueqiu.com/v5/stock/chart/kline.json"

	var uri url.URL
	param := uri.Query()
	param.Add("symbol", symbol)
	param.Add("begin", strconv.FormatInt(time.Now().UnixMilli(), 10))
	param.Add("period", "day")
	param.Add("type", "before")
	param.Add("count", strconv.Itoa(-1*int(days)))
	param.Add("indicator", "kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance")
	queryStr := param.Encode()
	url := fmt.Sprintf("%s?%s", api, queryStr)

	responseStr := request(url)
	if responseStr == "" {
		return klines, nil
	}

	klineListResponse := new(KlineListResponse)
	json.Unmarshal([]byte(responseStr), &klineListResponse)
	for _, item := range klineListResponse.Data.Item {
		kline := new(Kline)
		kline.Time = item[0]
		kline.Volume = item[1]
		kline.Open = item[2]
		kline.High = item[3]
		kline.Low = item[4]
		kline.Close = item[5]
		kline.Chg = item[6]
		kline.Percent = item[7]
		kline.Amount = item[9]
		klines = append(klines, kline)
	}
	return klines, nil
}

// request send http request
func request(url string) string {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//server.Log(server.LogleveInfo, err.Error())
		return ""
	}

	//增加header选项
	reqest.Header.Add("Referer", "https://xueqiu.com/")
	reqest.Header.Add("Origin", "https://xueqiu.com")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	reqest.Header.Add("cookie", config.Data.Xueqiu.Cookie)
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		//server.Log(server.LogleveInfo, err.Error())
		return ""
	}
	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		//server.Log(server.LogleveInfo, err.Error())
		return ""
	}
	return string(resp)
}
