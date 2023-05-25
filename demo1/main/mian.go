package main

import (
	"demo1/util"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// 原始层 1、获取excel参数  2、将数据写入excel文件
type ExcelTest struct {
	FileName  string //文件地址
	SheetName string //sheet页名称

}

// 保存excel文件
func (e *ExcelTest) save(name map[string]interface{}) {
	file := excelize.NewFile()
	sheet, _ := file.NewSheet(e.SheetName)
	file.SetActiveSheet(sheet)
	for k, v := range name {
		file.SetCellValue(e.SheetName, k, v)
	}
	if err := file.SaveAs(e.FileName); err != nil {
		fmt.Println(err)
		return
	}
}
func (e *ExcelTest) GetUsData() map[string]string {
	xlFile, err := xlsx.OpenFile(e.FileName)
	if err != nil {
		fmt.Println("获取我方设备数据异常")
	}
	var emacMap = make(map[string]string)
	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			if index == 0 {
				continue
			}
			emac := row.Cells[1].String()
			if strings.Contains(emac, ":") {
				emac = strings.Replace(emac, " ", "", -1)
				emac = strings.Replace(emac, ":", "", -1)
			}
			emac = strings.ToLower(emac)
			emacMap[emac] = emac
		}
	}
	return emacMap
}

type TvBox struct {
	TvBoxs []string
}
type TvInfo struct {
	RoomId      int
	DeviceId    string
	ChannelId   int
	WifiMac     string
	EthernetMac string
	CreateTime  time.Time
	TvBox       string
}

func (name *TvBox) GetTvInfoByTvBox() ([]TvInfo, error) {
	str := strings.Join(name.TvBoxs, "','")
	fmt.Println("str", str)
	sql := "select room_id,device_id,channel_id,wifi_mac,ethernet_mac,create_time,tv_box from  hero_ok_web.tv_info where tv_box in ('" + str + "')"

	query, err := util.GetDbHerookWebRead().Query(sql)
	if err != nil {
		return nil, err
	}
	var list []TvInfo
	for query.Next() {
		var s TvInfo
		query.Scan(&s.RoomId, &s.DeviceId, &s.ChannelId, &s.WifiMac, &s.EthernetMac, &s.CreateTime, &s.TvBox)
		list = append(list, s)
	}
	return list, nil

}
func duplicateRemoval(t []TvInfo) []TvInfo {
	var StrMap = make(map[string]string)
	var list []TvInfo
	for _, info := range t {
		_, ok := StrMap[info.EthernetMac+" "+info.WifiMac]
		if !ok {
			list = append(list, info)
			StrMap[info.EthernetMac+" "+info.WifiMac] = info.TvBox
		}
	}
	return list
}

func ToMapString(ts []TvInfo) map[string]interface{} {
	var mapStr = make(map[string]interface{})
	//title数据
	mapStr["A1"] = "channel"
	mapStr["B1"] = "月份"
	mapStr["C1"] = "日期"
	mapStr["D1"] = "RoomID"
	mapStr["E1"] = "MAc+wifi"
	mapStr["F1"] = "Mac"
	mapStr["G1"] = "wif"
	mapStr["H1"] = "TvBox"
	mapStr["I1"] = "pid"
	mapStr["J1"] = "vid"
	mapStr["K1"] = "是否为我放Id"
	mapStr["L1"] = "是否有评分"
	//获取我方设备
	e := ExcelTest{
		FileName:  "/home/ec2-user/UpdateTwTvTagImg/static/TW/excel/我方设备数据.xlsx",
		SheetName: "工作表1",
	}
	usData := e.GetUsData()
	fmt.Println("我方設備數據", len(usData))
	for i, info := range ts {
		mapStr["A"+strconv.Itoa(i+2)] = info.ChannelId
		mapStr["B"+strconv.Itoa(i+2)] = info.CreateTime.Format("01")
		mapStr["C"+strconv.Itoa(i+2)] = info.CreateTime.Format("2006-01-02 15:04:05")
		mapStr["D"+strconv.Itoa(i+2)] = info.RoomId
		mapStr["E"+strconv.Itoa(i+2)] = info.EthernetMac + info.WifiMac
		mapStr["F"+strconv.Itoa(i+2)] = info.EthernetMac
		mapStr["G"+strconv.Itoa(i+2)] = info.WifiMac
		mapStr["H"+strconv.Itoa(i+2)] = info.TvBox
		//pid vid 处理
		var pid, vid int
		loginLog := GetMicrophoneInfoByRoomId(info.EthernetMac, info.CreateTime.Format("2006-01-02"), info.DeviceId)
		pid = loginLog.Pid
		vid = loginLog.Vid
		if pid == 0 || vid == 0 {
			paybackRecord := GetPbRecordBYRoomId(info.RoomId)
			pid, _ = strconv.Atoi(paybackRecord.Pid)
			vid, _ = strconv.Atoi(paybackRecord.Vid)
		}
		mapStr["I"+strconv.Itoa(i+2)] = pid
		mapStr["J"+strconv.Itoa(i+2)] = vid
		//是否是评分歌曲
		id := GetPlayBackRecord(info.RoomId)
		if id > 0 {
			mapStr["L"+strconv.Itoa(i+2)] = "是"
		} else {
			mapStr["L"+strconv.Itoa(i+2)] = "否"
		}
		//是否是我方设备
		_, ok := usData[info.EthernetMac]
		if ok && info.EthernetMac != "" {
			mapStr["K"+strconv.Itoa(i+2)] = "是"
		} else {
			mapStr["K"+strconv.Itoa(i+2)] = "否"
		}
	}
	return mapStr
}

func main() {
	tb := TvBox{
		TvBoxs: []string{"DreamTV Glory", "Dream_Tablet_king", "Dream TV Evolution", "Dream TV Revolution", "DreamTV", "8P"},
	}
	list, err := tb.GetTvInfoByTvBox()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("去重前数据", len(list))
	removal := duplicateRemoval(list)
	fmt.Println("去重后数据", len(removal))
	mapString := ToMapString(removal)
	et := ExcelTest{
		FileName:  "/home/ec2-user/GOTest/songinfo/特定設備資料統計.xlsx",
		SheetName: "sheet1",
	}
	et.save(mapString)

}

func GetMicrophoneInfoByRoomId(ethernetMac, curDate, deviceId string) MicrophoneAccessRecord {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	//sql := "select * from tv_login_log  where room_id = ? and create_time >= '" +curDate+" '"
	//sql := "select id, channel_id, pid, vid, tv_box, create_time from microphone_access_record  where ethernet_mac = ? and create_time like '" +curDate+" %'"
	sql := "select id, channel_id, pid, vid, tv_box, create_time from hero_ok_web.microphone_access_record  where ethernet_mac = ?and device_id=? order by id desc limit 1"
	row := util.GetDbHerookWebRead().QueryRow(sql, ethernetMac, deviceId)
	var log MicrophoneAccessRecord
	row.Scan(&log.Id, &log.ChannelId, &log.Pid, &log.Vid, &log.TvBox, &log.CreateTime)
	return log
}

type MicrophoneAccessRecord struct {
	Id         int
	ChannelId  int
	Pid        int
	Vid        int
	TvBox      string
	CreateTime time.Time
}

func GetPbRecordBYRoomId(roomId int) PlaybackRecord {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "select channel_id, pid,vid, tv_box from hero_ok_web. playback_record  where room_id = ? order by id desc limit 1"
	row := util.GetDbHerookWebRead().QueryRow(sql, roomId)
	var PbRecord PlaybackRecord
	row.Scan(&PbRecord.ChannelId, &PbRecord.Pid, &PbRecord.Vid, &PbRecord.TvBox)
	return PbRecord
}

type PlaybackRecord struct {
	Id               int
	UserId           int
	SongId           int
	RoomId           int
	YoutubeID        string
	RecordId         int
	FeedbackRecordId int
	DeviceId         string
	ChannelId        int
	TvBox            string
	WifiMac          string
	EthernetMac      string
	Version          string
	Pid              string
	Vid              string
	SerialNumber     string
	CreateTime       time.Time
}

func GetPlayBackRecord(roomId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "select id from herook.room_song_upload  where user_id = ? and (score > 0 or score2 > 0 ) limit 1 "
	//row := util.GetDbHerookRead().QueryRow(sql, roomId)
	row := util.GetDbHeroOkReab().QueryRow(sql, roomId)
	id := 0
	row.Scan(&id)
	return id
}
