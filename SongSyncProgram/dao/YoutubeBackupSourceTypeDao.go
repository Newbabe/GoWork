package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func Save(SongId, sourceId, youtubeSongtype, songVersion, lrcVersion int, lrcChannel string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert ignore into song.youtube_backup_source_type (song_id,source_id,source,lrc_channel,song_version,lrc_version,youtube_song_type,update_time,create_time) values (?,?,?,?,?,?,?,?,?)"
	stmt, err := util.DbSongTest.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	result, err := stmt.Exec(SongId, sourceId, 3, lrcChannel, songVersion, lrcVersion, youtubeSongtype, util.GetNowTime(), util.GetNowTime())

	if err != nil {
		return
	}
	i, err := result.RowsAffected()
	if err != nil {
		return
	}
	if i <= 0 {
		sql = "update song.youtube_backup_source_type set lrc_channel = ?, song_version = song_version+1 ,update_time = ? where source_id = ?"
		stmt2, _ := util.DbSongTest.Prepare(sql)
		defer stmt2.Close()
		stmt2.Exec(lrcChannel, songVersion, util.GetNowTime(), SongId)
	}

}
