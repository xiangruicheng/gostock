package util

import (
	"time"
)

// FormatDate 格式化日期
func FormatDate(timestamp float64) string {
	t := time.Unix(int64(timestamp/1000), 0)
	date := t.Format("20060102")
	return date
}
