package malysiaDao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
	"time"
)

func GetLastUpdateTimeBySourceId(sourceId int) time.Time {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select last_update_time from song.last_update_time_malysia WHERE source_id=?"
	stmt := util.GetDbMalysia().QueryRow(sql, sourceId)
	var LsatUpdateTime time.Time
	stmt.Scan(&LsatUpdateTime)
	return LsatUpdateTime
}

func AddLasUpdateTime(sourceId int, lastUpdateTime string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.last_update_time_malysia(source_id,last_update_time) values (?,?)"
	stmt, err := util.GetDbMalysia().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec(sourceId, lastUpdateTime)

}
func UpdateLasUpdateTime(sourceId int, lastUpdateTime string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update  song.last_update_time_malysia set last_update_time=? where source_id=?"
	stmt, err := util.GetDbMalysia().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec(lastUpdateTime, sourceId)
}
