package service

import (
	"redis_test/cache"
	"redis_test/config"
	"redis_test/database"
	"redis_test/internal/log"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var srv *Service

type Service struct {
	Rdb *redis.Client
	Db  *gorm.DB
}

func NewService(config *config.Config) (*Service, error) {
	rdb, err := cache.NewCache(config.Cache)
	if err != nil {
		log.Sugar.Fatal("新建缓存错误", zap.Error(err))
	}

	db, err := database.NewDatabase(config.Database)
	if err != nil {
		log.Sugar.Fatal("新建数据库错误", zap.Error(err))
	}

	srv := &Service{
		Rdb: rdb,
		Db:  db,
	}

	return srv, nil
}

func InitService(config *config.Config) {
	var err error
	srv, err = NewService(config)
	if err != nil {
		log.Sugar.Fatal("新建服务错误", zap.Error(err))
	}
}
