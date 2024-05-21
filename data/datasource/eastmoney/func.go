package eastmoney

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Hs300() (*Hs300Response, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECURITY_CODE&sortTypes=-1&pageSize=400&pageNumber=1&reportName=RPT_INDEX_TS_COMPONENT&columns=SECUCODE%2CSECURITY_CODE&quoteColumns=f2%2Cf3&source=WEB&client=WEB&filter=(TYPE%3D%221%22)"
	hs300Response := new(Hs300Response)
	responseStr := request(url)
	if responseStr == "" {
		return hs300Response, errors.New("resp is empty")
	}
	err := json.Unmarshal([]byte(responseStr), &hs300Response)
	if err != nil {
		return hs300Response, err
	}
	return hs300Response, nil
}

func People(code string) (*PeopleResponse, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_HOLDERNUM_DET&columns=END_DATE%2CTOTAL_A_SHARES%2CAVG_HOLD_NUM%2CTOTAL_MARKET_CAP%2CAVG_MARKET_CAP%2CHOLDER_NUM&filter=(SECURITY_CODE%3D%22" + code + "%22)&source=WEB&client=WEB&sortColumns=END_DATE&sortTypes=-1&pageSize=500"
	peopleResponse := new(PeopleResponse)
	responseStr := request(url)
	if responseStr == "" {
		return peopleResponse, errors.New("resp is empty")
	}
	err := json.Unmarshal([]byte(responseStr), &peopleResponse)
	if err != nil {
		return peopleResponse, err
	}
	return peopleResponse, nil
}

func Cyb() (*StockAllResponse, error) {
	url := "https://45.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f3&fs=m:0+t:80&fields=f12,f14"
	stockAllRespons := new(StockAllResponse)
	responseStr := request(url)
	if responseStr == "" {
		return stockAllRespons, errors.New("resp is empty")
	}
	err := json.Unmarshal([]byte(responseStr), &stockAllRespons)
	if err != nil {
		return stockAllRespons, err
	}
	return stockAllRespons, nil
}

func StockAll(market string) (*StockAllResponse, error) {
	url := ""
	if market == "SH" {
		url = "https://45.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f3&fs=m:1+t:2,m:1+t:23&fields=f12,f14"
	}
	if market == "SZ" {
		url = "https://45.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f3&fs=m:0+t:6,m:0+t:80&fields=f12,f14"
	}
	if url == "" {
		return nil, errors.New("market error")
	}

	stockAllRespons := new(StockAllResponse)
	responseStr := request(url)
	if responseStr == "" {
		return stockAllRespons, errors.New("resp is empty")
	}
	err := json.Unmarshal([]byte(responseStr), &stockAllRespons)
	if err != nil {
		return stockAllRespons, err
	}
	return stockAllRespons, nil
}

func request(url string) string {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//server.Log(server.LogleveInfo, err.Error())
		return ""
	}
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
