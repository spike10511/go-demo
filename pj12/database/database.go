package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dsn := "root:P110040593abc@tcp(127.0.0.1:3306)/pj12db?charset=utf8mb4&parseTime=True&loc=Local" // 数据库信息
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 连接数据库
	return DB.Ping()
}
