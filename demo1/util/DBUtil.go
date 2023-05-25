package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Debug = false
var dbHerook *sql.DB
var dbHerookTest *sql.DB
var dbHerooWebRead *sql.DB
var dbSong *sql.DB
var dbSongTest *sql.DB
var dbSongRead *sql.DB
var dbHerookRead *sql.DB
var dbHerookTv *sql.DB
var dbHerookTvtest *sql.DB
var dbHerookStatisticRead *sql.DB
var dbHerookStatistic *sql.DB

// var lock *sync.Mutex = &sync.Mutex {}
func init() {
	readUrl := "herook2read.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"

	if Debug {
		var testurl = "herooktest.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"

		dbHerookTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+testurl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookTest.SetMaxOpenConns(10)
		dbHerookTest.SetMaxIdleConns(5)
		dbHerookTest.Ping()

		dbHerookTv, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+testurl+":3306)/herook_tv?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookTv.SetMaxOpenConns(10)
		dbHerookTv.SetMaxIdleConns(5)
		dbHerookTv.Ping()

		dbHerookTvtest, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+testurl+":3306)/herook_tv?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookTvtest.SetMaxOpenConns(10)
		dbHerookTvtest.SetMaxIdleConns(5)
		dbHerookTvtest.Ping()

		dbSongRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2read.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSongRead.SetMaxOpenConns(5)
		dbSongRead.SetMaxIdleConns(2)
		dbSongRead.Ping()
		dbHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2read.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRead.SetMaxOpenConns(5)
		dbHerookRead.SetMaxIdleConns(2)
		dbHerookRead.Ping()
	} else {

		dbHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerook.SetMaxOpenConns(5)
		dbHerook.SetMaxIdleConns(2)
		dbHerook.Ping()

		dbHerooWebRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerooWebRead.SetMaxOpenConns(10)
		dbHerooWebRead.SetMaxIdleConns(5)
		dbHerooWebRead.Ping()

		dbSong, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSong.SetMaxOpenConns(5)
		dbSong.SetMaxIdleConns(2)
		dbSong.Ping()

		dbSongRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSongRead.SetMaxOpenConns(5)
		dbSongRead.SetMaxIdleConns(2)
		dbSongRead.Ping()

		dbHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRead.SetMaxOpenConns(5)
		dbHerookRead.SetMaxIdleConns(2)
		dbHerookRead.Ping()
		dbHerookStatisticRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook_statistic?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookStatisticRead.SetMaxOpenConns(10)
		dbHerookStatisticRead.SetMaxIdleConns(5)
		dbHerookStatisticRead.Ping()
		dbHerookStatistic, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook_statistic?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookStatistic.SetMaxOpenConns(10)
		dbHerookStatistic.SetMaxIdleConns(5)
		dbHerookStatistic.Ping()

	}

}

func GetDbHeroOK() *sql.DB {
	return dbHerook
}
func GetDbHeroOkReab() *sql.DB {
	return dbHerookRead
}
func GetDbHerookWebRead() *sql.DB {
	return dbHerooWebRead
}
func GetSong() *sql.DB {
	return dbSong
}
func GetSongRead() *sql.DB {
	return dbSongRead
}
func GetDbHerookStatisticRead() *sql.DB {
	return dbHerookStatisticRead
}
func GetDbHerookStatistic() *sql.DB {
	return dbHerookStatistic
}
