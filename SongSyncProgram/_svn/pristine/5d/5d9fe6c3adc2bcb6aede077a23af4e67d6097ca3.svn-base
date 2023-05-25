package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetYoutubeSongFilterBySongId(songId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id from song.youtube_song_filter where song_id = ? and status = 1 limit 1"
	row := util.GetDbSongRead().QueryRow(sql, songId)
	var id int
	row.Scan(&id)
	return id
}
func SaveYoutubeSongFilter(id int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT INTO song.`youtube_song_filter` (`song_id`) VALUES(?) ;"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(id)
}
