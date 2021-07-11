package cache

import (
	"fmt"
	"redis_test/config"

	"github.com/go-redis/redis/v8"
)

func NewCache(configCache *config.Cache) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", configCache.Host, configCache.Port),
		Username: configCache.UserName,
		Password: configCache.Password,
	})

	return rdb, nil
}
