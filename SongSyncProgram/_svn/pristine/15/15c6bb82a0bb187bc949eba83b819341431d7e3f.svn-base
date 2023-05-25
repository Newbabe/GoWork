package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
)

func SaveLrcLog(sourceId, lrc1State, lrc2State, duet int) int {
	return dao.SaveLrcLog(sourceId, lrc1State, lrc2State, duet)

}

func UpdateLrc1Log(sourceId, lrc1State int) int {
	return dao.UpdateLrc1Log(sourceId, lrc1State)
}
func UpdateLrc2Log(sourceId, lrc2State int) int {
	return dao.UpdateLrc2Log(sourceId, lrc2State)
}

func QueryLrcLogBySourceId(sourceId int) model.LrcLog {
	return dao.QueryLrcLogBySourceId(sourceId)
}
func QueryAllYoutubeLrcLog() []model.LrcLog {
	return dao.QueryAllYoutubeLrcLog()
}
