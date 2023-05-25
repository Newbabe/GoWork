package action

import (
	"SongSyncProgram/service"
	"SongSyncProgram/util"
	"fmt"
	"strconv"
	"time"
)

func StartYoutubePhone() {
	var str string
	t1 := util.GetNowTime()
	tt1 := time.Now().Unix()
	str = "手机歌曲同步开始:" + t1
	LogUtil(str)
	groupNum := service.GetGroupNum() + 1 //每次运行加1
	// 先统计平差值表的信息
	// 插入log表
	str = "开始记录日志:" + util.GetNowTime()
	LogUtil(str)
	service.SaveAutomaticUpdateLog(groupNum)
	// 复制表
	str = "开始复制表:" + util.GetNowTime()
	LogUtil(str)
	service.CopyTable()
	//修改临时表数据状态
	str = "开始修改临时表的数据:" + util.GetNowTime()
	LogUtil(str)
	service.UpdateMergeTmpStatus()
	//临时表添加数据【】
	str = "开始添加数据:" + util.GetNowTime()
	LogUtil(str)
	service.AddYoutubeSong()
	str = "第二次记录日志:" + util.GetNowTime()
	LogUtil(str)
	service.SaveAutomaticUpdateLog2(groupNum)
	//更新singer表
	str = "更新表:" + util.GetNowTime()
	LogUtil(str)
	service.UpdateAllTables()
	//更新CacheMap文件
	str = "更新txt文件:" + util.GetNowTime()
	LogUtil(str)
	service.UpdateCacheMap()
	str = "手机歌曲同步结束:" + util.GetNowTime()
	LogUtil(str)
	tt2 := time.Now().Unix()
	str = "共用时" + strconv.FormatInt(tt2-tt1, 10)
	LogUtil(str)

}

func LogUtil(str string) {
	util.RecordLogUtil([]byte(str))
	fmt.Println(str)
}
