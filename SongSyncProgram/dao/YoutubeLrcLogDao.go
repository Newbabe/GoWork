package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func SaveLrcLog(sourceId, lrc1State, lrc2State, duet int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.youtube_lrc_log(source_id,lrc1_state ,lrc2_state,duet ,create_time) values (?,?,?,?,?)"
	i, err := util.URDUtil(sql, sourceId, lrc1State, lrc2State, duet, util.GetNowTime())
	if err != nil {
		panic(err)
		return 0
	}
	return i
}
func UpdateLrc1Log(sourceId, lrc1State int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update song.youtube_lrc_log set lrc1_state=? where source_id=?"
	i, err := util.URDUtil(sql, lrc1State, sourceId)
	if err != nil {
		panic(err)
		return 0
	}
	return i
}
func UpdateLrc2Log(sourceId, lr2State int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update song.youtube_lrc_log set lrc2_state=? where source_id=?"
	i, err := util.URDUtil(sql, lr2State, sourceId)
	if err != nil {
		panic(err)
		return 0
	}
	return i
}

func QueryLrcLogBySourceId(sourceId int) model.LrcLog {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id, source_id, lrc1_state, duet, lrc2_state from song.youtube_lrc_log where source_id=?"
	query := util.GetDbSongRead().QueryRow(sql, sourceId)
	var lrcLog model.LrcLog
	query.Scan(&lrcLog.Id, &lrcLog.SourceId, &lrcLog.Lrc1State, &lrcLog.Duet, &lrcLog.Lrc2State)
	return lrcLog
}
func QueryAllYoutubeLrcLog() []model.LrcLog {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "select id, source_id, lrc1_state, duet, lrc2_state from song.youtube_lrc_log where lrc1_state=1 or lrc2_state=1"
	query, err := util.GetDbSongRead().Query(sql)
	defer query.Close()
	if err != nil {
		panic(err)
		return nil
	}
	var lrcLogList []model.LrcLog
	for query.Next() {
		var lrcLog model.LrcLog
		query.Scan(&lrcLog.Id, &lrcLog.SourceId, &lrcLog.Lrc1State, &lrcLog.Duet, &lrcLog.Lrc2State)
		lrcLogList = append(lrcLogList, lrcLog)
	}
	return lrcLogList
}
