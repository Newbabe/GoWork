package malysiaDao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func DelMalysiaMergeSongDb(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DELETE from " + tableName + "  where source  = 1"
	_, err := util.MalysiaURDUtil(sql)
	if err != nil {
		panic(err)
		return
	}

}

func AddSongInfoMergeTmp(tableName string, songInfo model.SongInfoMergeTvTmp) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := " INSERT ignore  INTO " + tableName + " (`id`,`source_id`,`source`,`word_part`,`song_name`," +
		"  `song_name_phonetic`,`singer`,`singer_phonetic`,`album`,`line_writer`,`song_writer`,`gender`,`lang`,`year`," +
		"  `dossier_name`,`is_duet`,`song_size`,`lrc_size`,`background_image_size`,`icon_size`,`song_time`,`status`," +
		"  `lrc_type`,`update_time`,`intonation`,`create_date`,`mv_size`,`song_path`,`song_name_stroke_num`," +
		"  `singer_stroke_num`,`song_name_simple`,`singer_simple`,`song_version`,`lrc_version`,`lrc_channel`,`lrc_size2`," +
		"  `song_chinese_phonetic`,`singer_chinese_phonetic`,`search_word`,`youtube_song_type`) " +
		"  values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	prepare, err := util.GetDbMalysia().Prepare(sql)
	if err != nil {
		return 0
	}
	defer prepare.Close()
	result, err := prepare.Exec(songInfo.Id, songInfo.SourceId, songInfo.Source, songInfo.WordPart, songInfo.SongName,
		songInfo.SongNamePhonetic, songInfo.Singer, songInfo.SingerNamePhonetic, songInfo.Album, songInfo.LineWriter, songInfo.SongWrite, songInfo.Gender, songInfo.Lang, songInfo.Year,
		songInfo.DossierName, songInfo.IsDuet, songInfo.SongSize, songInfo.LrcSize, songInfo.BackgroundImageSize, songInfo.IconSize, songInfo.SongTime, songInfo.Status,
		songInfo.LrcType, songInfo.UpdateTime, songInfo.Intonation, songInfo.CreateDate, songInfo.MvSize, songInfo.SongPath, songInfo.SongNameStrokeNum,
		songInfo.SingerStrokeNum, songInfo.SongNameSimple, songInfo.SingerSimple, songInfo.SongVersion, songInfo.LrcVersion, songInfo.LrcChannel, songInfo.LrcSize2,
		songInfo.SongChinesePhonetic, songInfo.SingerChinesePhonetic, songInfo.SearchWord, songInfo.YouTubeSongType)

	if err != nil {
		panic(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	songId := int(id)
	return songId
}
func AddSongInfoMergeTmp2(tableName string, songInfo model.SongInfoMergeTmp) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sql := "INSERT ignore  INTO " + tableName + " (  `source_id`," + "  `source`," +
		"  `word_part`," + "  `song_name`," + "  `song_name_phonetic`," + "  `singer`," +
		"  `singer_phonetic`," + "  `album`," + "  `line_writer`," + "  `song_writer`," + "  `gender`," +
		"  `lang`," + "  `year`," + "  `dossier_name`," + "  `is_duet`," + "  `song_size`," + "  `lrc_size`," +
		"  `background_image_size`," + "  `icon_size`," + "  `song_time`," + "  `status`," + "  `lrc_type`," +
		"  `update_time`," + "  `intonation`," + "  `create_date`," + "  `mv_size`," + "  `song_path`," +
		"  `song_name_stroke_num`," + "  `singer_stroke_num`," + "  `song_name_simple`," +
		"  `singer_simple`," + "  `song_version`," +
		"  `lrc_version`,lrc_channel, lrc_size2,song_chinese_phonetic,singer_chinese_phonetic," +
		"search_word,youtube_song_type" + ") " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) "
	prepare, err := util.GetDbMalysia().Prepare(sql)
	if err != nil {
		panic(err)
		return 0
	}
	defer prepare.Close()
	result, err := prepare.Exec(songInfo.SourceId, songInfo.Source,
		songInfo.WordPart, songInfo.SongName, songInfo.SongNamePhonetic, songInfo.Singer,
		songInfo.SingerNamePhonetic, songInfo.Album, songInfo.LineWriter, songInfo.SongWrite, songInfo.Gender,
		songInfo.Lang, songInfo.Year, songInfo.DossierName, songInfo.IsDuet, songInfo.SongSize, songInfo.LrcSize,
		songInfo.BackgroundImageSize, songInfo.IconSize, songInfo.SongTime, songInfo.Status, songInfo.LrcType,
		songInfo.UpdateTime, songInfo.Intonation, songInfo.CreateDate, songInfo.MvSize, songInfo.SongPath,
		songInfo.SongNameStrokeNum, songInfo.SingerStrokeNum, songInfo.SongNameSimple,
		songInfo.SingerSimple, songInfo.SongVersion,
		songInfo.LrcVersion, songInfo.LrcChannel, songInfo.LrcSize2, songInfo.SongChinesePhonetic, songInfo.SingerNamePhonetic,
		songInfo.SearchWord, songInfo.YouTubeSongType)

	if err != nil {
		panic(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	i := int(id)
	return i
}

func UpdateMergeSongDb(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update " + tableName + " SET status = -1 where source  = 3 and status = 1"
	_, err := util.MalysiaURDUtil(sql)
	if err != nil {
		panic(err)
		return
	}

}
func UpdateMergeSongDb2(tableName string, sourceId int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "update " + tableName + " SET status = 1 where source_id=?"
	_, err := util.MalysiaURDUtil(sql, sourceId)
	if err != nil {
		panic(err)
		return
	}

}

func UpdateSongInfoMergeTmp(tableName string, songInfo model.SongInfoMergeTmp) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	updateSql := "update " + tableName + "  SET  `word_part` = ?, `song_name` = ?," +
		"  `singer` = ?,  `album` = ?,  `line_writer` = ?,  `song_writer` = ?," +
		"  `gender` = ?, `lang` = ?,  `year` = ?, `dossier_name` = ?,  `is_duet` = ?," +
		"  `song_size` = ?, `lrc_size` = ?,  `background_image_size` = ?,  `icon_size` = ?," +
		"  `song_time` = ?, `status` = ?,  `lrc_type` = ?, `update_time` = ?, `intonation` = ?," +
		"  `mv_size` = ?,  `song_path` = ?,  `song_name_stroke_num` = ?," +
		"  `singer_stroke_num` = ?,  `song_name_simple` = ?," + "  `singer_simple` = ?," +
		"  lrc_channel = ?, lrc_size2 = ?,song_chinese_phonetic = ?,singer_chinese_phonetic = ?," +
		"search_word=?,youtube_song_type = ? where source = 3 and `source_id` = ? "
	prepare, err := util.GetDbMalysia().Prepare(updateSql)
	if err != nil {
		return 0
	}
	defer prepare.Close()
	result, err := prepare.Exec(songInfo.WordPart, songInfo.SongName,
		songInfo.Singer, songInfo.Album, songInfo.LineWriter, songInfo.SongWrite,
		songInfo.Gender, songInfo.Lang, songInfo.Year, songInfo.DossierName, songInfo.IsDuet,
		songInfo.SongSize, songInfo.LrcSize, songInfo.BackgroundImageSize, songInfo.IconSize,
		songInfo.SongTime, songInfo.Status, songInfo.LrcType, songInfo.UpdateTime, songInfo.Intonation,
		songInfo.MvSize, songInfo.SongPath, songInfo.SongNameStrokeNum,
		songInfo.SingerStrokeNum, songInfo.SongNameSimple, songInfo.SingerSimple,
		songInfo.LrcChannel, songInfo.LrcSize2, songInfo.SongChinesePhonetic, songInfo.SingerChinesePhonetic,
		songInfo.SearchWord, songInfo.YouTubeSongType, songInfo.SourceId)
	if err != nil {
		panic(err)
		return 0
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	i := int(id)
	return i
}
