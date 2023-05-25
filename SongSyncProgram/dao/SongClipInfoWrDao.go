package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetSongClipInfoDaoBySongId(songId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id from herook.song_clip_info_wr where song_id = ?"
	row := util.GetDbSongRead().QueryRow(sql, songId)
	var id int
	row.Scan(&id)
	return id

}

func SaveSongClipInfo(SongId int, SongName, clip1, clip2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into herook.song_clip_info_wr (`song_id`, `song_name`, `clip`, `clip2`) values (?,?,?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(SongId, SongName, clip1, clip2)
}
func UpdateSongClipInfo(SongId int, SongName, clip1, clip2 string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update  herook.song_clip_info_wr set `song_name` = ? , `clip` = ?, `clip2` = ?, lrc_version = lrc_version + 1 where song_id = ?"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(SongName, clip1, clip2, SongId)

}
func UpdateSongClipInfo2(SongId int, SongName, clip1 string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update  herook.song_clip_info_wr set `song_name` = ? , `clip` = ?, lrc_version = lrc_version + 1 where song_id = ?"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(SongName, clip1, SongId)

}
func UpdateSongClipInfo3(SongId int, SongName, clip2 string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update  herook.song_clip_info_wr set `song_name` = ? , `clip2` = ?, lrc_version = lrc_version + 1 where song_id = ?"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(SongName, clip2, SongId)

}
