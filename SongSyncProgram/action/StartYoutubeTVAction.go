package action

import (
	"SongSyncProgram/service"
	"SongSyncProgram/util"
	"strconv"
	"time"
)

func StartYoutubeTV() {
	var str string
	groupNum := service.GetGroupNum() + 1 //每次运行加1
	t1 := time.Now().Unix()
	service.SaveAutomaticUpdateLog(groupNum)
	t2 := time.Now().Unix()

	str = "第一次记录日志结束,用时:" + strconv.FormatInt(t2-t1, 10) + "秒"
	LogUtil(str)
	service.TVCopyTable()
	t3 := time.Now().Unix()
	//修改songinfoMergrTvStatus
	str = "复制表结束,用时:" + strconv.FormatInt(t3-t2, 10) + "秒"
	LogUtil(str)
	service.UpdateMergeTVSong()
	t4 := time.Now().Unix()
	str = "修改临时表结束,用时:" + strconv.FormatInt(t4-t3, 10) + "秒"
	LogUtil(str)
	//更新数据
	service.AddYoutubeTVSong()
	t5 := time.Now().Unix()
	str = "操作临时表结束,用时:" + strconv.FormatInt(t5-t4, 10) + "秒"
	LogUtil(str)
	service.SaveAutomaticUpdateLog2(groupNum)
	t6 := time.Now().Unix()
	str = "第二次记录日志结束,用时:" + strconv.FormatInt(t6-t5, 10) + "秒"
	LogUtil(str)
	//跟新txt文件
	//更新TV
	list := service.UpdateTv()
	t7 := time.Now().Unix()
	str = "更新TV表结束,用时:" + strconv.FormatInt(t7-t6, 10) + "秒"
	LogUtil(str)
	service.UpdateCacheMap()
	t8 := time.Now().Unix()
	str = "更新cacheMap文件结束,用时:" + strconv.FormatInt(t8-t7, 10) + "秒"
	LogUtil(str)
	service.SaveNewMv(list)
	t9 := time.Now().Unix()
	str = "更新mv歌曲标签结束,用时:" + strconv.FormatInt(t9-t8, 10) + "秒"
	LogUtil(str)
	t10 := time.Now().Unix()
	//更新备源文件前调用再次下载歌词文件逻辑
	for i := 0; i < 3; i++ { //重复下载3次
		AgainDownloadSongFile()
	}
	t13 := time.Now().Unix()
	str = "再次下载歌曲文件,用时:" + strconv.FormatInt(t13-t10, 10) + "秒"
	LogUtil(str)
	UpdateDoremiLastUpdateTime()
	t11 := time.Now().Unix()
	str = "无音挡歌曲Doremi时间更新,用时:" + strconv.FormatInt(t11-t13, 10) + "秒"
	LogUtil(str)
	service.UpdateBeiYuanFileSize()
	t12 := time.Now().Unix()
	str = "备源歌词文件更新结束,用时:" + strconv.FormatInt(t12-t11, 10) + "秒" + "\n" + "总用时" + strconv.FormatInt(t12-t1, 10)
	LogUtil(str)

}

func UpdateDoremiLastUpdateTime() {

	//查询无音挡的sourceId
	DownloadFailedSourceIdList := service.QueryAllYoutubeMp3Log()
	for _, sourceId := range DownloadFailedSourceIdList {
		service.UpdateDoremiLastUpdateTime(sourceId)
		//修改歌曲状态
		service.UpdateMp3LogState(sourceId)
		util.RecordLogUtil([]byte("更新无音挡的歌曲Doremi时间" + strconv.Itoa(sourceId)))
	}

}
