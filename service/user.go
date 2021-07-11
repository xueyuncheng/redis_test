package service

import (
	"fmt"
	"redis_test/internal/log"
	"redis_test/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddUser(ctx *gin.Context, req *model.AddUserReq) error {
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := srv.Db.Create(user).Error; err != nil {
		log.Sugar.Error("创建人员错误", zap.Error(err))
		return fmt.Errorf("创建人员错误 %w", err)
	}

	key := fmt.Sprintf("user:" + strconv.Itoa(int(user.ID)))
	if err := srv.Rdb.HMSet(ctx, key, "name", user.Name, "age", user.Age).Err(); err != nil {
		log.Sugar.Error("hset 人员信息错误", zap.Error(err))
		return fmt.Errorf("hset 人员信息错误 %w", err)
	}

	if err := srv.Rdb.Expire(ctx, key, 10*time.Second).Err(); err != nil {
		log.Sugar.Error("expire 设置key的过期时间错误", zap.Error(err))
		return fmt.Errorf("expire 设置key的过期时间错误 %w", err)
	}

	if err := srv.Rdb.LPush(ctx, "user", user.ID).Err(); err != nil {
		log.Sugar.Error("lpush 人员信息错误", zap.Error(err))
		return fmt.Errorf("lpush 人员信息错误 %w", err)
	}

	if err := srv.Rdb.Expire(ctx, "user", 10*time.Second).Err(); err != nil {
		log.Sugar.Error("设置user列表超时时间错误", zap.Error(err))
		return fmt.Errorf("设置user列表超时时间错误 %w", err)
	}

	return nil
}

func ListUser(ctx *gin.Context, req *model.ListUserReq) ([]*model.ListUserResp, error) {
	ids := make([]int, 0, 16)
	if err := srv.Rdb.LRange(ctx, "user", int64(req.Offset()), int64(req.PageSize)).ScanSlice(&ids); err != nil {
		log.Sugar.Error("获取用户列表错误", zap.Error(err))
		return nil, fmt.Errorf("获取用户列表错误 %w", err)
	}

	users := make([]*model.User, 0, 128)
	for _, v := range ids {
		tmp := &model.User{}
		if err := srv.Rdb.HGetAll(ctx, "user:"+strconv.Itoa(v)).Scan(&tmp); err != nil {
			log.Sugar.Error("获取人员信息错误", zap.Error(err))
			return nil, fmt.Errorf("获取人员信息错误 %w", err)
		}

		users = append(users, tmp)
	}

	respUsers := make([]*model.ListUserResp, 0, 16)
	for _, v := range users {
		tmp := &model.ListUserResp{
			ID:   int(v.ID),
			Name: v.Name,
			Age:  v.Age,
		}

		respUsers = append(respUsers, tmp)
	}

	return respUsers, nil
}

func ListUserCount(ctx *gin.Context, req *model.ListUserReq) (int, error) {
	length, err := srv.Rdb.LLen(ctx, "user").Result()
	if err != nil {
		log.Sugar.Error("获取user列表长度错误", zap.Error(err))
		return 0, fmt.Errorf("获取user列表长度错误 %w", err)
	}

	return int(length), nil

}
