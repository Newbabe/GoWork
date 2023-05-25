package model

import "time"

type SongInfoMerger struct {
	Status     int
	SourceId   int
	SongName   string
	Singer     string
	SongSize   int
	LrcSize    int
	LrcSize2   int
	UpdateTime time.Time
}
type SourceIdAndChannel struct {
	SourceId   int
	LrcChannel int
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
