package malysiaDao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func SelectSongInfoMergeDb(tableName string) []model.SongInfoMerger {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "SELECT `source_id`,`song_name`,`singer`,`status` FROM " + tableName + "  where source = 3"
	query, err := util.GetDbMalysia().Query(sql)
	defer query.Close()
	if err != nil {

		return []model.SongInfoMerger{}
	}
	var list []model.SongInfoMerger
	for query.Next() {
		var info model.SongInfoMerger
		query.Scan(&info.SourceId, &info.SongName, &info.Singer, &info.Status)
		list = append(list, info)
	}
	return list
}
func GetSongInfoBySongNameAndSinger(SongName, Singer, tableName string) []map[string]int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select  id,status,source_id from  " + tableName + " where song_name=? and singer=? and source=3"
	query, err := util.GetDbMalysia().Query(sql, SongName, Singer)
	if err != nil {
		panic(err)
		return nil
	}
	defer query.Close()
	ListMap := make([]map[string]int, 0)
	for query.Next() {
		songInfoMap := make(map[string]int)
		var songId, status, sourceId int
		query.Scan(&songId, &status, &sourceId)
		songInfoMap["songId"] = songId
		songInfoMap["status"] = status
		songInfoMap["sourceId"] = sourceId
		ListMap = append(ListMap, songInfoMap)
	}
	return ListMap
}
func UpdateOldSongStatue(SongId int, tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update " + tableName + " set `status`=-1 where id=? "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		panic(err)
		return
	}
	_, err = stmt.Exec(SongId)
	if err != nil {
		return
	}

}
