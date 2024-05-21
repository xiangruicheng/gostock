package indicator

import (
	"fmt"
	"gostock/model"
	"gostock/server"
	"time"
)

type Macd struct{}

func (s *Macd) BatchRun() {
	stockInfos, _ := new(model.StockInfoModel).GetAll()
	count := len(stockInfos)
	for k, item := range stockInfos {
		server.Log.Info(fmt.Sprintf("(%d/%d) %s run macd", k, count, item.Code))
		s.Run(item.Code)
	}
}

// Run 计算单个股票MACD
func (s *Macd) Run(code string) {
	var newBulidTag bool
	var prevEma12 float64
	var prevEma26 float64
	var prevDea float64
	var lastRow *model.MacdRecord
	var beginDate string

	hisMacdDatas, _ := new(model.MacdModel).GetByCode(code)
	if len(hisMacdDatas) == 0 {
		newBulidTag = true
		beginDate = "20030101"
	} else {
		newBulidTag = false
		lastRow = hisMacdDatas[len(hisMacdDatas)-1]
		beginDate = lastRow.Date
	}

	//行情数据
	klines, _ := new(model.KlineModel).GetByCodeGtDate(code, beginDate)
	if len(klines) <= 0 {
		server.Log.Error(fmt.Sprintf("%s not kline data,can not macd", code))
		return
	}
	rows := s.klineRecordConvertMacdRecord(klines)

	for key, row := range rows {
		if newBulidTag {
			if key == 0 {
				continue
			} else if key == 1 {
				prevEma12 = rows[0].Close
				prevEma26 = rows[0].Close
				prevDea = 0
			} else {
				prevEma12 = rows[key-1].Ema12
				prevEma26 = rows[key-1].Ema26
				prevDea = rows[key-1].Dea
			}
		} else {
			if key == 0 {
				prevEma12 = lastRow.Ema12
				prevEma26 = lastRow.Ema26
				prevDea = lastRow.Dea
			} else {
				prevEma12 = rows[key-1].Ema12
				prevEma26 = rows[key-1].Ema26
				prevDea = rows[key-1].Dea
			}
		}

		ema12, ema26, diff, dea, macd := s.getMacd(prevEma12, prevEma26, prevDea, row.Close)
		row.Ema12 = ema12
		row.Ema26 = ema26
		row.Diff = diff
		row.Dea = dea
		row.Macd = macd
		row.CTime = time.Now().Unix()
		row.UTime = time.Now().Unix()
	}
	_, err := new(model.MacdModel).BatchInsert(rows)
	if err != nil {
		server.Log.Error(fmt.Sprintf("%s update macd fail", code))
	}
}

func (s *Macd) klineRecordConvertMacdRecord(list []*model.KlineRecord) []*model.MacdRecord {
	resp := []*model.MacdRecord{}
	for _, row := range list {
		macdModel := new(model.MacdRecord)
		macdModel.Code = row.Code
		macdModel.Date = row.Date
		macdModel.Close = row.Close
		resp = append(resp, macdModel)
	}
	return resp
}

/**
 * core MACD核心公式
 * 算法 12 26 9
 * EMA（12）= 前一日EMA（12）×11/13＋今日收盘价×2/13
 * EMA（26）= 前一日EMA（26）×25/27＋今日收盘价×2/27
 * DIFF=今日EMA（12）- 今日EMA（26）
 * DEA（MACD）= 前一日DEA×8/10＋今日DIF×2/10
 * MACD=2×(DIFF－DEA)
 * @param string $prev 前一日的值
 * @param string $today 当日的值
 * @param string $n 周期
 * @return float|int
 */
func (s *Macd) core(prev float64, today float64, n float64) float64 {
	return prev*(n-1)/(n+1) + today*2/(n+1)
}

func (s *Macd) getMacd(prev_ema12 float64, prev_ema26 float64, prev_dea float64, today float64) (float64, float64, float64, float64, float64) {
	ema12 := s.core(prev_ema12, today, 12)
	ema26 := s.core(prev_ema26, today, 26)
	diff := ema12 - ema26
	dea := s.core(prev_dea, diff, 9)
	macd := 2*diff - 2*dea

	return ema12, ema26, diff, dea, macd
}
