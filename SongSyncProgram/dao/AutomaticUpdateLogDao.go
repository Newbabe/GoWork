package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

//获取组数据
func GetGroupNum() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select group_num from song.automatic_update_log order by group_num DESC LIMIT 0 , 1"
	row := util.GetDbSongRead().QueryRow(sql)
	var num int
	row.Scan(&num)
	return num
}

/**
 * 保存log
 * @param createTime
 * @param status
 * @param offset
 * @param groupNum
 */
func SaveAutomaticUpdateLog(createTime string, status int, offset string, groupNum int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.automatic_update_log (create_time,status,offset,group_num) values (?,?,?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("SaveAutomaticUpdateLog占位符赋值异常:", sql)
		panic(err)
		return
	}
	_, err = stmt.Exec(createTime, status, offset, groupNum)
	if err != nil {

		panic(err)
		return
	}

}
