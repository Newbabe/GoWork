package main

import (
	"demo1/util"
	"fmt"
	"strconv"
	"time"
)

type tvLog struct {
	Id        int
	RoomId    int
	ChannelId string
	Mac       string
	Wifi      string
	VId       int
}

type LEInfo struct {
	LEInfo []tvLog
}

func main() {
	li := new(LEInfo)
	li.GetMothData()
}

func (lei *LEInfo) GetMothData() {
	fmt.Println(" 日活數量统计")
	num := make([]tvLog, 0)
	for i := 0; i <= 20; i++ {
		t, _ := time.Parse("2006-01-02 15:04:05", "2023-05-01 00:00:00")
		createTime := t.AddDate(0, 0, i)
		endTime := createTime.AddDate(0, 0, 1)
		//fmt.Println("createTime", createTime.Format("2006-01-02 15:04:05"), "endTime", endTime.Format("2006-01-02 15:04:05"))
		list := GetTvLogData(createTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
		lei.LEInfo = list
		removal := lei.duplicateRemoval()
		num = append(num, removal...)
		fmt.Println(createTime.Format("2006-01-02"), "=", "去重前:", len(list), "去重后:", len(removal))
	}
	fmt.Println("总数", len(num))
}

func (lei *LEInfo) duplicateRemoval() []tvLog {
	var list = make([]tvLog, 0, len(lei.LEInfo))
	var m = make(map[string]tvLog)
	for _, v := range lei.LEInfo {
		/*base16 := GetBase16(strconv.Itoa(v.VId))
		if base16 == "0xd8c" {
			//fmt.Println("base16", base16)

		}*/
		m[v.Mac+" "+v.Wifi] = v
	}
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

func GetTvLogData(createTime, endTime string) []tvLog {
	sql := "select id ,room_id ,ethernet_mac ,wifi_mac,channel_id,vid from hero_ok_web.tv_login_log where create_time > ? and create_time <= ? and channel_id in (10000,10014,10015) "
	query, err := util.GetDbHerookWebRead().Query(sql, createTime, endTime)
	if err != nil {
		return nil
	}
	var list []tvLog
	for query.Next() {
		var tv tvLog
		query.Scan(&tv.Id, &tv.RoomId, &tv.Mac, &tv.Wifi, &tv.ChannelId, &tv.VId)
		list = append(list, tv)
	}
	return list
}
func GetBase16(str string) string {
	num, _ := strconv.Atoi(str)
	str16 := fmt.Sprintf("%x", num)
	if num == 0 {
		return "0x0"
	}
	return "0x" + str16
}
func GetPlaySql(createTime, endTime string) []tvLog {
	sql := "select id ,room_id ,ethernet_mac ,wifi_mac,channel_id,vid from hero_ok_web.playback_record where create_time > ? and create_time <= ? and channel_id "
	query, err := util.GetDbHerookWebRead().Query(sql, createTime, endTime)
	if err != nil {
		return nil
	}
	var list []tvLog
	for query.Next() {
		var tv tvLog
		query.Scan(&tv.Id, &tv.RoomId, &tv.Mac, &tv.Wifi, &tv.ChannelId, &tv.VId)
		list = append(list, tv)
	}
	return list
}
func (lei *LEInfo) GetPlayCount() {
	num := make([]tvLog, 0)
	for i := 0; i <= 30; i++ {
		t, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
		createTime := t.AddDate(0, 0, i)
		endTime := createTime.AddDate(0, 0, 1)
		//fmt.Println("createTime", createTime.Format("2006-01-02 15:04:05"), "endTime", endTime.Format("2006-01-02 15:04:05"))
		list := GetPlaySql(createTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
		lei.LEInfo = list
		removal := lei.duplicateRemoval()
		num = append(num, removal...)
		fmt.Println(createTime.Format("2006-01-02"), ":", len(removal))
	}
	fmt.Println("总数", len(num))
}
