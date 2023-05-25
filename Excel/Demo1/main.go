package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	/*data1 := GetInputData2("D:\\GoWork\\Excel\\3.10号私域10703.xlsx", "数据段1")
	fmt.Println("数量1", len(data1))
	data2 := GetInputDataNew2("D:\\GoWork\\Excel\\10703.xlsx", "Sheet")
	fmt.Println("数量2", len(data2))
	for s, s2 := range data2 {
		if data1[s] != s2 {
			fmt.Println("数据不一致的", s)
		}
	}*/

	data1 := GetInputData("D:\\GoWork\\Excel\\3.10号云仓10627(1).xlsx", "Demo")
	fmt.Println("数量1", len(data1))
	dataNew2 := GetInputDataNew("D:\\GoWork\\Excel\\10627(2).xlsx", "Sheet")
	fmt.Println("数量2", len(dataNew2))
	for s, s2 := range data1 {
		if dataNew2[s] != s2 {
			fmt.Println("数据不一致的", s)
		}
	}
}

func GetInputData(url, sheetName string) map[string]string {

	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
	}
	//只获取
	var dataMap = make(map[string]string)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}
	for i, row := range rows { //获取所有行数据
		if i < 1 {
			continue
		}

		dataMap[row[2]] = row[4]

	}

	return dataMap
}
func GetInputData2(url, sheetName string) map[string]string {

	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
	}
	//只获取
	var dataMap = make(map[string]string)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}
	for i, row := range rows { //获取所有行数据
		if i < 1 {
			continue
		}
		dataMap[row[0]] = row[3]
	}

	return dataMap
}
func GetInputDataNew(url, sheetName string) map[string]string {

	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
	}
	//只获取
	var dataMap = make(map[string]string)
	rows, _ := f.GetRows(sheetName)
	for i, row := range rows { //获取所有行数据
		if i < 1 {
			continue
		}

		dataMap[row[2]] = row[3]

	}

	return dataMap
}
func GetInputDataNew2(url, sheetName string) map[string]string {

	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
	}
	//只获取
	var dataMap = make(map[string]string)
	rows, _ := f.GetRows(sheetName)
	for i, row := range rows { //获取所有行数据
		if i < 1 {
			continue
		}

		dataMap[row[2]] = row[4]

	}

	return dataMap
}
