package malysiaDao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetAdjustmentValue(typeInt int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var sql string
	if typeInt == 1 {
		sql = "SELECT count(*) FROM song.`song_info_youtube_info` where source = 3"
	} else if typeInt == 2 {
		sql = "SELECT count(*) FROM song.`song_info_youtube_info` where source = 3 and time_reduce<0"
	} else if typeInt == 3 {
		sql = "SELECT count(*) FROM song.`song_info_youtube_info` where source = 3 and time_reduce>0"
	}
	row := util.GetDbMalysia().QueryRow(sql)
	var num int
	row.Scan(&num)
	return num
}

func AddSongInfoYoutubeInfo(SourceId int, tableName, YouTubeId, vidoStartTimeOffset string) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert ignore into " + tableName + " (lrc_song_id, video_id, time_reduce) values (?,?,?) "
	urdUtil, err := util.MalysiaURDUtil(sql, SourceId, YouTubeId, vidoStartTimeOffset)
	if err != nil {
		panic(err)
		return 0
	}
	return urdUtil
}

//清除数据
func ClearYoutubeOffest(tableNaem string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DELETE FROM " + tableNaem + " where source = 3"
	stmt, err := util.GetDbMalysia().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)

		return
	}
	stmt.Exec()

}
