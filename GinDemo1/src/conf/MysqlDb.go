package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Debug = false

var dbHerook *sql.DB
var dbHeroOkWeb *sql.DB
var dbSong *sql.DB
var dbSongRead *sql.DB
var dbCompetition *sql.DB
var dbHeroOkShareComment *sql.DB
var dbLrcBackstage *sql.DB
var dbLogHerook *sql.DB
var dbLogHerookRead *sql.DB
var dbHerookWebRead *sql.DB
var dbHerookRecordData *sql.DB
var dbHerookShareCommentRead *sql.DB
var dbHerookStatistic *sql.DB
var dbHerookStatisticRead *sql.DB
var dbChat *sql.DB
var dbHerookRead *sql.DB
var dbHerookSongStar *sql.DB
var dbMemory *sql.DB

//var lock *sync.Mutex = &sync.Mutex {}
func init() {
	url := "herooktest.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"
	dburl := "herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"
	readUrl := "herook2read.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"
	//马来西亚
	//url := "herook-mylaysia.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"
	//dburl := "herook-mylaysia.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"

	if Debug {
		dbHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerook.SetMaxOpenConns(30)
		dbHerook.SetMaxIdleConns(5)
		dbHerook.Ping()

		dbSong, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSong.SetMaxOpenConns(30)
		dbSong.SetMaxIdleConns(5)
		dbSong.Ping()

		dbHeroOkWeb, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkWeb.SetMaxOpenConns(30)
		dbHeroOkWeb.SetMaxIdleConns(5)
		dbHeroOkWeb.Ping()

		dbCompetition, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/competition?charset=utf8mb4&parseTime=true&loc=Local")
		dbCompetition.SetMaxOpenConns(10)
		dbCompetition.SetMaxIdleConns(5)
		dbCompetition.Ping()

		dbHeroOkShareComment, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook_share_comment?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkShareComment.SetMaxOpenConns(5)
		dbHeroOkShareComment.SetMaxIdleConns(5)
		dbHeroOkShareComment.Ping()

		dbLrcBackstage, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/lrc_backstage?charset=utf8mb4&parseTime=true&loc=Local")
		dbLrcBackstage.SetMaxOpenConns(5)
		dbLrcBackstage.SetMaxIdleConns(5)
		dbLrcBackstage.Ping()

		dbLogHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/log_herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbLogHerook.SetMaxOpenConns(5)
		dbLogHerook.SetMaxIdleConns(5)
		dbLogHerook.Ping()

		dbLogHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/log_herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbLogHerookRead.SetMaxOpenConns(2)
		dbLogHerookRead.SetMaxIdleConns(2)
		dbLogHerookRead.Ping()

		dbHerookRecordData, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook_record_data?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRecordData.SetMaxOpenConns(5)
		dbHerookRecordData.SetMaxIdleConns(5)
		dbHerookRecordData.Ping()

		dbHerookSongStar, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook_song_star?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookSongStar.SetMaxOpenConns(5)
		dbHerookSongStar.SetMaxIdleConns(5)
		dbHerookSongStar.Ping()

		dbMemory, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/memory_db?charset=utf8mb4&parseTime=true&loc=Local")
		dbMemory.SetMaxOpenConns(30)
		dbMemory.SetMaxIdleConns(5)
		dbMemory.Ping()

	} else {
		dbHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerook.SetMaxOpenConns(10)
		dbHerook.SetMaxIdleConns(2)
		dbHerook.Ping()

		dbSong, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSong.SetMaxOpenConns(10)
		dbSong.SetMaxIdleConns(2)
		dbSong.Ping()

		dbSongRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSongRead.SetMaxOpenConns(10)
		dbSongRead.SetMaxIdleConns(2)
		dbSongRead.Ping()

		dbHeroOkWeb, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkWeb.SetMaxOpenConns(10)
		dbHeroOkWeb.SetMaxIdleConns(2)
		dbHeroOkWeb.Ping()

		dbCompetition, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/competition?charset=utf8mb4&parseTime=true&loc=Local")
		dbCompetition.SetMaxOpenConns(10)
		dbCompetition.SetMaxIdleConns(2)
		dbCompetition.Ping()

		dbHeroOkShareComment, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook_share_comment?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkShareComment.SetMaxOpenConns(5)
		dbHeroOkShareComment.SetMaxIdleConns(2)
		dbHeroOkShareComment.Ping()

		dbLrcBackstage, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/lrc_backstage?charset=utf8mb4&parseTime=true&loc=Local")
		dbLrcBackstage.SetMaxOpenConns(5)
		dbLrcBackstage.SetMaxIdleConns(2)
		dbLrcBackstage.Ping()

		dbLogHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/log_herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbLogHerook.SetMaxOpenConns(5)
		dbLogHerook.SetMaxIdleConns(2)
		dbLogHerook.Ping()

		dbLogHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/log_herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbLogHerookRead.SetMaxOpenConns(2)
		dbLogHerookRead.SetMaxIdleConns(2)
		dbLogHerookRead.Ping()

		dbHerookWebRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookWebRead.SetMaxOpenConns(5)
		dbHerookWebRead.SetMaxIdleConns(2)
		dbHerookWebRead.Ping()

		dbHerookRecordData, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook_record_data?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRecordData.SetMaxOpenConns(5)
		dbHerookRecordData.SetMaxIdleConns(2)
		dbHerookRecordData.Ping()

		dbHerookShareCommentRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook_share_comment?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookShareCommentRead.SetMaxOpenConns(10)
		dbHerookShareCommentRead.SetMaxIdleConns(2)
		dbHerookShareCommentRead.Ping()

		dbHerookStatistic, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook_statistic?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookStatistic.SetMaxOpenConns(5)
		dbHerookStatistic.SetMaxIdleConns(2)
		dbHerookStatistic.Ping()

		dbChat, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook_chat?charset=utf8mb4&parseTime=true&loc=Local")
		dbChat.SetMaxOpenConns(5)
		dbChat.SetMaxIdleConns(2)
		dbChat.Ping()

		dbHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRead.SetMaxOpenConns(10)
		dbHerookRead.SetMaxIdleConns(2)
		dbHerookRead.Ping()

		dbHerookSongStar, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/herook_song_star?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookSongStar.SetMaxOpenConns(5)
		dbHerookSongStar.SetMaxIdleConns(2)
		dbHerookSongStar.Ping()

		dbMemory, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+dburl+":3306)/memory_db?charset=utf8mb4&parseTime=true&loc=Local")
		dbMemory.SetMaxOpenConns(5)
		dbMemory.SetMaxIdleConns(2)
		dbMemory.Ping()

		dbHerookStatisticRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook_statistic?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookStatisticRead.SetMaxOpenConns(10)
		dbHerookStatisticRead.SetMaxIdleConns(5)
		dbHerookStatisticRead.Ping()
	}

}

func GetDbHerook() *sql.DB {
	return dbHerook
}

func GetDbSong() *sql.DB {
	return dbSong
}

func GetDbSongRead() *sql.DB {
	return dbSongRead
}

func GetDbHeroOkWeb() *sql.DB {
	return dbHeroOkWeb
}

func GetDbCompetition() *sql.DB {
	return dbCompetition
}

func GetDbHeroOkShareComment() *sql.DB {
	return dbHeroOkShareComment
}
func GetDbLrcBackstage() *sql.DB {
	return dbLrcBackstage
}
func GetDbLogHerook() *sql.DB {
	return dbLogHerook
}

func GetDbLogHerookRead() *sql.DB {
	return dbLogHerookRead
}
func GetDbHerookWebRead() *sql.DB {
	return dbHerookWebRead
}
func GetDbHerookRecordData() *sql.DB {
	return dbHerookRecordData
}
func GetDbShareCommentRead() *sql.DB {
	return dbHerookShareCommentRead
}

func GetDbHerookStatistic() *sql.DB {
	return dbHerookStatistic
}
func GetDbHerookStatisticRead() *sql.DB {
	return dbHerookStatisticRead
}
func GetDbChat() *sql.DB {
	return dbChat
}
func GetDbHerookRead() *sql.DB {
	return dbHerookRead
}

func GetDbHerookSongStar() *sql.DB {
	return dbHerookSongStar
}

func GetDbMemory() *sql.DB {
	return dbMemory
}
