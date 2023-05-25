package malysiaDao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

//查询台湾数据
func SelectTwSongInfoDb() []model.SongInfoMergeTvTmp {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "SELECT `id`,`source_id`,`source`,`word_part`,`song_name`," +
		" `song_name_phonetic`,`singer`,`singer_phonetic`,`album`," +
		" `line_writer`,`song_writer`,`gender`,`lang`,`year`,`dossier_name`," +
		" `is_duet`,`song_size`,`lrc_size`,`background_image_size`,`icon_size`," +
		" `song_time`,`status`,`lrc_type`,`update_time`,`intonation`,`create_date`," +
		" `mv_size`,`song_path`,`song_name_stroke_num`,`singer_stroke_num`," +
		" `song_name_simple`,`singer_simple`,`song_version`,`lrc_version`," +
		" `lrc_channel`,`lrc_size2`,`song_chinese_phonetic`,`singer_chinese_phonetic`," +
		" `search_word`,`youtube_song_type`" +
		"  FROM `song_info_merge_tv` WHERE source = 1 "
	query, err := util.GetDbSongRead().Query(sql)

	defer query.Close()
	if err != nil {
		panic(err)
		return []model.SongInfoMergeTvTmp{}
	}
	var songInfoList []model.SongInfoMergeTvTmp
	for query.Next() {
		var songInfo model.SongInfoMergeTvTmp
		query.Scan(&songInfo.Id, &songInfo.SourceId, &songInfo.Source, &songInfo.WordPart, &songInfo.SongName,
			&songInfo.SongNamePhonetic, &songInfo.Singer, &songInfo.SingerNamePhonetic, &songInfo.Album,
			&songInfo.LineWriter, &songInfo.SongWrite, &songInfo.Gender, &songInfo.Lang, &songInfo.Year, &songInfo.DossierName,
			&songInfo.IsDuet, &songInfo.SongSize, &songInfo.LrcSize, &songInfo.BackgroundImageSize, &songInfo.IconSize,
			&songInfo.SongTime, &songInfo.Status, &songInfo.LrcType, &songInfo.UpdateTime, &songInfo.Intonation, &songInfo.CreateDate,
			&songInfo.MvSize, &songInfo.SongPath, &songInfo.SongNameStrokeNum, &songInfo.SingerStrokeNum,
			&songInfo.SongNameSimple, &songInfo.SingerSimple, &songInfo.SongVersion, &songInfo.LrcVersion,
			&songInfo.LrcChannel, &songInfo.LrcSize2, &songInfo.SongChinesePhonetic, &songInfo.SingerChinesePhonetic,
			&songInfo.SearchWord, &songInfo.YouTubeSongType)
		songInfoList = append(songInfoList, songInfo)
	}
	return songInfoList
}
