package database

import (
	"fmt"
	"redis_test/config"
	"redis_test/internal/log"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(database *config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		database.UserName, database.Password, database.Host, database.Port, database.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Sugar.Fatal("初始化数据库连接错误", zap.Error(err))
	}

	return db, nil
}
