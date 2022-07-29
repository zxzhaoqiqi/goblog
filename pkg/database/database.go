package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/zxzhaoqiqi/goblog/pkg/logger"
)

// 连接池对象
var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error

	// 设置数据库链接信息
	config := mysql.Config{
		User:                 "root",
		Passwd:               "zhaoxian",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	fmt.Println(config.FormatDSN())
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.Error(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)

	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)

	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 常识链接，失败会报错
	err = DB.Ping()
	logger.Error(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	);`

	_, err := DB.Exec(createArticlesSQL)
	logger.Error(err)
}
