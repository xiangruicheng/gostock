package eastmoney

import (
	"encoding/json"
	"errors"
	"fmt"
	"gostock/model"
	"io"
	"net/http"
	"time"
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

func StockAll(market string) ([]*model.StockInfoRecord, error) {
	records := []*model.StockInfoRecord{}
	pz := "20"
	for pn := 1; pn < 2000; pn++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("执行分页pn=%d\n", pn)
		url := ""
		if market == "SH" {
			url = "https://push2.eastmoney.com/api/qt/clist/get?np=1&fltt=1&invt=2&fs=m%3A1%2Bt%3A2%2Cm%3A1%2Bt%3A23&fields=f12,f14&fid=f3&pn=" + fmt.Sprintf("%d", pn) + "&pz=" + pz + "&po=1&dect=1&ut=fa5fd1943c7b386f172d6893dbfba10b&wbp2u=%7C0%7C0%7C0%7Cweb&_=1754745689393"
		}
		if market == "SZ" {
			url = "https://push2.eastmoney.com/api/qt/clist/get?np=1&fltt=1&invt=2&fs=m%3A0%2Bt%3A6%2Cm%3A0%2Bt%3A80&fields=f12,f14&fid=f3&pn=" + fmt.Sprintf("%d", pn) + "&pz=" + pz + "&po=1&dect=1&ut=fa5fd1943c7b386f172d6893dbfba10b&wbp2u=%7C0%7C0%7C0%7Cweb&_=1754745689397"
		}
		if url == "" {
			return nil, errors.New("market error")
		}

		stockAllRespons := new(StockAllResponse)
		responseStr := request(url)

		if responseStr == "" {
			return records, errors.New("resp is empty")
		}
		err := json.Unmarshal([]byte(responseStr), &stockAllRespons)
		if len(stockAllRespons.Data.Diff) <= 0 {
			break
		}
		if err != nil {
			return records, err
		}
		for _, item := range stockAllRespons.Data.Diff {
			records = append(records, &model.StockInfoRecord{
				Market: market,
				Name:   item.Name,
				Code:   item.Code,
			})
		}
	}
	return records, nil
}

func Block(blockType int64) (*StockAllResponse, error) {
	url := ""
	if blockType == 1 {
		url = "https://39.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=1000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f3&fs=m:90+t:2+f:!50&fields=f12,f14"
	}
	if blockType == 2 {
		url = "https://39.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=1000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f3&fs=m:90+t:3+f:!50&fields=f12,f14"
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

func BlockCode(bkCode string) (*StockAllResponse, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get?po=1&pz=500&pn=1&np=1&fltt=2&invt=2&fs=b%3A" + bkCode + "&fields=f12,f14"
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

func Zdfb() (*ZdfbResponse, error) {
	url := "https://push2ex.eastmoney.com/getTopicZDFenBu?ut=7eea3edcaed734bea9cbfc24409ed989&dpt=wz.ztzt"
	zdfbResponse := new(ZdfbResponse)
	responseStr := request(url)
	if responseStr == "" {
		return zdfbResponse, errors.New("resp is empty")
	}
	err := json.Unmarshal([]byte(responseStr), &zdfbResponse)
	if err != nil {
		return zdfbResponse, err
	}
	return zdfbResponse, nil
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
