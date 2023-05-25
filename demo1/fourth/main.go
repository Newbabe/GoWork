package main

import (
	"bufio"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

type ExcelTest struct {
	FileName  string //文件地址
	SheetName string //sheet页名称

}

// 保存excel文件
func (e *ExcelTest) save(userMap map[string]int) {
	file, _ := excelize.OpenFile(e.FileName)
	rows, _ := file.GetRows(e.SheetName)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		file.SetCellValue(e.SheetName, "H"+strconv.Itoa(1+i), userMap[row[0]])
	}
	file.Save()
}

// 获取excel文件
func (e *ExcelTest) GetExcelUserId() []int {
	file, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return nil
	}
	var list []int
	rows, _ := file.GetRows(e.SheetName)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for i2, cell := range row {
			if i2 == 0 {
				userId, _ := strconv.Atoi(cell)
				list = append(list, userId)
			}
			continue
		}
	}
	return list
}

// 日志结构
type Log struct {
	FileName string
	Url      string
}

func (l *Log) GetLogData() []string {
	open, err := os.Open(l.Url + l.FileName)
	if err != nil {
		return nil
	}
	var list []string
	reader := bufio.NewReader(open)
	//按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		list = append(list, string(line))
	}
	return list
}
func GetDatelog() [][]string {
	var data = make([][]string, 0, 14)
	parse, _ := time.Parse("2006-01-02", "2023-04-26")
	url := "/home/ec2-user/ExpDiamondCoinLogRpc/static/exp_coin_diamond/exp_coin_diamond_"
	for i := 0; i <= 13; i++ {
		format := parse.AddDate(0, 0, i).Format("2006-01-02")
		l := Log{
			FileName: format,
			Url:      url,
		}
		logData := l.GetLogData()
		data = append(data, logData)
	}
	return data
}
func main() {
	e := ExcelTest{
		FileName:  "/home/ec2-user/GOTest/Log/星币超10万用户数据详情.xlsx",
		SheetName: "sheet",
	}
	userIdList := e.GetExcelUserId()
	userNumMap := make(map[string]int, len(userIdList))
	for _, userId := range userIdList {
		userNumMap[strconv.Itoa(userId)] = 0
	}
	datelogList := GetDatelog()
	for _, stringList := range datelogList {
		for _, str := range stringList {
			split := strings.Split(str, "|")
			userIdStr := split[0]
			value, ok := userNumMap[userIdStr]
			if ok && split[3] == "55" {
				userNumMap[userIdStr] = value + 1
			}
		}
	}
	e.save(userNumMap)
}
