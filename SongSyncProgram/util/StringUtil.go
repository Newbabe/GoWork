package util

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

//把字符串转换成16进制
func GetBase16(str string) string {
	num, _ := strconv.Atoi(str)
	str16 := fmt.Sprintf("%x", num)
	if num == 0 {
		return "0x0"
	}
	return "0x" + str16
}
func GetFile(filePath string) *os.File {
	filename := "access_log_" + time.Now().Format("2006-01-02") + ".txt"
	//filePath := "/home/ec2-user/HeroOkApi/logs/"

	if !IsExist(filePath) {
		os.MkdirAll(filePath, os.ModePerm)
	}
	var f *os.File
	if !IsExist(filePath + filename) {
		f, _ = os.Create(filePath + filename)
	} else {
		f, _ = os.OpenFile(filePath+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	}
	return f

}
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
