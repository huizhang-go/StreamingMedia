package dbops

import (
	"database/sql" // 自带的
	_ "github.com/go-sql-driver/mysql" // 第三方的
)

var (
	dbConn *sql.DB
	err error
)

// 初始化连接,避免重复连接，自带的已经给做好了
func init()  {
	dbConn, err = sql.Open("mysql", "root:guoyuzhao123@#@tcp(localhost:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
