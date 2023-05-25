package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Debug = false

var dbHerook *sql.DB
var DbSongTest *sql.DB
var Malysia *sql.DB

var dbHerookWebRead *sql.DB
var dbHeroOkWeb *sql.DB
var dbSong *sql.DB
var dbSongRead *sql.DB

var dbHerookRead *sql.DB

var dbHerookTvTest *sql.DB

//var lock *sync.Mutex = &sync.Mutex {}
func init() {
	readUrl := "herook2read.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"
	//var dburl = "herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"

	if Debug {
		fmt.Println("测试数据库")
		var url = "herooktest.c9ni3auhkoyx.us-east-2.rds.amazonaws.com"

		dbHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerook.SetMaxOpenConns(10)
		dbHerook.SetMaxIdleConns(5)
		dbHerook.Ping()
		//测试数据库
		dbSong, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSong.SetMaxOpenConns(5)
		dbSong.SetMaxIdleConns(2)
		dbSong.Ping()

		DbSongTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		DbSongTest.SetMaxOpenConns(10)
		DbSongTest.SetMaxIdleConns(5)
		DbSongTest.Ping()

		dbHerookTvTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/herook_tv?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookTvTest.SetMaxOpenConns(10)
		dbHerookTvTest.SetMaxIdleConns(5)
		dbHerookTvTest.Ping()

		dbSongRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSongRead.SetMaxOpenConns(5)
		dbSongRead.SetMaxIdleConns(2)
		dbSongRead.Ping()

		dbHerookTvTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herooktest.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook_tv?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookTvTest.SetMaxOpenConns(5)
		dbHerookTvTest.SetMaxIdleConns(2)
		dbHerookTvTest.Ping()

		dbHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRead.SetMaxOpenConns(5)
		dbHerookRead.SetMaxIdleConns(2)
		dbHerookRead.Ping()

		dbHeroOkWeb, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+url+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkWeb.SetMaxOpenConns(10)
		dbHeroOkWeb.SetMaxIdleConns(5)
		dbHeroOkWeb.Ping()
		//【，测试环境测试数据库】
		DbSongTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herooktest.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		DbSongTest.SetMaxOpenConns(10)
		DbSongTest.SetMaxIdleConns(5)
		DbSongTest.Ping()

	} else {

		Malysia, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook-mylaysia-cluster-1.cluster-c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		Malysia.SetMaxOpenConns(5)
		Malysia.SetMaxIdleConns(2)
		Malysia.Ping()

		dbHerook, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerook.SetMaxOpenConns(5)
		dbHerook.SetMaxIdleConns(2)
		dbHerook.Ping()
		//【正式环境用的是正式数据库】
		DbSongTest, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		DbSongTest.SetMaxOpenConns(10)
		DbSongTest.SetMaxIdleConns(5)
		DbSongTest.Ping()

		dbHerookRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/herook?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookRead.SetMaxOpenConns(5)
		dbHerookRead.SetMaxIdleConns(2)
		dbHerookRead.Ping()

		dbHerookWebRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHerookWebRead.SetMaxOpenConns(5)
		dbHerookWebRead.SetMaxIdleConns(2)
		dbHerookWebRead.Ping()

		dbSong, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSong.SetMaxOpenConns(5)
		dbSong.SetMaxIdleConns(2)
		dbSong.Ping()

		dbSongRead, _ = sql.Open("mysql", "root:Mykuro1818@tcp("+readUrl+":3306)/song?charset=utf8mb4&parseTime=true&loc=Local")
		dbSongRead.SetMaxOpenConns(5)
		dbSongRead.SetMaxIdleConns(2)
		dbSongRead.Ping()

		dbHeroOkWeb, _ = sql.Open("mysql", "root:Mykuro1818@tcp(herook2.c9ni3auhkoyx.us-east-2.rds.amazonaws.com:3306)/hero_ok_web?charset=utf8mb4&parseTime=true&loc=Local")
		dbHeroOkWeb.SetMaxOpenConns(5)
		dbHeroOkWeb.SetMaxIdleConns(2)
		dbHeroOkWeb.Ping()

	}

}

func GetDbSong() *sql.DB {
	return dbSong
}
func GetDbMalysia() *sql.DB {
	return Malysia
}

func GetDbSongRead() *sql.DB {
	return dbSongRead
}
func GetDbSongTest() *sql.DB {
	return DbSongTest
}
