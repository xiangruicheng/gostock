package util

import (
	"strconv"
	"time"
)

// FormatAmout 格式化金额
func FormatAmout(amount float64) string {
	if amount > 10000 {
		if amount > 100000000 {
			return strconv.FormatFloat(amount/100000000, 'f', 2, 64)
		} else {
			return strconv.FormatFloat(amount/10000, 'f', 2, 64)
		}
	}
	return strconv.FormatFloat(amount, 'f', 2, 64)
}

// FormatFloat 格式化日期
func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

// FormatDate 格式化日期
func FormatDate(timestamp float64) string {
	t := time.Unix(int64(timestamp/1000), 0)
	date := t.Format("20060102")
	return date
}

// NowDateTime 获取当前时间
func NowDateTime() string {
	t := time.Now()
	date := t.Format("2006-01-02 15:04:05")
	return date
}

// code2Market 根据股票代码得到是那个市场 SZ SH BJ
// 规则如下
// 沪市A股股票代码均以6开头
// 沪市B股股票代码以900
//
// 深市主板股票代码以0
// 深市B股股票代码以200开头
// 3开头创业板
//
// 北京新三板挂牌公司，股票代码以4或8开头
func Code2Market(code string) string {
	codeBytes := []byte(code)
	//深圳
	if string(codeBytes[0]) == "0" ||
		string(codeBytes[0]) == "3" ||
		string(codeBytes[0]+codeBytes[1]+codeBytes[2]) == "200" {
		return "SZ"
	}

	if string(codeBytes[0]) == "6" ||
		string(codeBytes[0]+codeBytes[1]+codeBytes[2]) == "900" {
		return "SH"
	}
	if string(codeBytes[0]) == "4" ||
		string(codeBytes[0]) == "8" {
		return "BJ"
	}
	return ""
}
