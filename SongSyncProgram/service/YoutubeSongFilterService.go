package service

import "SongSyncProgram/dao"

func GetYoutubeSongFilterBySongId(songId int) int {
	return dao.GetYoutubeSongFilterBySongId(songId)
}

func SaveYoutubeSongFilter(id int) {
	dao.SaveYoutubeSongFilter(id)
}
