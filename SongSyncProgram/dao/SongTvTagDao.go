package dao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetSongTvTagId(songId int, tagId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id from song.song_tv_tag where song_id =? and tag =? "
	row := util.GetDbSongRead().QueryRow(sql, songId, tagId)
	var id int
	row.Scan(&id)
	return id

}
func SaveSongTvTag(songId, sourceId, source, tagId, ranking int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into  song.song_tv_tag(song_id, source_id, source, tag, ranking) values(?,?,?,?,?)"
	stmt, err := util.DbSongTest.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(songId, sourceId, source, tagId, ranking)

}

func GetSongTvTagMaxRank(tag int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select max(ranking) from song.song_tv_tag where tag = ?"
	row := util.GetDbSongRead().QueryRow(sql, tag)
	var maxRanking int
	row.Scan(&maxRanking)
	return maxRanking
}
func GetSongTvTagIdBySongId(songId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select tag from song.song_tv_tag where song_id =?"
	row := util.GetDbSongRead().QueryRow(sql, songId)
	var tag int
	row.Scan(&tag)
	return tag

}
