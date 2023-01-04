package bootstrap

import (
	"time"

	"github.com/zxzhaoqiqi/goblog/models"
)

func SetupDB() {

	// 建立连接
	db := models.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
