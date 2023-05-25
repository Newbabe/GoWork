package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/util"
	"strconv"
)

func GetGroupNum() int {
	return dao.GetGroupNum()
}

//添加日志
func SaveAutomaticUpdateLog(groupNum int) {
	createTime := util.GetNowTime()
	youtubeInfoCount1 := GetAdjustmentValue(1)
	youtubeInfoCount2 := GetAdjustmentValue(2)
	youtubeInfoCount3 := GetAdjustmentValue(3)
	dao.SaveAutomaticUpdateLog(createTime, 1, strconv.Itoa(youtubeInfoCount1), groupNum)
	dao.SaveAutomaticUpdateLog(createTime, 2, strconv.Itoa(youtubeInfoCount2), groupNum)
	dao.SaveAutomaticUpdateLog(createTime, 3, strconv.Itoa(youtubeInfoCount3), groupNum)
}
func SaveAutomaticUpdateLog2(groupNum int) {
	createTime := util.GetNowTime()
	youtubeInfoCount1 := GetAdjustmentValue(1)
	youtubeInfoCount2 := GetAdjustmentValue(2)
	youtubeInfoCount3 := GetAdjustmentValue(3)
	dao.SaveAutomaticUpdateLog(createTime, 4, strconv.Itoa(youtubeInfoCount1), groupNum)
	dao.SaveAutomaticUpdateLog(createTime, 5, strconv.Itoa(youtubeInfoCount2), groupNum)
	dao.SaveAutomaticUpdateLog(createTime, 6, strconv.Itoa(youtubeInfoCount3), groupNum)
}
