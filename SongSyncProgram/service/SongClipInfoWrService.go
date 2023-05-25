package service

import "SongSyncProgram/dao"

func GetSongClipInfoDaoBySongId(songId int) int {
	return dao.GetSongClipInfoDaoBySongId(songId)
}
func SaveSongClipInfo(songId int, SongName, clip1, clip2 string) {
	dao.SaveSongClipInfo(songId, SongName, clip1, clip2)
}
func UpdateSongClipInfo(songId int, SongName, clip1, clip2 string) {
	if clip1 != "" && clip2 != "" {
		dao.UpdateSongClipInfo(songId, SongName, clip1, clip2)
	} else if clip1 != "" {
		dao.UpdateSongClipInfo2(songId, SongName, clip1)
	} else {
		dao.UpdateSongClipInfo3(songId, SongName, clip2)
	}

}
