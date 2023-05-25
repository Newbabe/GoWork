package action

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
	"fmt"
	"github.com/vanng822/go-solr/solr"
	"runtime/debug"
	"time"
)

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

const solr_server = "http://172.31.7.72:8983/solr"
const searchKeyword = "search_key_word_tv"
const songInfoMerge = "song_info_merge"
const songInfoMergeTv = "song_info_merge_tv"

// UpdateNewSongSolr 更新新进歌曲到 solr服务器
func UpdateNewSongSolr() {
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
	newSong := dao.GetSongInfoMergeNewSong()
	fmt.Println(time.Now().Format("2006-01-02"), "歌曲数量:", len(newSong))
	for _, info := range newSong {
		UpdateSongInfoMergeNewSongSolr(info, songInfoMergeConfig.NewSolrInterface())
		UpdateSongInfoMergeTVNewSongSolr(info, songInfoMergeTvConfig.NewSolrInterface())
		//通过source_id查询searcher_key_word_tv表数据
		searchKeyWord := dao.GetSearcherKeyWordTvBySourceId(info.SourceId)
		UpdateSearcherKeyWordNewSongSolr(searchKeyWord, SearchKeywordConfig.NewSolrInterface())

	}
}
func UpdateSongInfoMergeNewSongSolr(info model.SongInfoMerge, si *solr.SolrInterface) {
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
func UpdateSongInfoMergeTVNewSongSolr(info model.SongInfoMerge, si *solr.SolrInterface) {
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
func UpdateSearcherKeyWordNewSongSolr(info model.SearcherKeyWord, si *solr.SolrInterface) {
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
