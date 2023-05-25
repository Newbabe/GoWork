package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func SaveMp3Log(sourceId int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.youtube_mp3_log(source_id,state,create_time ) values (?,?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(sourceId, 1, util.GetNowTime())

}
func UpdateMp3Log(sourceId int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update song.youtube_mp3_log set state=0 where source_id=?"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(sourceId)
}
func QueryAllYoutubeMp3Log() []int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select source_id from song.youtube_mp3_log where state=1"
	stmt, err := util.GetDbSongRead().Query(sql)
	defer stmt.Close()
	if err != nil {
		return []int{}
	}
	var sourceIdList []int
	for stmt.Next() {
		var sourceId int
		stmt.Scan(&sourceId)
		sourceIdList = append(sourceIdList, sourceId)
	}
	return sourceIdList
}
func QueryAllYoutubeMp3LogDownloadFailed() []int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select source_id from song.youtube_mp3_log where state=-1"
	stmt, err := util.GetDbSongRead().Query(sql)
	defer stmt.Close()
	if err != nil {
		return []int{}
	}
	var sourceIdList []int
	for stmt.Next() {
		var sourceId int
		stmt.Scan(&sourceId)
		sourceIdList = append(sourceIdList, sourceId)
	}
	return sourceIdList
}
func GetMp3LogBSourceId(sourceId int) model.Mp3Log {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id, source_id, state from song.youtube_mp3_log where source_id=? "
	query := util.GetDbSongTest().QueryRow(sql, sourceId)
	var mp3Log model.Mp3Log
	query.Scan(&mp3Log.Id, &mp3Log.SourceId, &mp3Log.State)
	return mp3Log

}
func UpdateMp3LogState(sourceId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update song.youtube_mp3_log set state=-1 where state=1 and source_id=?"
	i, err := util.URDUtil(sql, sourceId)
	if err != nil {
		panic(err)
		return 0
	}
	return i
}
