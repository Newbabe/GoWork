package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func SaveUploadLrcFailRecord(sourceId int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.upload_lrc_fail_record (`source_id`) values(?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(sourceId)
}
