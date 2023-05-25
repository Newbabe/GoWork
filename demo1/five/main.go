package main

import (
	"demo1/util"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type ExcelTest struct {
	FileName  string //文件地址
	SheetName string //sheet页名称

}

// 保存excel文件
func (e *ExcelTest) save(name map[string]string) {
	file, _ := excelize.OpenFile(e.FileName)
	/*sheet, _ := file.NewSheet(e.SheetName)
	file.SetActiveSheet(sheet)*/
	for k, v := range name {
		file.SetCellValue(e.SheetName, k, v)
	}
	if err := file.SaveAs(e.FileName); err != nil {
		fmt.Println(err)
		return
	}
}
func (e *ExcelTest) GetExcelData() []ExcelInfo {
	open, err := excelize.OpenFile(e.FileName)
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := open.GetRows(e.SheetName)
	var list = make([]ExcelInfo, 0, len(rows))
	for i, row := range rows {
		if i < 2 {
			continue
		}
		var excelInfo ExcelInfo
		excelInfo.Channel = row[0]
		excelInfo.MacWifi = row[1]
		excelInfo.Mac = row[2]
		excelInfo.Wifi = row[3]
		excelInfo.TvBox = row[4]
		list = append(list, excelInfo)
	}
	return list
}

type ExcelInfo struct {
	Channel string
	MacWifi string
	Mac     string
	Wifi    string
	TvBox   string
}

type LEInfo struct {
	LEInfo []ExcelInfo
}

/*
	func (lei *LEInfo) duplicateRemoval() []ExcelInfo {
		var list = make([]ExcelInfo, 0, len(lei.LEInfo))
		var m = make(map[string]ExcelInfo)
		for _, v := range lei.LEInfo {
			m[v.MacWifi] = v
		}
		for _, v := range m {
			list = append(list, v)
		}
		return list
	}
*/
func (lei *LEInfo) GetCount() map[string]string {
	colArray := []string{"F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC"}

	var excelMap = make(map[string]string)
	for i := 0; i < 24; i++ {
		startTime, _ := time.Parse("2006-01-02 15:04:05", "2021-06-01 00:00:00")
		overTime, _ := time.Parse("2006-01-02 15:04:05", "2021-07-01 00:00:00")
		startTime = startTime.AddDate(0, i, 0)
		overTime = overTime.AddDate(0, i, 0)
		fmt.Println("startTime", startTime, "overTime", overTime)
		for j, info := range lei.LEInfo {
			if j == 3 {
				break
			}
			count := getCount(info.Mac, info.Wifi, startTime.Format("2006-01-02 15:04:05"), overTime.Format("2006-01-02 15:04:05"))
			excelMap[colArray[i]+strconv.Itoa(j+3)] = strconv.Itoa(count)
			fmt.Println("Mac", info.Mac, "wifi", info.Wifi, "colArray[i]+strconv.Itoa(j+3)", colArray[i]+strconv.Itoa(j+3), "count", count)

		}

	}
	return excelMap
}
func getCount(mac, wifi, startTime, endTime string) int {
	sql := "select count(*) from hero_ok_web.playback_record where ethernet_mac = ? and wifi_mac = ? and create_time >= ? and create_time <= ?"
	row := util.GetDbHerookWebRead().QueryRow(sql, mac, wifi, startTime, endTime)
	var count int
	if err := row.Scan(&count); err != nil {
		fmt.Println(err)
	}
	return count
}

func main() {
	test := ExcelTest{
		FileName:  "/home/ec2-user/GOTest/songinfo/副本0517特定設備資料統計.xlsx",
		SheetName: "總表",
	}
	data := test.GetExcelData()
	l := new(LEInfo)
	l.LEInfo = data
	count := l.GetCount()
	test.save(count)
}
