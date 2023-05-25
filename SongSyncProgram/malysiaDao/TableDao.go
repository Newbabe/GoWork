package malysiaDao

import (
	"SongSyncProgram/util"
	"fmt"
	"runtime/debug"
)

func CreateTable(table, tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "CREATE TABLE " + table + " SELECT * FROM " + tableName
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()

}

func TruncateTable(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "truncate " + tableName
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()
}

func CopyTable(OldTableName, newTableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT INTO " + newTableName + " SELECT * FROM " + OldTableName
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()
}
func DropTable(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "DROP TABLE  if EXISTS " + tableName
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()
}
func DropId(tableName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "alter table " + tableName + " drop id"
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()
}
func AddPrimaryKey(tabName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "alter table " + tabName + " add id int not null primary key auto_increment first;"
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)
		return
	}
	prepare.Exec()
}

func AddTableBByTableA(tableNameA, tableNameB string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	sql := "INSERT IGNORE INTO " + tableNameB + " SELECT * FROM " + tableNameA
	prepare, err := util.GetDbMalysia().Prepare(sql)
	defer prepare.Close()
	if err != nil {
		panic(err)

		return
	}
	prepare.Exec()
}
