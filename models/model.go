package models

import (
	"github.com/zxzhaoqiqi/goblog/pkg/logger"
	"github.com/zxzhaoqiqi/goblog/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:zhaoxian@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})

	logger.Error(err)

	return DB
}

type BaseMode struct {
	ID uint64
}

func (a BaseMode) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
