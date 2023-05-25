package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
)

func GetDoremiLastUpdateTime() []model.DoremiUpdateTime {
	return dao.GetDoremiUpdateTime()
}
func UpdateDoremiLastUpdateTime(sourceId int) {
	updateTime := GetLasUpdateTimeBySourceId(sourceId).String()
	dao.UpdatelastUpdate(sourceId, updateTime)

}
