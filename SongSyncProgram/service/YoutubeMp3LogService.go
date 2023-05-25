package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
)

func SaveMp3Log(sourceId int) {
	dao.SaveMp3Log(sourceId)

}

func UpdateMp3Log(sourceId int) {
	dao.UpdateMp3Log(sourceId)
}
func QueryAllYoutubeMp3Log() []int {
	return dao.QueryAllYoutubeMp3Log()
}

func GetMp3LogBySourceId(sourceId int) model.Mp3Log {
	return dao.GetMp3LogBSourceId(sourceId)
}

func UpdateMp3LogState(sourceId int) int {
	return dao.UpdateMp3LogState(sourceId)
}

func QueryAllYoutubeMp3LogDownloadFailed() []int {
	return dao.QueryAllYoutubeMp3LogDownloadFailed()
}
