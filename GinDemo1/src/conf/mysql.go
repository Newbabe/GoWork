package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	DB, _ = sql.Open("mysql", "cowboy:123456Cxp@@tcp(rm-uf627n6r3e46m3b4pso.mysql.rds.aliyuncs.com:3306)/oa?charset=utf8")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
