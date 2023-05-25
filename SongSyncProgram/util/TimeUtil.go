package util

import (
	"time"
)

const timeStr = "2006-01-02 15:04:05"
const data = "20060102"

func GetNowTime() string {
	dataStr := time.Now().Format(timeStr)
	return dataStr
}
func GetNowData() string {
	strTime := time.Now().Format(data)

	return strTime
}
func GetNowData2() string {
	strTime := time.Now().AddDate(0, 0, -3).Format(data)

	return strTime
}
