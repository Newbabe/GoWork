package malysiaDao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetGroupNum() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select group_num from song.automatic_update_log_tv order by group_num DESC LIMIT 0 , 1"
	row := util.GetDbMalysia().QueryRow(sql)
	var num int
	row.Scan(&num)
	return num

}
func SaveAutomaticUpdateLog(tableName, createTime string, status, offset, groupNUm int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into " + tableName + " (create_time,status,offset,group_num) values(?,?,?,?)"
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		return
	}
	prepare.Exec(createTime, status, offset, groupNUm)

}
