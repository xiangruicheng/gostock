package eastmoney

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Hs300() (*Hs300Response, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECURITY_CODE&sortTypes=-1&pageSize=400&pageNumber=1&reportName=RPT_INDEX_TS_COMPONENT&columns=SECUCODE%2CSECURITY_CODE%2CTYPE%2CSECURITY_NAME_ABBR%2CCLOSE_PRICE%2CINDUSTRY%2CREGION%2CWEIGHT%2CEPS%2CBPS%2CROE%2CTOTAL_SHARES%2CFREE_SHARES%2CFREE_CAP&quoteColumns=f2%2Cf3&source=WEB&client=WEB&filter=(TYPE%3D%221%22)"
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
