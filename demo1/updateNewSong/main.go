package main

import (
	"demo1/util"
	"fmt"
	"github.com/vanng822/go-solr/solr"
	"runtime/debug"
	"time"
)

type song struct {
}

const solr_server = "http://172.31.7.72:8983/solr"
const searchKeyword = "search_key_word_tv"
const songInfoMerge = "song_info_merge"
const songInfoMergeTv = "song_info_merge_tv"

func main() {
	songInfoMergeConfig := sortConfig{
		Host: solr_server,
		Core: songInfoMerge,
	}
	songInfoMergeTvConfig := sortConfig{
		Host: solr_server,
		Core: songInfoMergeTv,
	}
	SearchKeywordConfig := sortConfig{
		Host: solr_server,
		Core: searchKeyword,
	}
	newSong := GetSongInfoMergeNewSong()
	fmt.Println("len", len(newSong))
	for _, info := range newSong {
		UpdateSongInfoMergeNewSongSolr(info, songInfoMergeConfig.NewSolrInterface())
		UpdateSongInfoMergeTVNewSongSolr(info, songInfoMergeTvConfig.NewSolrInterface())
		//通过source_id查询searcher_key_word_tv表数据
		searchKeyWord := GetSearcherKeyWordTvBySourceId(info.SourceId)
		UpdateSearcherKeyWordNewSongSolr(searchKeyWord, SearchKeywordConfig.NewSolrInterface())

	}

}

type sortConfig struct {
	Host string
	Core string
}

func (s *sortConfig) NewSolrInterface() *solr.SolrInterface {
	solrInterface, err := solr.NewSolrInterface(s.Host, s.Core)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return solrInterface
}

// 查询 曲库金和tv曲库更新歌曲
func GetSongInfoMergeNewSong() []SongInfoMerge {
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
	query, _ := util.GetSongRead().Query(sql, time.Now().Format("2006-01-02"))
	defer query.Close()
	var list []SongInfoMerge
	for query.Next() {
		var info SongInfoMerge
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
func GetSearcherKeyWordTvBySourceId(sourceId int) SearcherKeyWord {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "select id, source_id, song_name, singer, song_name_keyword, singer_keyword, digital_id from song.search_key_word_tv where source_id = ?"
	query := util.GetSongRead().QueryRow(sql, sourceId)
	var info SearcherKeyWord
	query.Scan(&info.Id, &info.SourceId, &info.SongName, &info.Singer, &info.SongNameKeyWord, &info.SingerKeyWord, &info.DigitalId)
	return info
}

func UpdateSongInfoMergeNewSongSolr(info SongInfoMerge, si *solr.SolrInterface) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var solrMapList []solr.Document
	var solrMap = solr.Document{}

	solrMap["id"] = info.Id
	solrMap["source"] = info.Source
	solrMap["song_name"] = info.SongName
	solrMap["singer"] = info.Singer
	solrMap["singer_phonetic"] = info.SingerPhonetic
	solrMap["song_name_phonetic"] = info.SongNamePhonetic
	solrMap["song_name_stroke_num"] = info.SongNameStrokeNum
	solrMap["singer_stroke_num"] = info.SingerStrokeNum
	solrMap["song_name_simple"] = info.SongNameSimple
	solrMap["singer_simple"] = info.SingerSimple
	solrMap["song_chinese_phonetic"] = info.SongPhinesePhonetic
	solrMap["singer_chinese_phonetic"] = info.SingerChinesePhonetic
	solrMap["song_version"] = info.SongVersion
	solrMapList = append(solrMapList, solrMap)

	add, err := si.Add(solrMapList, 0, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(add.Result)
	}
	_, err = si.Commit()
	if err != nil {
		fmt.Println(err)
	}
}
func UpdateSongInfoMergeTVNewSongSolr(info SongInfoMerge, si *solr.SolrInterface) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var solrMapList []solr.Document
	var solrMap = solr.Document{}

	solrMap["id"] = info.Id
	solrMap["source"] = info.Source
	solrMap["song_name"] = info.SongName
	solrMap["singer"] = info.Singer
	solrMap["singer_phonetic"] = info.SingerPhonetic
	solrMap["song_name_phonetic"] = info.SongNamePhonetic
	solrMap["song_name_stroke_num"] = info.SongNameStrokeNum
	solrMap["singer_stroke_num"] = info.SingerStrokeNum
	solrMap["song_name_simple"] = info.SongNameSimple
	solrMap["singer_simple"] = info.SingerSimple
	solrMap["song_chinese_phonetic"] = info.SongPhinesePhonetic
	solrMap["singer_chinese_phonetic"] = info.SingerChinesePhonetic
	solrMap["search_word"] = info.SearchWord
	solrMapList = append(solrMapList, solrMap)

	add, err := si.Add(solrMapList, 0, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(add.Result)
	}
	_, err = si.Commit()
	if err != nil {
		fmt.Println(err)
	}
}
func UpdateSearcherKeyWordNewSongSolr(info SearcherKeyWord, si *solr.SolrInterface) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	var solrMapList []solr.Document
	var solrMap = solr.Document{}
	solrMap["id"] = info.Id
	solrMap["source_id"] = info.SourceId
	solrMap["song_name"] = info.SongName
	solrMap["singer"] = info.Singer
	solrMap["song_name_keyword"] = info.SongNameKeyWord
	solrMap["singer_keyword"] = info.SongNameKeyWord
	solrMap["digital_id"] = info.DigitalId
	solrMapList = append(solrMapList, solrMap)

	add, err := si.Add(solrMapList, 0, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println(add.Result)
	}
	_, err = si.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

type SongInfoMerge struct {
	Id                    int
	SourceId              int
	Source                int
	WordPart              int
	SongName              string
	SongNamePhonetic      string
	Singer                string
	SingerPhonetic        string
	Album                 string
	LineWriter            string
	SongWriter            string
	Gender                int
	Lang                  int
	Year                  int
	DossierName           string
	IsDuet                int
	SongSize              int
	LrcSize               int
	BackgroundImageSize   int
	IconSize              int
	SongTime              int
	Status                int
	LrcType               int
	UpdateTime            time.Time
	Intonation            int
	CreateDate            time.Time
	MvSize                int
	SongPath              string
	SongNameStrokeNum     int
	SingerStrokeNum       int
	SongNameSimple        string
	SingerSimple          string
	SongVersion           string
	LrcVersion            string
	LrcChannel            int
	LrcSize2              int
	SongPhinesePhonetic   string
	SingerChinesePhonetic string
	SearchWord            string
	YoutubeSongType       int
}

type SearcherKeyWord struct {
	Id              int
	SourceId        int
	SongName        string
	Singer          string
	SongNameKeyWord string
	SingerKeyWord   string
	DigitalId       int
}
