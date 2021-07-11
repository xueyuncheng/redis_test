package http

import (
	"redis_test/internal/ecode"
	"redis_test/internal/log"
	"redis_test/model"
	"redis_test/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListUser(ctx *gin.Context) interface{} {
	req := &model.ListUserReq{}
	if err := ctx.ShouldBind(req); err != nil {
		log.Sugar.Error("参数绑定错误", zap.Error(err))
		return ecode.ErrInvalidParam.WithError(err)
	}

	users, err := service.ListUser(ctx, req)
	if err != nil {
		return ecode.ErrSystemError.WithError(err)
	}

	count, err := service.ListUserCount(ctx, req)
	if err != nil {
		return ecode.ErrSystemError.WithError(err)
	}

	return ecode.OK.WithPageData(users, count, req.PageSize)
}
