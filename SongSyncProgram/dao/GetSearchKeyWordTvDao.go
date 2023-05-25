package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func GetSearchKeyWordTvSongIdBySourceId(sourceId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id from song.search_key_word_tv  where source_id = ? "
	stmt := util.GetDbSongRead().QueryRow(sql, sourceId)
	var SourceId int
	stmt.Scan(&SourceId)

	return SourceId

}

func UpdateSearKey(searchKeyWordTv model.SearchKeyWordTv) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := " update   song.search_key_word_tv set song_name = ? , singer = ? ,song_name_keyword = ? ,singer_keyword = ? ,album_keyword = ?   where source_id = ?"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(searchKeyWordTv.SongName, searchKeyWordTv.Singer, searchKeyWordTv.SongNameKeyWord, searchKeyWordTv.SingerKeyWord, searchKeyWordTv.AlbumKeyWord, searchKeyWordTv.SourceId)

}
func SaveSearchKeyWordTv(searchKeyWordTv model.SearchKeyWordTv) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.search_key_word_tv(source_id,song_name,singer,song_name_keyword,singer_keyword,album_keyword) values(?,?,?,?,?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(searchKeyWordTv.SourceId, searchKeyWordTv.SongName, searchKeyWordTv.Singer, searchKeyWordTv.SongNameKeyWord, searchKeyWordTv.SingerKeyWord, searchKeyWordTv.AlbumKeyWord)
}
