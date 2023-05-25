package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
	"time"
)

// CreateBackupTable 创建备份表 参数 TableName：主表  BackupTableName:备份表
func CreateBackupTable(TableName, BackupTableName string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "CREATE TABLE " + BackupTableName + " SELECT * FROM  " + TableName
	//fmt.Println(sql)
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("出错的sql:=", err)

		return
	}

}
func CreateBackupTableTestDb(TableName, BackupTableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "CREATE TABLE " + BackupTableName + " SELECT * FROM  " + TableName
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("出错的sql:=", sql, err)
		panic(err)
		return
	}

}
func RemoveTmpTableData(TmpTableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "truncate " + TmpTableName
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("出错的sql:=", sql)
		panic(err)
		return
	}
}
func RemoveTmpTableDataTestDb(TmpTableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "truncate " + TmpTableName
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("出错的sql:=", sql)
		panic(err)
		return
	}
}
func TransferDataBToA(tableNameA, tableNameB string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT INTO " + tableNameA + " SELECT * FROM " + tableNameB
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("出错的sql:=", sql)
		return
	}
}
func TransferDataBToATestDb(tableNameA, tableNameB string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT INTO " + tableNameA + " SELECT * FROM " + tableNameB
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
		return
	}
}

func UpdateYoutubeStatus(songInfoMergeTmp string, oleStatus, newStatus int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "update song.song_info_merge_tmp_test SET status =? where source  = 3 and status = ? "
	sql := "update song.song_info_merge_tmp SET status =? where source  = ? and status = ? "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	r, err := stmt.Exec(newStatus, 3, oleStatus)
	if err != nil {
		fmt.Println("修改youtube状态异常:", sql)
		panic(err)
		return 0
	}
	affected, err := r.RowsAffected()
	return int(affected)
}
func UpdateYoutubeStatus2(songInfoMergeTmp string, newStatus, sourceId int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "update song.song_info_merge_tmp_test SET status =? where source  = 3 and status = ? "
	sql := "update song.song_info_merge_tmp SET status =? where source  = ?  and source_id=?"
	//	sql := "update song.song_info_merge SET `status` =? where `source`  = 3 and status = ? and source_id=?"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	r, err := stmt.Exec(newStatus, 3, sourceId)
	if err != nil {
		fmt.Println("修改youtube状态异常:", sql)
		panic(err)
		return 0
	}
	affected, err := r.RowsAffected()
	return int(affected)
}
func UpdateYoutubeStatusTestDb(songInfoMergeTmp string, oleStatus, newStatus int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update song." + songInfoMergeTmp + " SET status =? where source  = 3 and status = ? "
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	r, err := stmt.Exec(newStatus, oleStatus)
	if err != nil {
		fmt.Println("修改youtube状态异常:", sql)
		panic(err)
		return 0
	}
	affected, err := r.RowsAffected()
	return int(affected)
}
func SelectSongInfoMerge(songInfoMerge string) []model.SongInfoMerger {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "SELECT `source_id`,`song_name`,`singer`,`status`, song_size, lrc_size, lrc_size2, update_time FROM song.song_info_merge_test  where source =3  "
	sql := "SELECT `source_id`,`song_name`,`singer`,`status`, song_size, lrc_size, lrc_size2, update_time FROM song." + songInfoMerge + "  where source =3  "
	rows, err := util.GetDbSongRead().Query(sql)
	if err != nil {
		fmt.Println("出错的sql:", sql)
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

// 添加唯一索引
func AddSongInfoMerge2UNIQUEIndex(songInfoMerge2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "ALTER TABLE `song_info_merge2_test` ADD UNIQUE source_id( `source_id`,`source`)"
	sql := "ALTER TABLE `" + songInfoMerge2 + "` ADD UNIQUE source_id( `source_id`,`source`)"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()

}

// 添加主键
func AddSongInfoMerge2PrimaryKey(songInfoMerge2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "ALTER TABLE `song_info_merge2_test` MODIFY COLUMN `id` int(11) NOT NULL AUTO_INCREMENT FIRST ,ADD PRIMARY KEY (`id`) "
	sql := "ALTER TABLE `" + songInfoMerge2 + "` MODIFY COLUMN `id` int(11) NOT NULL AUTO_INCREMENT FIRST ,ADD PRIMARY KEY (`id`) "
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()

}

// 添加索引
func AddSongInfoMerge2Index(songInfoMerge2, ColumnName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "ALTER TABLE `song_info_merge2_test` ADD INDEX " + ColumnName + " ( `" + ColumnName + "` )"
	sql := "ALTER TABLE `" + songInfoMerge2 + "` ADD INDEX " + ColumnName + " ( `" + ColumnName + "` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}
func AddSongInfoMerge2Index2(songInfoMerge2, ColumnName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "ALTER TABLE `song_info_merge2_test` ADD INDEX " + ColumnName + " ( `source`,`status` )"
	sql := "ALTER TABLE `" + songInfoMerge2 + "` ADD INDEX " + ColumnName + " ( `source`,`status` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}

// 添加全文索引
func AddSongInfoMerge2FULLTEXTIndex(songInfoMerge2, ColumnName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "ALTER TABLE `" + songInfoMerge2 + "` ADD FULLTEXT " + ColumnName + " ( `song_name` )"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	stmt.Exec()
}

func UpdateSongInfoMerge2Data(songInfoMerge2, songInfoMergeTmp string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "insert into song_info_merge2_test select a.* from song_info_merge_tmp_test a INNER JOIN (SELECT max(id) maxId FROM `song_info_merge2_test`) b on a.id > b.maxId "
	sql := "insert into " + songInfoMerge2 + " select a.* from " + songInfoMergeTmp + " a INNER JOIN (SELECT max(id) maxId FROM `" + songInfoMerge2 + "`) b on a.id > b.maxId "
	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}

func UpdateSongInfoMerge2Status(songInfoMerge2 string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "UPDATE song_info_merge2_test set status = 1 where status = -2"
	sql := "UPDATE " + songInfoMerge2 + " set status = 1 where status = -2"
	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", err)
		return
	}
	prepare.Exec()
}
func DelTable(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DROP TABLE " + tableName
	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()

}
func DelTableTestDb(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DROP TABLE " + tableName
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()

}

// UpdateTableName 修改表名 newTableName  oldTableName
func UpdateTableName(newTableName, oldTableName string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "RENAME TABLE " + oldTableName + " TO  " + newTableName
	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}
func UpdateTableNameTestDb(newTableName, oldTableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "RENAME TABLE " + oldTableName + " TO  " + newTableName
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}
func IfExistsDelTable(tableName string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DROP TABLE  if EXISTS  " + tableName
	//fmt.Println("删除备份表的sql", sql)
	prepare, err := util.GetDbSong().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错的sql=", sql)
		return
	}
	prepare.Exec()

}
func IfExistsDelTableTestDb(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DROP TABLE  if EXISTS  " + tableName
	//fmt.Println("删除备份表的sql", sql)
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错的sql=", sql)
		return
	}
	prepare.Exec()

}

func DeleteIdByTableName(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "alter table " + tableName + " drop id;"
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}

func AddIdByTableName(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "alter table " + tableName + " add id int not null primary key auto_increment first;"
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()

}
func CreateTable() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "CREATE TABLE `song_info_youtube_info_tmp`  ( " +
		" `id` int(11) NOT NULL AUTO_INCREMENT,  " +
		"`lrc_song_id` int(11) NOT NULL,  " +
		"`video_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',  " +
		"`time_reduce` int(11) NOT NULL DEFAULT 0," +
		"  `source` int(11) NOT NULL DEFAULT 3," +
		"  PRIMARY KEY (`id`) USING BTREE" +
		") ENGINE = InnoDB AUTO_INCREMENT = 31386 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;"
	prepare, err := util.GetDbSongTest().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		return
	}
	prepare.Exec()
}
func AddSongInfoMergeTmp(songInfoMergeTmp string, tmp model.SongInfoMergeTmp, num int) int {
	//fmt.Println("参数", tmp)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	//sql := "insert ignore into song.song_info_merge_tmp_test" +
	sql := "insert ignore into song." + songInfoMergeTmp +
		" (source_id, source, word_part, song_name, song_name_phonetic, " +
		" singer, singer_phonetic,album, line_writer, " +
		" song_writer, gender, lang, year, dossier_name,is_duet," +
		" song_size, lrc_size, background_image_size, icon_size, " +
		" song_time, status, lrc_type, update_time, intonation, create_date, mv_size," +
		" song_path, song_name_stroke_num, singer_stroke_num, song_name_simple,singer_simple," +
		" song_version, lrc_version, lrc_channel, lrc_size2, song_chinese_phonetic, singer_chinese_phonetic, search_word, youtube_song_type) " +
		" values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)

		panic(err)
		return num
	}
	exec, err := stmt.Exec(tmp.SourceId, tmp.Source, tmp.WordPart, tmp.SongName, tmp.SongNamePhonetic,
		tmp.Singer, tmp.SingerNamePhonetic, tmp.Album, tmp.LineWriter,
		tmp.SongWrite, tmp.Gender, tmp.Lang, tmp.Year, tmp.DossierName, tmp.IsDuet,
		tmp.SongSize, tmp.LrcSize, tmp.BackgroundImageSize, tmp.IconSize,
		tmp.SongTime, tmp.Status, tmp.LrcType, tmp.UpdateTime, tmp.Intonation, tmp.CreateDate, tmp.MvSize,
		tmp.SongPath, tmp.SongNameStrokeNum, tmp.SingerStrokeNum, tmp.SongNameSimple, tmp.SingerSimple,
		tmp.SongVersion, tmp.LrcVersion, tmp.LrcChannel, tmp.LrcSize2, tmp.SongChinesePhonetic, tmp.SingerChinesePhonetic, tmp.SearchWord, tmp.YouTubeSongType)
	if err != nil {
		return num
	}
	_, err = exec.RowsAffected()
	if err != nil {
		return num
	}
	//fmt.Println("添加受影响的行数", affected)
	num += 1
	return num
}

func UpdateSongInfoMergeTmp(songInfoMergeTmp string, tmp model.SongInfoMergeTmp, num int) int {
	//fmt.Println("canshu ", tmp)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	//	sql := "update song.song_info_merge_tmp_test  SET " +
	sql := "update song." + songInfoMergeTmp + "  SET " +
		"  `word_part` = ?, `song_name` = ?, `singer` = ?,  `album` = ?,  `line_writer` = ?," +
		"  `song_writer` = ?,  `gender` = ?,  `lang` = ?,  `year` = ?,  `dossier_name` = ?," +
		"  `is_duet` = ?,  `song_size` = ?,  `lrc_size` = ?,  `background_image_size` = ?,  `icon_size` = ?," +
		"  `song_time` = ?,  `status` = ?,  `lrc_type` = ?, `update_time` = ?,  `intonation` = ?," +
		"  `mv_size` = ?,  `song_path` = ?,  `song_name_stroke_num` = ?, `singer_stroke_num` = ?,  `song_name_simple` = ?," +
		"  `singer_simple` = ?, lrc_channel = ?, lrc_size2 = ?,song_chinese_phonetic = ?,singer_chinese_phonetic = ?," +
		" search_word=?,youtube_song_type = ? where source = 3 and `source_id` = ? "

	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		panic(err)
		return num
	}
	exec, err := stmt.Exec(tmp.WordPart, tmp.SongName, tmp.Singer, tmp.Album, tmp.LineWriter,
		tmp.SongWrite, tmp.Gender, tmp.Lang, tmp.Year, tmp.DossierName,
		tmp.IsDuet, tmp.SongSize, tmp.LrcSize, tmp.BackgroundImageSize, tmp.IconSize,
		tmp.SongTime, tmp.Status, tmp.LrcType, tmp.UpdateTime, tmp.Intonation,
		tmp.MvSize, tmp.SongPath, tmp.SongNameStrokeNum, tmp.SingerStrokeNum, tmp.SongNameSimple,
		tmp.SingerSimple, tmp.LrcChannel, tmp.LrcSize2, tmp.SongChinesePhonetic, tmp.SingerChinesePhonetic,
		tmp.SearchWord, tmp.YouTubeSongType, tmp.SourceId)

	_, err = exec.RowsAffected()
	if err != nil {
		panic(err)
		return num
	}
	//fmt.Println("修改受影响的行数", affected)
	num += 1
	return num
}

func GetLasUpdateTimeBySourceId(sourceId int) time.Time {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select last_update_time from song.last_update_time  WHERE source_id=?"
	stmt := util.GetDbSongRead().QueryRow(sql, sourceId)
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
	sql := "insert into song.last_update_time(source_id,last_update_time) values (?,?)"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		panic(err)
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
	sql := "update  song.last_update_time set last_update_time=? where source_id=?"
	stmt, err := util.GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("出错sql", sql)
		panic(err)
		return
	}
	stmt.Exec(lastUpdateTime, sourceId)
}
func GetSongIdSngVersionLrcVersionBySourceId(songInfoMerge string, sourceId int) map[string]string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "SELECT id,song_version,lrc_version from song.song_info_merge WHERE source = 3 and source_id = ?"
	sql := "SELECT id,song_version,lrc_version from song." + songInfoMerge + " WHERE source = 3 and source_id = ?"
	//fmt.Println(sql)
	row := util.GetDbSongRead().QueryRow(sql, sourceId)
	Map := make(map[string]string)
	var id string
	var songVersion string
	var LrcVersion string
	row.Scan(&id, &songVersion, &LrcVersion)
	Map["id"] = id
	Map["songVersion"] = songVersion
	Map["LrcVersion"] = LrcVersion
	return Map
}

func UpdateFinallyTmp(songInfoMergeTmp string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	//sql := "delete from song.song_info_merge_tmp_test where status = -99"
	sql := "delete from song." + songInfoMergeTmp + " where status = -99"
	stmt, err := util.GetDbSong().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	stmt.Exec()
}

func GetSongNameAndSingerBySongInfoMergeSourceId(songInfoMerge string, sourceId int) model.IdSongNameSinger {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var model model.IdSongNameSinger
	//sql := "SELECT id, `song_name`, singer from song.song_info_merge_test WHERE source_id = ? "
	sql := "SELECT id, `song_name`, singer from song." + songInfoMerge + " WHERE source_id = ? "
	row := util.GetDbSongRead().QueryRow(sql, sourceId)
	row.Scan(&model.Id, &model.SongName, &model.Singer)
	return model
}

// 查询 曲库金和tv曲库更新歌曲
func GetSongInfoMergeNewSong() []model.SongInfoMerge {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id, source_id, source, word_part, song_name, " +
		"	song_name_phonetic, singer, singer_phonetic, album, " +
		"	line_writer, song_writer, gender, lang, year, dossier_name, " +
		"	is_duet, song_size, lrc_size, background_image_size, icon_size, " +
		"	song_time, status, lrc_type, update_time, intonation, create_date, " +
		"	mv_size, song_path, song_name_stroke_num, singer_stroke_num, song_name_simple," +
		" 	singer_simple, song_version, lrc_version, lrc_channel, lrc_size2, song_chinese_phonetic, " +
		"	singer_chinese_phonetic, search_word, youtube_song_type from song.song_info_merge where create_date like concat(?,'%') order by create_date desc"
	query, _ := util.GetDbSongRead().Query(sql, time.Now().Format("2006-01-02"))
	defer query.Close()
	var list []model.SongInfoMerge
	for query.Next() {
		var info model.SongInfoMerge
		query.Scan(&info.Id, &info.SourceId, &info.Source, &info.WordPart, &info.SongName,
			&info.SongNamePhonetic, &info.Singer, &info.SingerPhonetic, &info.Album,
			&info.LineWriter, &info.SongWriter, &info.Gender, &info.Lang, &info.Year, &info.DossierName,
			&info.IsDuet, &info.SongSize, &info.LrcSize, &info.BackgroundImageSize, &info.IconSize,
			&info.SongTime, &info.Status, &info.LrcType, &info.UpdateTime, &info.Intonation, &info.CreateDate,
			&info.MvSize, &info.SongPath, &info.SongNameStrokeNum, &info.SingerStrokeNum, &info.SongNameSimple,
			&info.SingerSimple, &info.SongVersion, &info.LrcVersion, &info.LrcChannel, &info.LrcSize2, &info.SongPhinesePhonetic,
			&info.SingerChinesePhonetic, &info.SearchWord, &info.YoutubeSongType)
		list = append(list, info)
	}
	return list
}
func GetSearcherKeyWordTvBySourceId(sourceId int) model.SearcherKeyWord {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id, source_id, song_name, singer, song_name_keyword, singer_keyword, digital_id from song.search_key_word_tv where source_id = ?"
	query := util.GetDbSongRead().QueryRow(sql, sourceId)
	var info model.SearcherKeyWord
	query.Scan(&info.Id, &info.SourceId, &info.SongName, &info.Singer, &info.SongNameKeyWord, &info.SingerKeyWord, &info.DigitalId)
	return info
}
