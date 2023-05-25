package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

//得到调整值
func GetAdjustmentValue(typeInt int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var sql string
	if typeInt == 1 {
		sql = "SELECT count(*) FROM `song_info_youtube_info` where source = 3"
	} else if typeInt == 2 {
		sql = "SELECT count(*) FROM `song_info_youtube_info` where source = 3 and time_reduce<0"
	} else if typeInt == 3 {
		sql = "SELECT count(*) FROM `song_info_youtube_info` where source = 3 and time_reduce>0"
	}

	row := util.GetDbSongRead().QueryRow(sql)
	var count int
	row.Scan(&count)
	return count
}
func AddSongInfoYoutubeInfo(SourceId int, YouTubeId string, vidoStartTimeOffset string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert ignore into song.song_info_youtube_info (lrc_song_id, video_id, time_reduce) values (?,?,?) "
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {

		panic(err)
		return
	}
	_, err = stmt.Exec(SourceId, YouTubeId, vidoStartTimeOffset)
	if err != nil {
		return
	}
}

//清空YoutubeOffest
func ClearYoutubeOffest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DELETE FROM song.song_info_youtube_info where source = 3"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)

		return
	}
	stmt.Exec()

}
