package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
	"time"
)

func UpdateSongInfoMergeTvStatus(tableName string, newStatus, oldStatus int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//	sql := "update song.song_info_merge_tv_test_tmp SET status = ? where source  = 3 and status = ?"
	sql := "update song.song_info_merge_tv_tmp SET `status` = ? where `source`  = ?  and `status` = ? "
	stmt, err := util.GetDbSong().Prepare(sql)
	if err != nil {
		fmt.Println("异常sql:", sql)
		return
	}
	defer stmt.Close()
	stmt.Exec(newStatus, 3, oldStatus)

}
func UpdateSongInfoMergeTvStatus2(newStatus, sourceId int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//	sql := "update song.song_info_merge_tv_test_tmp SET status = ? where source  = 3 and status = ?"
	sql := "update song.song_info_merge_tv_tmp SET `status` = ? where `source`  =? and `source_id`=? "
	//sql := "update song.song_info_merge_tv_tmp SET status = ? where source  = 3  and source_id=? "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {

		panic(err)
		return
	}
	stmt.Exec(newStatus, 3, sourceId)

}

//添加唯一索引
func AddSongInfoMergeTV2UNIQUEIndex(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "ALTER TABLE `song_info_merge_tv2_test` ADD UNIQUE source_id( `source_id`,`source`)"
	sql := "ALTER TABLE `" + tableName + "` ADD UNIQUE source_id( `source_id`,`source`)"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()

}

//添加主键
func AddSongInfoMergeTV2PrimaryKey(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "ALTER TABLE " + tableName + " ADD PRIMARY KEY ( `id` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()

}

//添加索引
func AddSongInfoMergeTV2Index(tableName, ColumnName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "ALTER TABLE " + tableName + " ADD INDEX " + ColumnName + " ( `" + ColumnName + "` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}
func AddSongInfoMergeTV2Index2(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "ALTER TABLE " + tableName + " ADD INDEX source_status ( `source`,`status` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}

//添加全文索引
func AddSongInfoMergeTV2FULLTEXTIndex(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "ALTER TABLE " + tableName + "  ADD FULLTEXT song_name  ( `song_name` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}

func AddSongInfoMergeTV2BySongInfoMerge(songInfoMerge, songInfoMergeTv2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "insert into song_info_merge_tv2_test  select a.* from song_info_merge_test  a INNER JOIN (SELECT max(id) maxId FROM `song_info_merge_tv2_test`) b on a.id > b.maxId"
	sql := "insert into " + songInfoMergeTv2 + "  select a.* from " + songInfoMerge + "  a INNER JOIN (SELECT max(id) maxId FROM `" + songInfoMergeTv2 + "`) b on a.id > b.maxId"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec()
}
func UpdateSongInfoMergeTV2Status(songInfoMergeTv2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "UPDATE song_info_merge_tv2_test set status = 1 where status = -2"
	sql := "UPDATE " + songInfoMergeTv2 + " set status = 1 where status = -2"

	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}
func SelectNewSaveSongInfoMerge(songInfoMerge, songInfoMergeTv2 string) []model.NewSaveSongInfoMerge {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select a.id, a.song_name, a.singer, a.source_id, a.source from song." + songInfoMerge + "  a INNER JOIN (SELECT max(id) maxId FROM `" + songInfoMergeTv2 + "`) b on a.id > b.maxId "
	rows, err := util.GetDbSongRead().Query(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println("异常sql:", err)
		return nil
	}
	var NewSaveSongInfoMergeList []model.NewSaveSongInfoMerge
	for rows.Next() {
		var NewSaveSongInfoMerge model.NewSaveSongInfoMerge
		rows.Scan(&NewSaveSongInfoMerge.Id, &NewSaveSongInfoMerge.SongName, &NewSaveSongInfoMerge.Singer, &NewSaveSongInfoMerge.SourceId, &NewSaveSongInfoMerge.Source)
		NewSaveSongInfoMergeList = append(NewSaveSongInfoMergeList, NewSaveSongInfoMerge)
	}
	return NewSaveSongInfoMergeList
}

func GetSongLangById(songId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select lang from song.song_info_merge_tv where id =? "
	row := util.GetDbSongRead().QueryRow(sql, songId)
	var lang int
	row.Scan(&lang)
	return lang
}
func GetSongPitchInfoBySourceId(sourceId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id from song.song_pitch_info where source_id =?"
	stmt := util.GetDbSongRead().QueryRow(sql, sourceId)
	var id int
	stmt.Scan(&id)
	return id

}
func GetLasUpdateTimeTvBySourceId(sourceId int) time.Time {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select last_update_time from song.last_update_time_tv  WHERE source_id=?"
	stmt := util.GetDbSongRead().QueryRow(sql, sourceId)
	var LsatUpdateTime time.Time
	stmt.Scan(&LsatUpdateTime)
	return LsatUpdateTime
}
func AddLasUpdateTimeTV(sourceId int, lastUpdateTime string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "insert into song.last_update_time_tv(source_id,last_update_time) values (?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec(sourceId, lastUpdateTime)

}
func UpdateLasUpdateTimeTV(sourceId int, lastUpdateTime string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update  song.last_update_time_tv set last_update_time=? where source_id=?"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec(lastUpdateTime, sourceId)
}

func SaveSongPitchInfo(SourceId, minPitch, maxPitch, SongPitchInfoNum int, minMidi, maxMidi string) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT INTO song.`song_pitch_info` (`source_id`,min_pitch, max_pitch, min_midi, max_midi ) VALUES(?,?,?,?,?) ;"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(sql)
		return SongPitchInfoNum
	}
	result, err := stmt.Exec(SourceId, minPitch, maxPitch, minMidi, maxMidi)
	if err != nil {
		return SongPitchInfoNum
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return SongPitchInfoNum
	}

	i := SongPitchInfoNum + int(affected)
	return i
}

func SelectSongInfoMergeTV() []model.SongInfoMerger {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "SELECT `source_id`,`song_name`,`singer`,`status`, song_size, lrc_size, lrc_size2, update_time FROM song.song_info_merge_tv  where source =3  "
	rows, err := util.GetDbSongRead().Query(sql)
	if err != nil {
		//fmt.Println("出错的sql:", sql)
		panic(err)
		return nil
	}
	var songInfoMergerList []model.SongInfoMerger
	for rows.Next() {
		var songInfoMerger model.SongInfoMerger
		rows.Scan(&songInfoMerger.SourceId, &songInfoMerger.SongName, &songInfoMerger.Singer, &songInfoMerger.Status, &songInfoMerger.SongSize, &songInfoMerger.LrcSize, &songInfoMerger.LrcSize2, &songInfoMerger.UpdateTime)
		songInfoMergerList = append(songInfoMergerList, songInfoMerger)
	}
	return songInfoMergerList

}
func AddSongInfoMergeTvTmp(songInfoMerge string, tmp model.SongInfoMergeTvTmp, num int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	//sql := "INSERT ignore  INTO`song`.`song_info_merge_test`( " +
	sql := "INSERT ignore  INTO`song`.`" + songInfoMerge + "`( " +
		"source_id, source, word_part, song_name, song_name_phonetic, " +
		"singer, singer_phonetic, album, line_writer, song_writer, " +
		"gender, lang, year, dossier_name, is_duet, " +
		"song_size, lrc_size, background_image_size, icon_size, song_time, " +
		"status, lrc_type, update_time, intonation, create_date, " +
		"mv_size, song_path, song_name_stroke_num, singer_stroke_num," +
		" song_name_simple, singer_simple, song_version, lrc_version, lrc_channel, " +
		"lrc_size2, song_chinese_phonetic, singer_chinese_phonetic, search_word, youtube_song_type) " +
		"values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return num
	}
	stmt.Exec(tmp.SourceId, tmp.Source, tmp.WordPart, tmp.SongName, tmp.SongNamePhonetic,
		tmp.Singer, tmp.SingerNamePhonetic, tmp.Album, tmp.LineWriter, tmp.SongWrite,
		tmp.Gender, tmp.Lang, tmp.Year, tmp.DossierName, tmp.IsDuet,
		tmp.SongSize, tmp.LrcSize, tmp.BackgroundImageSize, tmp.IconSize, tmp.SongTime,
		tmp.Status, tmp.LrcType, tmp.UpdateTime, tmp.Intonation, tmp.CreateDate,
		tmp.MvSize, tmp.SongPath, tmp.SongNameStrokeNum, tmp.SingerStrokeNum,
		tmp.SongNameSimple, tmp.SingerSimple, tmp.SongVersion, tmp.LrcVersion, tmp.LrcChannel,
		tmp.LrcSize2, tmp.SongChinesePhonetic, tmp.SingerChinesePhonetic, tmp.SearchWord, tmp.YouTubeSongType)

	return num + 1
}
func UpdateSongInfoMergeTvTmp(songInfoMergeTvTmp string, tmp model.SongInfoMergeTvTmp, num int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "update song.song_info_merge_tv_test_tmp set word_part=?,song_name=? ," +
	sql := "update song." + songInfoMergeTvTmp + " set word_part=?,song_name=? ," +
		"singer=?,album=?,line_writer=?,song_writer=?," +
		"gender=?,lang=?,year=?,dossier_name=?,is_duet=?," +
		"song_size=?,lrc_size=?,background_image_size=?,icon_size=?," +
		"song_time=?,status=?,lrc_type=?,update_time=?,intonation=?," +
		"mv_size=?,song_path=?,song_name_stroke_num=?,singer_stroke_num=?,song_name_simple=?," +
		"singer_simple=?,lrc_channel=?,lrc_size2=?,song_chinese_phonetic=?,singer_chinese_phonetic=?," +
		"search_word=?,youtube_song_type=? where source=3 and source_id=?"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return num
	}

	_, err = stmt.Exec(tmp.WordPart, tmp.SongName, tmp.Singer, tmp.Album, tmp.LineWriter, tmp.SongWrite,
		tmp.Gender, tmp.Lang, tmp.Year, tmp.DossierName, tmp.IsDuet,
		tmp.SongSize, tmp.LrcSize, tmp.BackgroundImageSize, tmp.IconSize,
		tmp.SongTime, tmp.Status, tmp.LrcType, tmp.UpdateTime, tmp.Intonation,
		tmp.MvSize, tmp.SongPath, tmp.SongNameStrokeNum, tmp.SingerStrokeNum, tmp.SongNameSimple,
		tmp.SingerSimple, tmp.LrcChannel, tmp.LrcSize2, tmp.SongChinesePhonetic, tmp.SingerChinesePhonetic,
		tmp.SearchWord, tmp.YouTubeSongType, tmp.SourceId)
	if err != nil {
		return 0
	}
	return num + 1
}

func Update(songInfoMergeTv string, sourceId, lrcSize1, lrcSize2, songSize, songVersion int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//	sql := "UPDATE song.song_info_merge_tv_test SET lrc_size = ? ,lrc_size2=?, song_size=?, song_version = ? WHERE source_id = ? AND source = 3"
	sql := "UPDATE song." + songInfoMergeTv + " SET lrc_size = ? ,lrc_size2=?, song_size=?, song_version = ? WHERE source_id = ? AND source = 3"
	//fmt.Println("sql==", sql)
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt.Exec(lrcSize1, lrcSize2, songSize, songVersion, sourceId)

}
func UpdatePhone(songInfoMerge string, sourceId, lrcSize1, lrcSize2, songSize, songVersion int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "UPDATE song.song_info_merge_test SET lrc_size = ? ,lrc_size2=?, song_size=?, song_version = ? WHERE source_id = ? AND source = 3"
	sql := "UPDATE song." + songInfoMerge + " SET lrc_size = ? ,lrc_size2=?, song_size=?, song_version = ? WHERE source_id = ? AND source = 3"

	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec(lrcSize1, lrcSize2, songSize, songVersion, sourceId)

}
func UpdateFinallyTv(songInfoMergeTvTmp string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//	sql := "delete from song.song_info_merge_tv_test_tmp where status = -99"
	sql := "delete from song." + songInfoMergeTvTmp + " where status = -99"
	stmt, err := util.GetDbSong().Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	stmt.Exec()
}

func GetSongInfoMergeBySongNameAndSinger(songInfoMerge string, songName, singer string) []int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	//sql := "SELECT `id` from song.song_info_merge_test WHERE song_name like concat(?,'%') and singer = ?  and youtube_song_type = 3 and ((source =4  and song_version = '1' ) or (source =3  and song_version = '2' ) ) "
	sql := "SELECT `id` from song." + songInfoMerge + " WHERE song_name like concat(?,'%') and singer = ?  and youtube_song_type = 3 and ((source =4  and song_version = '1' ) or (source =3  and song_version = '2' ) ) "

	rows, err := util.GetDbSongRead().Query(sql, songName, singer)
	defer rows.Close()
	if err != nil {
		return nil
	}
	var songIdList []int
	for rows.Next() {
		var songId int
		rows.Scan(&songId)
		songIdList = append(songIdList, songId)
	}
	return songIdList

}

func GetSongInfoMergeTVInfoBySongNameAndSinger(SongName, Singer string) []int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "SELECT `id` from song.song_info_merge WHERE song_name like concat(?,'%') and singer = ? and youtube_song_type != 2 and status in (1,2,-2) and source = 3 "
	rows, err := util.GetDbSongRead().Query(sql, SongName, Singer)
	defer rows.Close()
	if err != nil {
		fmt.Println("关闭流报错", err)
		return nil
	}
	var idList []int
	for rows.Next() {
		var id int
		rows.Scan(&id)
		idList = append(idList, id)

	}
	return idList
}

func GetSongInfoBySongNameAndSinger(SongName, Singer, tableName string) []map[string]int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select  id,status,source_id from song." + tableName + " where song_name=? and singer=? and source=3"
	query, err := util.GetDbSongRead().Query(sql, SongName, Singer)
	if err != nil {
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
	sql := "update song." + tableName + " set `status`=-1 where id=? "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	_, err = stmt.Exec(SongId)
	if err != nil {
		return
	}

}
