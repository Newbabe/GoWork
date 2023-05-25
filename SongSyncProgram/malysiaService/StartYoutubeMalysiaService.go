package malysiaService

import (
	"SongSyncProgram/malysiaDao"
	"SongSyncProgram/model"
	"SongSyncProgram/service"
	"SongSyncProgram/util"
	"fmt"
	"github.com/ozgio/strutil"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

const (
	testTable = "_test"
)
const (
	songInfoMergeTmp     = "song.song_info_merge_tmp" + testTable
	automaticUpdateLogTv = "song.automatic_update_log_tv" + testTable
	songInfoMergeBakCopy = "song.song_info_merge_bak_copy" + testTable
	songInfoMerge        = "song.song_info_merge" + testTable
	songInfoYoutubeInfo  = "song.song_info_youtube_info" + testTable
)

func SaveAutomaticUpdateLog() {
	getGroupNum := malysiaDao.GetGroupNum() + 1
	now := time.Now().Format("2006-01_02 15:04:05")
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 1, malysiaDao.GetAdjustmentValue(1), getGroupNum)
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 2, malysiaDao.GetAdjustmentValue(2), getGroupNum)
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 3, malysiaDao.GetAdjustmentValue(3), getGroupNum)
}
func SaveAutomaticUpdateLog2() {
	getGroupNum := malysiaDao.GetGroupNum() + 1
	now := time.Now().Format("2006-01_02 15:04:05")
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 4, malysiaDao.GetAdjustmentValue(1), getGroupNum)
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 5, malysiaDao.GetAdjustmentValue(2), getGroupNum)
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 6, malysiaDao.GetAdjustmentValue(3), getGroupNum)
}

func CopyTable() {
	now := time.Now().Format("20060102")
	backupTable := "song.song_info_merge_copy_" + now + testTable
	malysiaDao.CreateTable(backupTable, songInfoMerge)
	malysiaDao.TruncateTable(songInfoMergeTmp)
	malysiaDao.TruncateTable(songInfoMergeBakCopy)
	malysiaDao.CopyTable(songInfoMerge, songInfoMergeTmp)
	malysiaDao.CopyTable(songInfoMerge, songInfoMergeBakCopy)
}
func DelMalysiaMergeSongDb() {

	malysiaDao.DelMalysiaMergeSongDb(songInfoMergeTmp)
}

func SelectTwSongInfoDb() []model.SongInfoMergeTvTmp {
	return malysiaDao.SelectTwSongInfoDb()
}
func AddSongInfoMergeTmp() []int {
	var songIdList []int
	// 1.删除song_info_merge_tmp source=1的数据（马来西亚）
	DelMalysiaMergeSongDb()
	// 2.读取台湾song_info source=1的数据（台湾）
	twSongInfoList := SelectTwSongInfoDb()
	// 3.将第二步的结果插入到song_info_merge_tmp（马来西亚）
	for _, twSongInfo := range twSongInfoList {
		//AddSongInfoMergeTmpsong.`song_info_merge_tmp`

		//fmt.Println("songId:=", twSongInfo.Id)

		songId := malysiaDao.AddSongInfoMergeTmp(songInfoMergeTmp, twSongInfo)
		//记录添加成功的songId
		songIdList = append(songIdList, songId)
	}
	return songIdList
}

func UpdateMergeSongDb() {
	malysiaDao.UpdateMergeSongDb(songInfoMergeTmp)
}

// 跟新数据逻辑判断
func UpdateData(TableLastUpdateTime, lastUpdateTime time.Time) (tableFlag int) {

	if TableLastUpdateTime.IsZero() { //判断获取到的时间是否是为空如果为空就添加数据
		//	fmt.Println("sourceId", sourceId)
		return 1
	} else {
		if TableLastUpdateTime.Before(lastUpdateTime) {
			return 2
		} else if TableLastUpdateTime == lastUpdateTime { //如果表中存在数据就判断数据改歌曲的lastupdate时间与页面获取的时间是或否一致如果一一致接跳出循环
			return 3
		} else { //其他情况跟新数据
			return 2
		}
	}
}

// 返回处理后的字段以及以存在的数据和存在标志
func getTheSongData(json gjson.Result, exist bool) model.SongInfoMergeTmp {
	jsonStr := json.String()
	var SongInfoMerge model.SongInfoMergeTmp

	songName := service.GetString("song_name", jsonStr)
	singer := service.GetString("singer_name", jsonStr)
	songTimeStr := service.GetInt("song_time", jsonStr)
	songTime, _ := strconv.Atoi(songTimeStr)
	albumName := service.GetString("album_name", jsonStr)

	songChinesePhonetic := service.SongNamePinYinProcessing2(songName)

	singerChinesePhonetic := service.SongNamePinYinProcessing2(singer)

	hasLrcStr := service.GetInt("has_lrc", jsonStr)
	hasLrc, _ := strconv.Atoi(hasLrcStr)
	duetModeEnabled := service.GetInt("duet_mode_enabled", jsonStr)
	var lrcSize int
	var lrc2Size int
	lrcSize, lrc2Size, hasLrc = getLrcSizeAndLrc2Size(hasLrc, duetModeEnabled)
	languageStr := service.GetInt("language", jsonStr)
	language := getLanguage(languageStr)
	sexTag := service.GetInt("song_tag", jsonStr)
	sex := getSex(sexTag)

	songIndex := service.GetString("song_index", jsonStr)
	singerIndex := service.GetString("singer_index", jsonStr)
	//albumIndex := getString("album_index", jsonStr)
	songPhoneticIndex := service.GetString("song_phonetic_index", jsonStr)
	singerPhoneticIndex := service.GetString("singer_phonetic_index", jsonStr)

	keyword := getKeyWord(songIndex, singerIndex, songPhoneticIndex, singerPhoneticIndex)

	YouTubeSongType := getYouTubeSongType(songName)

	SongInfoMerge.Source = 3

	SongInfoMerge.WordPart = service.GetWordPart(songName)

	SongInfoMerge.SongName = songName

	if !exist { //如果时不存在的歌曲就给新歌曲获取注音
		//	fmt.Println("需要获取注音的SongName和Singer", songName, singer)
		SongInfoMerge.SongNamePhonetic = service.GetZhuYinNew(songName)
		SongInfoMerge.SingerNamePhonetic = service.GetZhuYinNew(singer)
	} else {
		SongInfoMerge.SongNamePhonetic = ""
		SongInfoMerge.SingerNamePhonetic = ""
	}

	SongInfoMerge.Singer = singer
	SongInfoMerge.Album = albumName
	SongInfoMerge.LineWriter = ""
	SongInfoMerge.SongWrite = ""
	SongInfoMerge.Gender = sex
	SongInfoMerge.Lang = language
	SongInfoMerge.Year = 2018

	SongInfoMerge.IsDuet = getIsDuet(duetModeEnabled)
	SongInfoMerge.SongSize = 0
	SongInfoMerge.LrcSize = lrcSize
	SongInfoMerge.BackgroundImageSize = "0"
	SongInfoMerge.IconSize = 0
	SongInfoMerge.SongTime = songTime
	SongInfoMerge.Status = 1
	LrcType := hasLrc
	SongInfoMerge.LrcType = LrcType

	SongInfoMerge.Intonation = 1
	SongInfoMerge.CreateDate = util.GetNowTime()
	SongInfoMerge.MvSize = 0
	SongInfoMerge.SongPath = "youtube"
	firstNameStr, _ := strutil.Substring(songName, 0, 1)
	firstSingerStr, _ := strutil.Substring(singer, 0, 1)
	SongInfoMerge.SongNameStrokeNum = util.GetStrokeCount(firstNameStr)
	SongInfoMerge.SingerStrokeNum = util.GetStrokeCount(firstSingerStr)
	SongInfoMerge.SongNameSimple = util.Transform(songName, "zh-Hans")
	SongInfoMerge.SingerSimple = util.Transform(singer, "zh-Hans")
	SongInfoMerge.SongVersion = "1"
	SongInfoMerge.LrcVersion = "1"
	SongInfoMerge.LrcChannel = getChannel(sexTag)
	SongInfoMerge.LrcSize2 = lrc2Size
	SongInfoMerge.SongChinesePhonetic = songChinesePhonetic
	SongInfoMerge.SingerChinesePhonetic = singerChinesePhonetic
	SongInfoMerge.SearchWord = keyword
	SongInfoMerge.YouTubeSongType = YouTubeSongType
	/*t8 := time.Now().UnixNano() / 1e6
	fmt.Println("赋值用时", t8-t7, "毫秒")*/
	return SongInfoMerge
}
func getLrcSizeAndLrc2Size(hasLrc int, duetModeEnabledStr string) (int, int, int) {
	duetModeEnabled, _ := strconv.Atoi(duetModeEnabledStr)

	var lrcSize int
	var lrc2Size int
	if hasLrc == 1 || duetModeEnabled == 1 {
		hasLrc = 0
		lrcSize = 1
	} else {
		hasLrc = 4
	}
	if duetModeEnabled == 1 {
		lrc2Size = 1
	}
	return lrcSize, lrc2Size, hasLrc
}
func getLanguage(languageStr string) int {
	language, _ := strconv.Atoi(languageStr)
	if language == 1 || language == 2 {
		language = 1
	} else if language == 3 {
		language = 5
	} else if language == 4 {
		language = 4
	} else if language == 5 {
		language = 2
	} else if language == 6 {
		language = 7
	} else if language == 7 {
		language = 18
	} else if language == 8 {
		language = 20
	} else if language == 9 {
		language = 3
	} else if language == 10 {
		language = 19
	} else if language == 11 {
		language = 21
	}
	return language
}
func getKeyWord(songIndex, singerIndex, songPhoneticIndex, singerPhoneticIndex string) string {
	var keyword string
	if songIndex != "" {
		keyword += songIndex
	}
	if singerIndex != "" {
		keyword += " " + singerIndex
	}

	if songPhoneticIndex != "" {
		keyword += " " + songPhoneticIndex
	}
	if singerPhoneticIndex != "" {
		keyword += " " + singerPhoneticIndex
	}
	return keyword
}

func getSex(sexTagStr string) int {
	sexTag, _ := strconv.Atoi(sexTagStr)
	var sex int
	if (sexTag&1) != 1 && (sexTag&2) != 2 {
		if (sexTag & 4) == 4 {
			sex = 3
		} else {
			sex = 0
		}
	}
	if (sexTag & 1) == 1 {
		sex = 1
	}
	if (sexTag & 2) == 2 {
		sex = 2
	}
	return sex
}
func getIsDuet(duetModeEnabledStr string) int {
	var isDuet int
	duetModeEnabled, _ := strconv.Atoi(duetModeEnabledStr)
	if duetModeEnabled == 1 {
		isDuet = 1
	}
	return isDuet
}
func getChannel(SexTagStr string) int {
	sexTag, _ := strconv.Atoi(SexTagStr)
	lrcChannel := 1
	if (sexTag & 32768) == 32768 {
		lrcChannel = 3
	} else if (sexTag & 16384) == 16384 {
		lrcChannel = 7
	} else {
		lrcChannel = 5
	}
	return lrcChannel
}
func getYouTubeSongType(songName string) int {

	youtubeSongType := 1 // youtube伴奏类型1练唱，2MV,3纯伴奏
	if strings.Contains(songName, "(MV)") || strings.Contains(songName, "（MV）") {
		youtubeSongType = 2
	} else if strings.Contains(songName, "(純伴奏)") || strings.Contains(songName, "（純伴奏）") || strings.Contains(songName, "(Karaoke)") || strings.Contains(songName, "（Karaoke）") {
		youtubeSongType = 3
	} else if strings.Contains(songName, "(練唱版)") {
		youtubeSongType = 4
	}
	return youtubeSongType
}
func AddYoutubeSongDb() {
	//清除偏移量表
	ClearYoutubeOffest()
	SongInfoMergeList := malysiaDao.SelectSongInfoMergeDb(songInfoMerge)
	songUrl := "https://api.karadoremi.net/api/song/get_song_list_by_tag.php"
	pageSize := 1000
	onlineSongMap := make(map[int]string)
	var i, num int

	for page := 0; page < 20; page++ {

		//	t1 := time.Now().UnixNano() / 1e6
		fmt.Println("AddYoutubeSongDb Page=", page)
		songParam := "device_id=999&platform=8&client_version=0.0.0&service=3&song_tag_id=-1&song_learn_material=4&page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(pageSize)
		result := util.HttpsPost(songUrl, []byte(songParam))
		jsonString := string(result.Bytes)
		//处理得到的数据
		status := gjson.Get(jsonString, "status").String()
		//判断请求是否成功
		if status != "success" { //如果请求退出循环
			break
		}
		array := gjson.Get(jsonString, "song_list").Array()
		if len(array) == 0 {
			break
		}
		//循环songList
		for _, SongJson := range array {

			sourceIdStr := service.GetInt("song_id", SongJson.String())
			sourceId, _ := strconv.Atoi(sourceIdStr)
			onlineSongMap[sourceId] = strconv.Itoa(sourceId)
			youtubeId := service.GetString("youtube_video_id", SongJson.String())
			videoStartTimeOffset := service.GetString("youtube_video_start_time_offset", SongJson.String())
			//更新偏移量
			var exist bool
			var existSongInfoMerge model.SongInfoMerger
			//t11 := time.Now().UnixNano() / 1e6
			i += AddSongInfoYoutubeInfo(sourceId, youtubeId, videoStartTimeOffset)
			//t2 := time.Now().UnixNano() / 1e6
			//fmt.Println("AddSongInfoYoutubeInfo t2-t11", t2-t11)
			exist, existSongInfoMerge = service.GetIsExt(SongInfoMergeList, sourceId)
			//t3 := time.Now().UnixNano() / 1e6
			//fmt.Println("t3-t2 GetIsExt:", t3-t2)

			//从记录表中获取修改时间
			tableLastUpdateTime := malysiaDao.GetLastUpdateTimeBySourceId(sourceId)
			//t4 := time.Now().UnixNano() / 1e6
			//fmt.Println("GetLastUpdateTimeBySourceId t4-t3", t4-t3)
			lastUpdateDate := service.GetString("last_update_date", SongJson.String())
			//将从页面获取的时间改为CST【中国时间】时间
			loc, err := time.LoadLocation("Local")
			if err != nil {
				return
			}
			lastUpdateTime, err := time.ParseInLocation("2006-01-02 15:04:05", lastUpdateDate, loc)
			var tableFlage int
			/*
				tableFlage 为1 添加表数据
				tableFlage 为2 更新表数据表数据
			*/

			tableFlage = UpdateData(tableLastUpdateTime, lastUpdateTime)
			if sourceId == 35594 ||
				sourceId == 36099 ||
				sourceId == 36044 ||
				sourceId == 36043 ||
				sourceId == 36100 {
				fmt.Println("tableFlage", tableFlage, "source", sourceId)
			}
			if tableFlage == 3 {
				continue
			}
			//	t5 := time.Now().UnixNano() / 1e6
			SongInfoMergeTmp := getTheSongData(SongJson, exist)
			//	t6 := time.Now().UnixNano() / 1e6
			//	fmt.Println("getTheSongData t6-t5", t6-t5)
			/*if ISSourceId(sourceId) {
				fmt.Println("打印测试数据", SongInfoMergeTmp)
			}*/
			//这些数据在上面以获取无需二次获取
			SongInfoMergeTmp.SourceId = sourceId
			SongInfoMergeTmp.DossierName = youtubeId
			SongInfoMergeTmp.UpdateTime = lastUpdateDate
			if exist { //如果歌曲数据存在就进行修改操作
				if sourceId == 35594 ||
					sourceId == 36099 ||
					sourceId == 36044 ||
					sourceId == 36043 ||
					sourceId == 36100 {
					fmt.Println("修改数据", sourceId)
				}
				if tableFlage == 1 { //向记录表中添加数据
					//	t7 := time.Now().UnixNano() / 1e6
					malysiaDao.AddLasUpdateTime(sourceId, lastUpdateDate)
					//	t8 := time.Now().UnixNano() / 1e6
					//	fmt.Println("AddLasUpdateTime  t8-t7", t8-t7)
					num += updateSongInfoMergeTmp(existSongInfoMerge, SongInfoMergeTmp, lastUpdateTime, num)
					//	t9 := time.Now().UnixNano() / 1e6
					//	fmt.Println("updateSongInfoMergeTmp t9-t8 ", t9-t8)
				} else if tableFlage == 2 {
					//	t7 := time.Now().UnixNano() / 1e6
					//如果标志位为 2 表示该歌曲存在记录 就跟心该歌曲的记录时间
					malysiaDao.UpdateLasUpdateTime(sourceId, lastUpdateDate)
					//	t8 := time.Now().UnixNano() / 1e6
					//	fmt.Println("AddLasUpdateTime  t8-t7", t8-t7)
					num += updateSongInfoMergeTmp(existSongInfoMerge, SongInfoMergeTmp, lastUpdateTime, num)
					//	t9 := time.Now().UnixNano() / 1e6
					//	fmt.Println("updateSongInfoMergeTmp t9-t8 ", t9-t8)
				} else {
					//将未做修改的歌曲改为上架转态
					//malysiaDao.UpdateMergeSongDb2(songInfoMergeTmp, sourceId)
					continue
				}
			} else { //否则就进行添加操作【添加的时候需要判断添加的歌曲是否存在旧版本 如果存在旧版本就将旧版本下架】
				//	t7 := time.Now().UnixNano() / 1e6
				if sourceId == 35594 ||
					sourceId == 36099 ||
					sourceId == 36044 ||
					sourceId == 36043 ||
					sourceId == 36100 {
					fmt.Println("添加数据", sourceId)
				}
				songInfoList := malysiaDao.GetSongInfoBySongNameAndSinger(SongInfoMergeTmp.SongName, SongInfoMergeTmp.Singer, songInfoMerge)
				//t8 := time.Now().UnixNano() / 1e6
				//	fmt.Println("updateSongInfoMergeTmp t8-t7 ", t8-t7)
				if len(songInfoList) > 0 {
					for _, songInfoMap := range songInfoList {
						if songInfoMap["status"] == 1 { //如果存在上线歌曲就使其下架
							//t9 := time.Now().UnixNano() / 1e6
							malysiaDao.UpdateOldSongStatue(songInfoMap["songId"], songInfoMerge)
							//	t10 := time.Now().UnixNano() / 1e6
							//	fmt.Println("UpdateOldSongStatue  t10-t9", t10-t9)
						}
					}
				}
				//添加临时表数据
				//	t13 := time.Now().UnixNano() / 1e6
				num += malysiaDao.AddSongInfoMergeTmp2(songInfoMergeTmp, SongInfoMergeTmp)
				fmt.Println("添加数据", SongInfoMergeTmp)
				//	t12 := time.Now().UnixNano() / 1e6
				//	fmt.Println("AddSongInfoMergeTmp2  t12-t13", t12-t13)
				if tableFlage == 1 { //向记录表中添加数据
					malysiaDao.AddLasUpdateTime(sourceId, lastUpdateDate)
				} else if tableFlage == 2 { //如果标志位为 2 表示该歌曲存在记录 就跟心该歌曲的记录时间
					malysiaDao.UpdateLasUpdateTime(sourceId, lastUpdateDate)
				}
			}
		}
		//t2 := time.Now().UnixNano() / 1e6
		//	fmt.Println("一个循环用时", t2-t1)
	}
	sourceIdMap := service.SongWasTakenDown(onlineSongMap, SongInfoMergeList)
	for sourceId := range sourceIdMap {
		malysiaDao.UpdateMergeSongDb2(songInfoMergeTmp, sourceId)
	}
	fmt.Println("下架歌曲数量", len(sourceIdMap))
	fmt.Println("总共插入的offset数据量为--", i)
	fmt.Println("操作数据数据量为--", num)

	SaveAutomaticUpdateLog2()
}
func AddSongInfoYoutubeInfo(sourceId int, youtubeId, videoStartTimeOffset string) int {

	return malysiaDao.AddSongInfoYoutubeInfo(sourceId, songInfoYoutubeInfo, youtubeId, videoStartTimeOffset)
}

func updateSongInfoMergeTmp(existSongInfoMerge model.SongInfoMerger, SongInfoMergeTmp model.SongInfoMergeTmp, lastUpdateTime time.Time, num int) int {
	oldStatus := existSongInfoMerge.Status
	updateStatus := 1
	if oldStatus == 1 || oldStatus == -1 || oldStatus == -99 {
		if oldStatus == -99 {
			updateStatus = -2
		} else {
			updateStatus = oldStatus
		}
	}
	SongInfoMergeTmp.Status = updateStatus
	SongInfoMergeTmp.SongSize = existSongInfoMerge.SongSize
	SongInfoMergeTmp.LrcSize = existSongInfoMerge.LrcSize
	SongInfoMergeTmp.LrcSize2 = existSongInfoMerge.LrcSize2
	if lastUpdateTime.Before(existSongInfoMerge.UpdateTime) { //如果当前时间在修改时间之前
		SongInfoMergeTmp.UpdateTime = existSongInfoMerge.UpdateTime.Format("2006-01-02 15-04-05")
	} else {
		SongInfoMergeTmp.Status = 1
	}
	num += malysiaDao.UpdateSongInfoMergeTmp(songInfoMergeTmp, SongInfoMergeTmp)

	fmt.Println("修改的数据", SongInfoMergeTmp)
	return num
}

func ClearYoutubeOffest() {
	malysiaDao.ClearYoutubeOffest(songInfoYoutubeInfo)
}
func EndTable() {
	//截断表
	malysiaDao.TruncateTable(songInfoMerge)
	malysiaDao.AddTableBByTableA(songInfoMergeTmp, songInfoMerge)
	//删除备份表
	backupTable := "song.song_info_merge_copy_" + time.Now().AddDate(0, 0, -3).Format("20060102") + testTable
	malysiaDao.DropTable(backupTable)
	//删除Id字段
	malysiaDao.DropId(songInfoYoutubeInfo)
	//添加Id主键自增
	malysiaDao.AddPrimaryKey(songInfoYoutubeInfo)
}
