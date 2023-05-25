package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
)

func GetSearchKeyWordTvSongIdBySourceId(sourceId int) int {
	return dao.GetSearchKeyWordTvSongIdBySourceId(sourceId)
}

func UpdateSearKey(searchKeyWordTv model.SearchKeyWordTv) {
	dao.UpdateSearKey(searchKeyWordTv)

}

func SaveSearchKeyWordTv(searchKeyWordTv model.SearchKeyWordTv) {
	dao.SaveSearchKeyWordTv(searchKeyWordTv)
}
