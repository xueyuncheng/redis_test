package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"redis_test/internal/log"
	"redis_test/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func ListUser(ctx *gin.Context, req *model.User) ([]*model.User, error) {
	users := make([]*model.User, 0, 128)
	if err := srv.Db.Model(&model.User{}).Where("name like ?", "%"+req.Name+"%").Offset(req.Offset()).Limit(req.PageSize).
		Find(&users).Error; err != nil {
		log.Sugar.Error("获取用户错误", zap.Error(err))
		return nil, fmt.Errorf("获取用户错误 %w", err)
	}

	return users, nil
}

func ListUserCount(ctx *gin.Context, req *model.User) (int, error) {
	var count int64
	if err := srv.Db.Model(&model.User{}).Where("name like ?", req.Name).Count(&count).Error; err != nil {
		log.Sugar.Error("获取用户计数错误", zap.Error(err))
		return 0, fmt.Errorf("获取用户计数错误 %w", err)
	}

	return int(count), nil
}

func GetUser(ctx *gin.Context, id int) (*model.User, error) {
	user := &model.User{}
	bs, err := srv.Rdb.Get(ctx, "user:id:"+strconv.Itoa(id)).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Sugar.Error("获取用户信息错误", zap.Error(err))
			return nil, fmt.Errorf("获取用户信息错误 %w", err)
		}

		if err := srv.Db.Where("id = ?", id).First(&user).Error; err != nil {
			log.Sugar.Error("获取用户信息错误", zap.Error(err))
			return nil, fmt.Errorf("获取用户信息错误 %w", err)
		}

		v, err := json.Marshal(user)
		if err != nil {
			log.Sugar.Error("序列化错误", zap.Error(err))
			return nil, fmt.Errorf("序列化错误 %w", err)
		}

		if err := srv.Rdb.Set(ctx, "user:id:"+strconv.Itoa(id), v, 10*time.Hour).Err(); err != nil {
			log.Sugar.Error("缓存set错误", zap.Error(err))
			return nil, fmt.Errorf("缓存set错误 %w", err)
		}

		return user, nil
	}

	if err := json.Unmarshal(bs, user); err != nil {
		log.Sugar.Error("反序列化错误", zap.Error(err))
		return nil, fmt.Errorf("反序列化错误 %w", err)
	}

	return user, nil
}
