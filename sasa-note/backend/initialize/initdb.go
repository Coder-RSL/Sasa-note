package initialize

import (
	"backend/common"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// InitDB
func InitDB() {
	common.Mysql = InitMySQL()
	InitRedis()
}

func InitMySQL() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		"root", "hzau.0213", "47.120.52.204:3306", "gintest", "utf8")

	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err.Error())
	} else {
		conn.SetConnMaxLifetime(7 * time.Second) //设置空闲时间，这个是比mysql 主动断开的时候短
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)
		fmt.Println("数据库连接成功")
		return conn
	}
}
func InitRedis() {

}
