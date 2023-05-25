package malysiaService

import (
	"SongSyncProgram/malysiaDao"
	"testing"
)

/*func TestSaveAutomaticUpdateLog(t *testing.T) {
	getGroupNum := malysiaDao.GetGroupNum() + 1
	fmt.Println("getGroupNum", getGroupNum)
	now := time.Now().Format("2006-01_02 15:04:05")
	malysiaDao.SaveAutomaticUpdateLog(automaticUpdateLogTv, now, 1, malysiaDao.GetAdjustmentValue(1), getGroupNum)

}*/

/*func TestCopyTable(t *testing.T) {
	now := time.Now().Format("20060102")
	backupTable := "song.song_info_merge_copy_" + now + testTable
	malysiaDao.CreateTable(backupTable, songInfoMerge)
	malysiaDao.TruncateTable(songInfoMergeTmp)
	malysiaDao.TruncateTable(songInfoMergeBakCopy)
	malysiaDao.CopyTable(songInfoMerge, songInfoMergeTmp)
	malysiaDao.CopyTable(songInfoMerge, songInfoMergeBakCopy)
}*/
/*func TestAddSongInfoMergeTmp(t *testing.T) {
	// 2.读取台湾song_info source=1的数据（台湾）

	// 3.将第二步的结果插入到song_info_merge_tmp（马来西亚）

	//DelMalysiaMergeSongDb()

	twSongInfoList := SelectTwSongInfoDb()

	var i int
	for _, twSongInfo := range twSongInfoList {

		i += malysiaDao.AddSongInfoMergeTmp(songInfoMergeTmp, twSongInfo)
		//记录添加成功的songId
	}
	fmt.Println("添加成功的个数", i)
}*/

/*func TestUpdateMergeSongDb(t *testing.T) {

	malysiaDao.UpdateMergeSongDb(songInfoMergeTmp)

}*/

func TestAddYoutubeSongDb(t *testing.T) {
	//清除偏移量表
	//malysiaDao.ClearYoutubeOffest(songInfoYoutubeInfo)
	//malysiaDao.AddSongInfoYoutubeInfo(1111, songInfoYoutubeInfo, strconv.Itoa(11111), strconv.Itoa(1111))
	/*SongInfoMergeTmp := model.SongInfoMergeTmp{
		SourceId:              1,
		Source:                2,
		WordPart:              3,
		LrcChannel:            4,
		LrcSize2:              5,
		Gender:                6,
		Lang:                  7,
		Year:                  7,
		IsDuet:                8,
		SongSize:              8,
		LrcSize:               8,
		IconSize:              8,
		SongTime:              8,
		Status:                1,
		LrcType:               8,
		YouTubeSongType:       8,
		Intonation:            8,
		MvSize:                8,
		BackgroundImageSize:   "修改测试",
		DossierName:           "修改测试",
		UpdateTime:            "修改测试",
		CreateDate:            "修改测试",
		SongPath:              "修改测试",
		SongNameStrokeNum:     "修改测试",
		SingerStrokeNum:       "修改测试",
		SongNameSimple:        "修改测试",
		SingerSimple:          "修改测试",
		SongVersion:           "修改测试",
		LrcVersion:            "修改测试",
		SongChinesePhonetic:   "修改测试",
		SingerChinesePhonetic: "修改测试",
		SearchWord:            "修改测试",
		SongName:              "修改测试",
		SongNamePhonetic:      "修改测试",
		Singer:                "修改测试",
		SingerNamePhonetic:    "修改测试",
		Album:                 "修改测试",
		LineWriter:            "修改测试",
		SongWrite:             "1",
	}*/
	//num := malysiaDao.AddSongInfoMergeTmp2(songInfoMergeTmp, SongInfoMergeTmp)
	//num := malysiaDao.UpdateSongInfoMergeTmp(songInfoMergeTmp, SongInfoMergeTmp)
	//fmt.Print("num", num)
	//malysiaDao.AddLasUpdateTime(1, "2022-12-21 17:13:01")
	//malysiaDao.UpdateLasUpdateTime(1, "2022-12-21 17:20:46")
	//malysiaDao.UpdateMergeSongDb2(songInfoMergeTmp, 1)
	//songInfoList := malysiaDao.GetSongInfoBySongNameAndSinger("流浪到台北", "黃西田", songInfoMerge)
	//fmt.Println(songInfoList)
	/*for _, songInfoMap := range songInfoList {
		malysiaDao.UpdateOldSongStatue(songInfoMap["songId"], songInfoMergeTmp)
	}*/
	//
	/*SaveAutomaticUpdateLog()
	SaveAutomaticUpdateLog2()*/
}

func TestEndTable(t *testing.T) {
	//malysiaDao.TruncateTable(songInfoMerge)
	//malysiaDao.AddTableBByTableA(songInfoMergeTmp, songInfoMerge)
	//删除备份表
	//backupTable := "song.song_info_merge_copy_" + time.Now().AddDate(0, 0, -3).Format("20060102") + testTable
	//malysiaDao.DropTable(backupTable)
	//删除Id字段
	//malysiaDao.DropId(songInfoYoutubeInfo)
	//添加Id主键自增
	malysiaDao.AddPrimaryKey(songInfoYoutubeInfo)
}
