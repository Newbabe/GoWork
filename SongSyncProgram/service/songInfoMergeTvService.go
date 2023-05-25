package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"github.com/ozgio/strutil"
	"github.com/tidwall/gjson"
	"math"
	"strconv"
	"strings"
	"time"
)

// TVCopyTable  复制表
/*创建当天备份表并添加song_info_merge_tv表信息
清除 tv临时表信息
将song_info_merge_tv表信息添加进tv临时表中*/

const (
	//正式数据库测试时
	//suffix = "_test"
	//测试数据库运行时【正式运行】
	suffix = ""
)

func TVCopyTable() {
	//1、创建备份表并添加信息
	//当天备份表
	//newTableName := "song_info_merge_tv" + time.Now().Format("20060102")
	newTableName := "song_info_merge_tv_" + util.GetNowData() + suffix
	//dao.CreateBackupTableTestDb("song_info_merge_tv", newTableName)
	dao.CreateBackupTable("song_info_merge_tv"+suffix, newTableName)
	//清除TV临时表信息
	//dao.RemoveTmpTableDataTestDb("song_info_merge_tv_tmp")
	dao.RemoveTmpTableData("song_info_merge_tv" + suffix + "_tmp")
	//将tv曲库表数据复制给tv临时表
	//dao.TransferDataBToATestDb("song_info_merge_tv_tmp", "song_info_merge_tv")
	dao.TransferDataBToA("song_info_merge_tv"+suffix+"_tmp", "song_info_merge_tv"+suffix)

}

//修改tv临时表的status
// 获取不更新歌名及歌手的音乐信息
// 状态改为-1
// 删除youtube歌曲
/*String sql2 = "update song_info_merge_tv_tmp SET status = -1 where source  = 3 and status = 1 ";
// 删除youtube歌曲
String sql3 = "update song_info_merge_tv_tmp SET status = -99 where source  = 3 and status = -2 ";*/
func UpdateMergeTVSong() {
	//dao.UpdateSongInfoMergeTvStatus(-1, 1)
	//dao.UpdateSongInfoMergeTvStatus(-99, -2)
	dao.UpdateSongInfoMergeTvStatus("song_info_merge_tv"+suffix+"_tmp", -99, -2)
	dao.UpdateSongInfoMergeTvStatus("song_info_merge_tv"+suffix+"_tmp", -1, 1)
}

/*
创建表SongInfoTv2并复制tv临时表的数据
给TV2表添加ID主键 给TV2表天添加索引
补全TV2的数据
更新TV2歌曲状态
删除TV曲库表
将TV2改名为song_info_merge_tv
删除三天前的备份表
给song_info_youtube_inf表添加主键并启自增
		String sql9 = "alter table song_info_youtube_info drop id;";
		String sql10 = "alter table song_info_youtube_info add id int not null primary key auto_increment first;";
给song_info_youtube_info_check表添加主键并且自增
		String sql11 = "alter table song_info_youtube_info_check drop id;";
		String sql12 = "alter table song_info_youtube_info_check add id int not null primary key auto_increment first;";
*/

func UpdateTv() []model.NewSaveSongInfoMerge {
	//创建Tv2表并复制TV临时表信息
	//dao.CreateBackupTableTestDb("song_info_merge_tv_tmp", "song_info_merge_tv2")
	songInfoMergeTv2 := "song_info_merge_tv2" + suffix
	songInfoMergeTvTmp := "song_info_merge_tv" + suffix + "_tmp"
	dao.CreateBackupTable(songInfoMergeTvTmp, songInfoMergeTv2)
	//给TV2表添加各种索引
	dao.AddSongInfoMergeTV2PrimaryKey(songInfoMergeTv2)
	dao.AddSongInfoMergeTV2UNIQUEIndex(songInfoMergeTv2)
	dao.AddSongInfoMergeTV2Index(songInfoMergeTv2, "singer")
	dao.AddSongInfoMergeTV2Index(songInfoMergeTv2, "status")
	dao.AddSongInfoMergeTV2Index(songInfoMergeTv2, "dossier_name")
	dao.AddSongInfoMergeTV2Index2(songInfoMergeTv2)
	dao.AddSongInfoMergeTV2FULLTEXTIndex(songInfoMergeTv2)
	//获取要跟新的mv歌曲集合
	songInfoMerge := "song_info_merge" + suffix
	list := selectNewSaveSongInfoMerge(songInfoMerge, songInfoMergeTv2)
	dao.AddSongInfoMergeTV2BySongInfoMerge(songInfoMerge, songInfoMergeTv2)
	dao.UpdateSongInfoMergeTV2Status(songInfoMergeTv2)
	//删除tv曲库表
	//dao.DelTableTestDb("song_info_merge_tv")
	songInfoMergeTv := "song_info_merge_tv" + suffix
	dao.DelTable(songInfoMergeTv)
	//修改表名
	dao.UpdateTableName(songInfoMergeTv, songInfoMergeTv2)
	//判断3天前的备份表是否存在 如果存在就删除
	oldTableName := "song_info_merge_tv_" + util.GetNowData2() + suffix
	dao.IfExistsDelTable(oldTableName)
	//song_info_youtube_info表添加主键
	dao.DeleteIdByTableName("song_info_youtube_info")
	dao.AddIdByTableName("song_info_youtube_info")
	dao.DeleteIdByTableName("song_info_youtube_info_check")
	dao.AddIdByTableName("song_info_youtube_info_check")
	return list
}

func selectNewSaveSongInfoMerge(songInfoMerge, songInfoMergeTv2 string) []model.NewSaveSongInfoMerge {

	return dao.SelectNewSaveSongInfoMerge(songInfoMerge, songInfoMergeTv2)
}

func SaveNewMv(SaveSongInfoMergeList []model.NewSaveSongInfoMerge) {
	var ids string
	for _, saveSongInfoMerge := range SaveSongInfoMergeList {
		ids += strconv.Itoa(saveSongInfoMerge.Id) + ", "
		var num int
		fmt.Println("============新增歌曲songInfoMap", saveSongInfoMerge)
		if strings.Contains(saveSongInfoMerge.SongName, "(MV)") || strings.Contains(saveSongInfoMerge.SongName, "（MV）") {
			if saveSongInfoMerge.Source == 3 {
				tagId1 := GetSongTvTagId(saveSongInfoMerge.Id, 32)
				//TODO 此判断暂时无用
				if tagId1 == 32 {
					num += 1
				}
				fmt.Println("num:", num)
				lang := GetSongLangById(saveSongInfoMerge.Id)
				fmt.Println("将MV标签细分", saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, lang)
				switch lang {
				case 1: //国语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 40)
				case 2: //台语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 41)
				case 7: //粤语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 42)
				case 5: //英语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 43)
				case 18: //韩语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 44)
				case 4: //日语
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 45)
				default: //其他语言
					TagProcessing(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, 46)
				}
			}
		} else {
			if saveSongInfoMerge.Source == 3 {
				songIdList := dao.GetSongInfoMergeTVInfoBySongNameAndSinger(saveSongInfoMerge.SongName, saveSongInfoMerge.Singer)
				if len(songIdList) == 0 {
					var count int
					for _, i2 := range songIdList {
						tag := GetSongTvTagIdBySongId(i2)
						if tag == 32 {
							count += 1
							continue
						}
						MaxRanking := GetSongTvTagMaxRank(tag)
						fmt.Println("新添加标签--song_tv_tag : ", saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, tag)
						SaveSongTvTag(saveSongInfoMerge.Id, saveSongInfoMerge.SourceId, saveSongInfoMerge.Source, tag, MaxRanking+1)
					}
					fmt.Println("count:", count)
				}
			}
		}
	}
	fmt.Println("新增歌曲Id:", ids)

}
func TagProcessing(id, sourceId, source, tag int) {
	tagId := GetSongTvTagId(id, tag)
	if tagId == 0 {
		maxRanking := GetSongTvTagMaxRank(tag)
		SaveSongTvTag(id, sourceId, source, tag, maxRanking+1)
	}
}

func GetSongLangById(songId int) int {
	return dao.GetSongLangById(songId)
}

func AddYoutubeTVSong() {
	//doremiUpdateTimeMap保存旧数据，对比用
	doremiUpdateTimeMap := GetDoremiLastUpdateTime()
	SourceIdmMap := SelectSongInfoMergeTV()
	//System.out.println("doremiUpdateTimeMap size=" +doremiUpdateTimeMap.size());
	fmt.Println("doremiUpdateTimeMap size=", len(doremiUpdateTimeMap))
	fmt.Println("SourceIdmMap size=", len(SourceIdmMap))
	doremiNewLrcMap := make(map[int]int)
	doremiNewUpdateTimeMap := make(map[int]model.DoremiSong)
	//onlineSongMap := make(map[int]string)
	// 获取不更新歌名及歌手的音乐信息
	SongUrl := "https://api.karadoremi.net/api/song/get_song_list_by_tag.php"
	pageSize := 1000
	//noUpdateSongMAP := SelectNoUpdateSong()  未被调用
	//fmt.Println("noUpdateSongMAP:", noUpdateSongMAP)
	var num int
	var PitchInfoNum int
	for page := 0; page < 20; page++ {
		fmt.Println("addYoutubeSongDb page=", page)

		songParam := "device_id=999&platform=8&client_version=0.0.0&service=2&song_tag_id=-1&song_learn_material=4&page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(pageSize)
		result := util.HttpsPost(SongUrl, []byte(songParam))
		jsonString := string(result.Bytes)
		//获取链接状态
		status := gjson.Get(jsonString, "status").String()
		if status != "success" { //如果链接失败直接退出
			break
		}
		array := gjson.Get(jsonString, "song_list").Array()
		if len(array) == 0 { //如果歌曲集合没有数据直接退出
			break
		}

		for _, SongTvJson := range array {
			sourceIdStr := GetInt("song_id", SongTvJson.String())
			songName := GetString("song_name", SongTvJson.String())
			singer := GetString("singer_name", SongTvJson.String())
			sourceId, _ := strconv.Atoi(sourceIdStr)
			//onlineSongMap[sourceId] = strconv.Itoa(sourceId)
			IsExt, isExtSongInfoMergeTv := GetIsExt(SourceIdmMap, sourceId)
			lastUpdateDateStr := GetString("last_update_date", SongTvJson.String())
			loc, _ := time.LoadLocation("Local")
			lastUpdateDate, _ := time.ParseInLocation("2006-01-02 15:04:05", lastUpdateDateStr, loc)
			tableTime := GetLasUpdateTimeTvBySourceId(sourceId)
			var flag int
			if tableTime.IsZero() {
				flag = 1
			} else {
				if tableTime.Before(lastUpdateDate) {
					flag = 2
				} else if tableTime == lastUpdateDate {
					flag = 3
				}
			}
			//偏移量记录表添加数据【此处不需要删除偏移量表因为手机同步歌曲已经删除过】
			videoStartTimeOffset := GetString("youtube_video_start_time_offset", SongTvJson.String())
			youtubeId := GetString("youtube_video_id", SongTvJson.String())
			AddSongInfoYoutubeInfo(sourceId, youtubeId, videoStartTimeOffset)

			//0310_tdj测试限制
			//fmt.Println(time1)
			sexTagStr := GetInt("song_tag", SongTvJson.String())
			lrcChannel := getChannel(sexTagStr)
			duetModeEnabledStr := GetInt("duet_mode_enabled", SongTvJson.String())
			isDuet := getIsDuet(duetModeEnabledStr)

			//--------------------------------------更新伴奏类型--------------------------------
			var infoIntList model.SourceIdAndChannel
			if sourceId != 23033 && sourceId != 23100 && sourceId != 25860 {
				infoIntList = GetSourceIdAndLrcChannelBySongNameAndSinger(songName, singer)
				if infoIntList.SourceId != 0 {
					//4是伴奏类型为右伴左唱的歌曲
					if infoIntList.LrcChannel == 4 {
						//把状态恢复成上架状态
						UpdateStatus(sourceId)
						continue
					}
				}
			}
			//------------------------------doremi youtube 备援歌曲--------------------------------
			/*
				3647
				35342
				35343
				35382
				35383*/

			DoremiUpdateTime := GetDoremiUpdateTime(doremiUpdateTimeMap, sourceId)

			if DoremiUpdateTime.IsZero() || DoremiUpdateTime != lastUpdateDate {
				//改歌曲ID在DoremiUpdateTime表不存在时间信息或者改歌曲的最后修改时间有改动
				//3是左伴右唱，7是纯伴奏
				if lrcChannel == 3 || lrcChannel == 7 || sourceId == 35817 || sourceId == 35732 {
					str := "备源进来了<time>" + DoremiUpdateTime.String() + "<lastUpdateDate>" + lastUpdateDate.String() + "<sourceId>" + strconv.Itoa(sourceId)
					fmt.Println(str)
					util.RecordLogUtil([]byte(str))
					var DoremiSong model.DoremiSong
					DoremiSong.SourceId = sourceId
					if isDuet == 1 {
						DoremiSong.Duet = true
					}
					DoremiSong.UpdateTime = lastUpdateDate
					DoremiSong.LrcChannel = strconv.Itoa(lrcChannel)
					DoremiSong.YoutubeSongType = getYouTubeSongType(songName)
					doremiNewUpdateTimeMap[sourceId] = DoremiSong
					youtubeLearningEnavledStr := GetInt("youtube_learning_enabled", SongTvJson.String())
					youtubeLearningEnavled, _ := strconv.Atoi(youtubeLearningEnavledStr)
					duetModeEnabled, _ := strconv.Atoi(duetModeEnabledStr)
					if youtubeLearningEnavled == 1 {
						if duetModeEnabled == 1 {
							doremiNewLrcMap[sourceId] = 2
						} else {
							doremiNewLrcMap[sourceId] = 1
						}
					}
				}
			}

			if flag == 3 { //在跳过之前把未下架且为做修改的歌曲状态改为上架状态
				dao.UpdateSongInfoMergeTvStatus2(1, sourceId)
				continue
			}

			//--------------------------------拆分关键字--------------------------------------------
			songIndex := GetString("song_index", SongTvJson.String())
			singerIndex := GetString("singer_index", SongTvJson.String())
			//专辑名拼音首字母索引
			albumIndex := GetString("album_index", SongTvJson.String())
			//0112新增注⾳⾸字⺟索引
			songPhoneticIndex := GetString("song_phonetic_index", SongTvJson.String())
			singerPhoneticIndex := GetString("singer_phonetic_index", SongTvJson.String())
			//专辑名注音首字母索引
			albumPhoneticIndex := GetString("album_phonetic_index", SongTvJson.String())
			keyword, songNameKeyword, singerKeyword, albumKeyword := GetKeyWordTV(songIndex, singerIndex, songPhoneticIndex, albumIndex, singerPhoneticIndex, albumPhoneticIndex)
			//tdj  新增  关键字拆分
			keySongId := GetSearchKeyWordTvSongIdBySourceId(sourceId)
			var searchKeyWordTv model.SearchKeyWordTv
			searchKeyWordTv.SourceId = sourceId
			searchKeyWordTv.SongName = songName
			searchKeyWordTv.Singer = singer
			searchKeyWordTv.SongNameKeyWord = songNameKeyword
			searchKeyWordTv.SingerKeyWord = singerKeyword
			searchKeyWordTv.AlbumKeyWord = albumKeyword
			if keySongId > 0 {
				UpdateSearKey(searchKeyWordTv)
			} else {
				dao.SaveSearchKeyWordTv(searchKeyWordTv)
			}

			//-------------------获取字段值----------------------------
			SongInfoMergeTvTmp := getSongTVData(SongTvJson, IsExt)
			SongInfoMergeTvTmp.SearchWord = keyword
			SongInfoMergeTvTmp.UpdateTime = lastUpdateDateStr
			SongInfoMergeTvTmp.LrcChannel = lrcChannel
			SongInfoMergeTvTmp.IsDuet = isDuet
			//---------------------判断跟新还是添加------------------------
			if IsExt { //存在的情况下修改曲库表数据
				if flag == 1 {
					AddLasUpdateTimeTV(sourceId, lastUpdateDateStr)
				} else if flag == 2 {
					UpdateLasUpdateTimeTV(sourceId, lastUpdateDateStr)
				}
				//修改表数据特殊处理的字段
				oldStatus := isExtSongInfoMergeTv.Status
				updateStatus := 1
				if oldStatus == 1 || oldStatus == -1 || oldStatus == -99 {
					if oldStatus == -99 {
						updateStatus = -2
					}
				} else {
					updateStatus = oldStatus
				}
				SongInfoMergeTvTmp.Status = updateStatus
				SongInfoMergeTvTmp.SongSize = isExtSongInfoMergeTv.SongSize
				SongInfoMergeTvTmp.LrcSize = isExtSongInfoMergeTv.LrcSize
				SongInfoMergeTvTmp.LrcSize2 = isExtSongInfoMergeTv.LrcSize2
				//页面时间在已存在数据时间之前
				if lastUpdateDate.Before(isExtSongInfoMergeTv.UpdateTime) {
					SongInfoMergeTvTmp.UpdateTime = isExtSongInfoMergeTv.UpdateTime.String()
				} else {
					SongInfoMergeTvTmp.Status = 1
				}
				fmt.Println("更新的数据", SongInfoMergeTvTmp)
				num = UpdateSongInfoMergeTvTmp(SongInfoMergeTvTmp, num)
				//将高低音域更新
				id := GetSongPitchInfoBySourceId(sourceId)
				if id <= 0 {
					PitchInfoNum = SaveSongPitchInfo(SongInfoMergeTvTmp.SourceId, SongInfoMergeTvTmp.MinPitch, PitchInfoNum, SongInfoMergeTvTmp.MaxPitch)
				}
			} else {
				if flag == 1 {
					AddLasUpdateTimeTV(sourceId, lastUpdateDateStr)
				} else if flag == 2 {
					UpdateLasUpdateTimeTV(sourceId, lastUpdateDateStr)
				}
				/*2022年10月11日 在添加数据前
				1、通过歌手歌名查询临时表中 歌曲的数据
				2、判断谁否存在该歌曲的旧数据
				3、如果存在判断是否为上线状态
				4、如果是上线状态 就修改旧的歌曲状态十七为下架状态
				*/
				songInfoList := TvGetSongInfoBySongNameAndSinger(SongInfoMergeTvTmp.SongName, SongInfoMergeTvTmp.Singer)
				if len(songInfoList) > 0 {
					for _, songInfoMap := range songInfoList {
						if songInfoMap["status"] == 1 { //如果存在上线歌曲就使其下架
							TvUpdateOldSongStatue(songInfoMap["songId"])
						}
					}
				}

				//添加数据后添加音域
				num = AddSongInfoMergeTvTmp(SongInfoMergeTvTmp, num)
				fmt.Println("添加的数据", SongInfoMergeTvTmp)
				//保存高低音域
				id := GetSongPitchInfoBySourceId(sourceId)
				if id <= 0 {
					PitchInfoNum = SaveSongPitchInfo(SongInfoMergeTvTmp.SourceId, SongInfoMergeTvTmp.MinPitch, PitchInfoNum, SongInfoMergeTvTmp.MaxPitch)
				}
			}

		}
	}
	fmt.Println("临时表操作数据", num)
	fmt.Println("音域更新数据", PitchInfoNum)
	//更新doremi updatetime 和 doremi mp3      下载文件的平时同步的时候可注释掉
	fmt.Println("doremiNewUpdateTimeMap大小:", len(doremiNewUpdateTimeMap))
	util.RecordLogUtil([]byte("doremiNewUpdateTimeMap大小:" + strconv.Itoa(len(doremiNewUpdateTimeMap))))
	fmt.Println("doremiNewLrcMap:", len(doremiNewLrcMap))
	util.RecordLogUtil([]byte("doremiNewLrcMap:" + strconv.Itoa(len(doremiNewLrcMap))))

	if len(doremiNewUpdateTimeMap) != 0 {
		UpdateDoremiUpdateTimeAndFile(doremiNewUpdateTimeMap)
	}
	if len(doremiNewLrcMap) != 0 {
		updateDoreminNewLrcAndFile(doremiNewLrcMap)
	}
	/*sourceIdMap := SongWasTakenDown(onlineSongMap, SourceIdmMap)

	for sourceId := range sourceIdMap {
		dao.UpdateSongInfoMergeTvStatus2(1, sourceId)
	}*/
	updateFinallyTv()
}

func GetLasUpdateTimeTvBySourceId(sourceId int) time.Time {
	return dao.GetLasUpdateTimeTvBySourceId(sourceId)
}

func updateFinallyTv() {

	dao.UpdateFinallyTv("song_info_merge_tv" + suffix + "_tmp")
}
func AddLasUpdateTimeTV(sourceId int, lastUpdateTime string) {
	dao.AddLasUpdateTimeTV(sourceId, lastUpdateTime)
}
func UpdateLasUpdateTimeTV(sourceId int, lastUpdateTime string) {
	dao.UpdateLasUpdateTimeTV(sourceId, lastUpdateTime)
}

func SaveSongPitchInfo(SourceId, minPitch, SongPitchInfoNum, maxPitch int) int {
	minMidi := GetOriginalMidi(minPitch)
	maxMidi := GetOriginalMidi(maxPitch)
	return dao.SaveSongPitchInfo(SourceId, minPitch, maxPitch, SongPitchInfoNum, minMidi, maxMidi)
}

func GetOriginalMidi(input int) string {
	if input > 12 {
		baseMap := make(map[int]string)
		baseMap[0] = "C"
		baseMap[1] = "C#"
		baseMap[2] = "D"
		baseMap[3] = "D#"
		baseMap[4] = "E"
		baseMap[5] = "F"
		baseMap[6] = "F#"
		baseMap[7] = "G"
		baseMap[8] = "G#"
		baseMap[9] = "A"
		baseMap[10] = "A#"
		baseMap[11] = "B"
		var octaveNumber int
		octaveNumber = int(math.Floor(float64(input/12) - 1))
		//fmt.Println("111:", octaveNumber)
		//octaveNumber = (int) (Math.floor(input / 12) - 1);
		letterAccidental := baseMap[input%12]
		return letterAccidental + strconv.Itoa(octaveNumber)
	}
	return "N/A"
}

func GetSongPitchInfoBySourceId(sourceId int) int {
	return dao.GetSongPitchInfoBySourceId(sourceId)
}

func GetDoremiUpdateTime(DoremiUpdateTimeList []model.DoremiUpdateTime, sourceId int) time.Time {
	for _, DoremiUpdateTime := range DoremiUpdateTimeList {
		if DoremiUpdateTime.SourceId == sourceId {
			return DoremiUpdateTime.UpdateTime
		}
	}
	//如果没有就返回默认空时间
	return time.Time{} //赋零值
}

func getSongTVData(json gjson.Result, IsExt bool) model.SongInfoMergeTvTmp {
	var SongInfoMergeTvTmp model.SongInfoMergeTvTmp
	jsonStr := json.String()
	songName := GetString("song_name", jsonStr)
	singer := GetString("singer_name", jsonStr)
	sourceIdStr := GetInt("song_id", jsonStr)
	sourceId, _ := strconv.Atoi(sourceIdStr)
	//歌曲最低音, 以midi note number表示, C4為中央Do, 其midi note number為 60, type為number
	minPitchStr := GetInt("min_pitch", jsonStr)
	maxPitchStr := GetInt("max_pitch", jsonStr)
	minPitch, _ := strconv.Atoi(minPitchStr)
	maxPitch, _ := strconv.Atoi(maxPitchStr)
	songTimeStr := GetInt("song_time", jsonStr)
	songTime, _ := strconv.Atoi(songTimeStr)
	albumName := GetString("album_name", jsonStr)
	/*typePinYin := pinyin.NewArgs()
	typePinYin.Style = pinyin.FirstLetter //首字母
	songChinesePhoneticArray := pinyin.LazyPinyin(songName, typePinYin)
	singerChinesePhoneticArray := pinyin.LazyPinyin(singer, typePinYin)*/
	songChinesePhonetic := SongNamePinYinProcessing2(songName)
	singerChinesePhonetic := SongNamePinYinProcessing2(singer)
	//是否有全曲LRC檔可評分
	var hasLrc int
	var lrcSize int
	var lrc2Size int
	hasLrcStr := GetInt("has_lrc", jsonStr)
	hasLrc, _ = strconv.Atoi(hasLrcStr)
	//是否有合唱模式
	duetModeEnabled := GetInt("duet_mode_enabled", jsonStr)
	lrcSize, lrc2Size, hasLrc = getLrcSizeAndLrc2Size(hasLrc, duetModeEnabled)
	languageStr := GetInt("language", jsonStr)
	language := getLanguage(languageStr)
	sexTagStr := GetInt("song_tag", jsonStr)
	sex := getSex(sexTagStr)
	isDuet := getIsDuet(duetModeEnabled)
	youtubeId := GetString("youtube_video_id", jsonStr)
	//videoStartTimeOffset := GetString("youtube_video_start_time_offset", jsonStr)

	wordPart := GetWordPart(songName)

	youtubeSongType := getYouTubeSongType(songName)

	lrcChannel := getChannel(sexTagStr)
	SongInfoMergeTvTmp.SourceId = sourceId
	SongInfoMergeTvTmp.Source = 3
	SongInfoMergeTvTmp.WordPart = wordPart
	SongInfoMergeTvTmp.SongName = songName
	SongInfoMergeTvTmp.Singer = singer
	SongInfoMergeTvTmp.MinPitch = minPitch
	SongInfoMergeTvTmp.MaxPitch = maxPitch
	//	fmt.Println("是否存在", IsExt)
	if !IsExt {
		SongInfoMergeTvTmp.SongNamePhonetic = GetZhuYinNew(songName)
		//fmt.Println(SongInfoMergeTvTmp.SongNamePhonetic)
		SongInfoMergeTvTmp.SingerNamePhonetic = GetZhuYinNew(singer)
		//	fmt.Println(SongInfoMergeTvTmp.SongNamePhonetic)
	} else {
		SongInfoMergeTvTmp.SongNamePhonetic = ""
		SongInfoMergeTvTmp.SingerNamePhonetic = ""
	}
	SongInfoMergeTvTmp.Album = albumName
	SongInfoMergeTvTmp.LineWriter = ""
	SongInfoMergeTvTmp.SongWrite = ""
	SongInfoMergeTvTmp.Gender = sex
	SongInfoMergeTvTmp.Lang = language
	SongInfoMergeTvTmp.Year = 2018
	SongInfoMergeTvTmp.DossierName = youtubeId
	SongInfoMergeTvTmp.IsDuet = isDuet
	SongInfoMergeTvTmp.SongSize = 0
	SongInfoMergeTvTmp.LrcSize = lrcSize
	SongInfoMergeTvTmp.BackgroundImageSize = "0"
	SongInfoMergeTvTmp.IconSize = 0
	SongInfoMergeTvTmp.SongTime = songTime
	SongInfoMergeTvTmp.Status = 1
	SongInfoMergeTvTmp.LrcType = hasLrc

	SongInfoMergeTvTmp.Intonation = 1
	SongInfoMergeTvTmp.CreateDate = util.GetNowTime()
	SongInfoMergeTvTmp.MvSize = 0
	SongInfoMergeTvTmp.SongPath = "youtube"
	firstNameStr, _ := strutil.Substring(songName, 0, 1)
	firstSingerStr, _ := strutil.Substring(singer, 0, 1)
	SongInfoMergeTvTmp.SongNameStrokeNum = util.GetStrokeCount(firstNameStr)
	//	fmt.Println("SongNameStrokeNum", SongInfoMergeTvTmp.SongNameStrokeNum)
	SongInfoMergeTvTmp.SingerStrokeNum = util.GetStrokeCount(firstSingerStr)
	//	fmt.Println("SongNameStrokeNum", SongInfoMergeTvTmp.SingerStrokeNum)
	SongInfoMergeTvTmp.SongNameSimple = util.Transform(songName, "zh-Hans")
	SongInfoMergeTvTmp.SingerSimple = util.Transform(singer, "zh-Hans")
	SongInfoMergeTvTmp.SongVersion = "1"
	SongInfoMergeTvTmp.LrcVersion = "1"
	SongInfoMergeTvTmp.LrcChannel = lrcChannel
	SongInfoMergeTvTmp.LrcSize2 = lrc2Size
	SongInfoMergeTvTmp.SongChinesePhonetic = songChinesePhonetic
	SongInfoMergeTvTmp.SingerChinesePhonetic = singerChinesePhonetic

	SongInfoMergeTvTmp.YouTubeSongType = youtubeSongType
	return SongInfoMergeTvTmp
}

func AddSongInfoMergeTvTmp(tmp model.SongInfoMergeTvTmp, num int) int {
	songInfoMerge := "song_info_merge" + suffix
	return dao.AddSongInfoMergeTvTmp(songInfoMerge, tmp, num)

}
func UpdateSongInfoMergeTvTmp(tmp model.SongInfoMergeTvTmp, num int) int {
	songInfoMergeTvTmp := "song_info_merge_tv" + suffix + "_tmp"
	return dao.UpdateSongInfoMergeTvTmp(songInfoMergeTvTmp, tmp, num)
}

func GetKeyWordTV(songIndex, singerIndex, songPhoneticIndex, albumIndex, singerPhoneticIndex, albumPhoneticIndex string) (string, string, string, string) {
	var keyword string
	var songNameKeyword string
	var singerKeyword string
	var albumKeyword string

	if songIndex != "" {
		keyword += songIndex
		songNameKeyword += songIndex
	}
	if singerIndex != "" {
		keyword += " " + singerIndex
		singerKeyword += singerIndex
	}
	if albumIndex != "" {
		keyword += " " + albumIndex
		albumKeyword += albumIndex
	}

	if songPhoneticIndex != "" {
		keyword += " " + songPhoneticIndex
		songNameKeyword += " " + songPhoneticIndex
	}
	if singerPhoneticIndex != "" {
		keyword += " " + singerPhoneticIndex
		singerKeyword += " " + singerPhoneticIndex
	}
	if albumPhoneticIndex != "" {
		keyword += " " + albumPhoneticIndex
		albumKeyword += " " + albumPhoneticIndex
	}
	return keyword, songNameKeyword, singerKeyword, albumKeyword
}

// 获取添加数据
func SelectSongInfoMergeTV() []model.SongInfoMerger {
	return dao.SelectSongInfoMergeTV()
}
func SelectNoUpdateSong() []model.UpdateSongInfoStatus {
	return dao.SelectNoUpdateSong()
}
func GetSourceIdAndLrcChannelBySongNameAndSinger(songName, Singer string) model.SourceIdAndChannel {
	songInfoMergeTv := "song_info_merge_tv" + suffix

	return dao.GetSourceIdAndLrcChannelBySongNameAndSinger(songInfoMergeTv, songName, Singer)
}
func UpdateStatus(sourceId int) {
	//fmt.Println("更新状态")
	songInfoMergeTvTmp := "song_info_merge_tv" + suffix + "_tmp"
	dao.UpdateStatus(songInfoMergeTvTmp, sourceId)
}
func Update(sourceId, lrcSize1, lrcSize2, songSize, songVersion int) {
	songInfoMergeTv := "song_info_merge_tv" + suffix

	dao.Update(songInfoMergeTv, sourceId, lrcSize1, lrcSize2, songSize, songVersion)
}
func TvGetSongInfoBySongNameAndSinger(songName, singer string) []map[string]int {
	tableName := "song_info_merge_tv_tmp"
	return dao.GetSongInfoBySongNameAndSinger(songName, singer, tableName)
}
func TvUpdateOldSongStatue(songId int) {
	tableName := "song_info_merge_tv_tmp"
	dao.UpdateOldSongStatue(songId, tableName)
}
