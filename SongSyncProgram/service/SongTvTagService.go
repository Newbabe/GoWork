package service

import (
	"SongSyncProgram/dao"
)

func GetSongTvTagId(songId int, tag int) int {
	return dao.GetSongTvTagId(songId, tag)
}

func GetSongTvTagMaxRank(tag int) int {
	return dao.GetSongTvTagMaxRank(tag)

}
func SaveSongTvTag(songId, sourceId, source, tagId, ranking int) {
	dao.SaveSongTvTag(songId, sourceId, source, tagId, ranking)
}
func GetSongTvTagIdBySongId(songId int) int {
	return dao.GetSongTvTagIdBySongId(songId)
}
