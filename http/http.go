package http

import (
	"net"
	"redis_test/config"
	"redis_test/service"

	"github.com/gin-gonic/gin"
)

func InitHttp(cfg *config.Config) {
	service.InitService(cfg)

	r := gin.Default()
	initRouter(r)
	r.Run(net.JoinHostPort(cfg.Http.Host, cfg.Http.Port))
}

func initRouter(r *gin.Engine) {
	api := r.Group("/api")

	users := api.Group("/users")
	{
		users.GET("", wrap(ListUser))
	}
}

func wrap(f func(ctx *gin.Context) interface{}) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		f(ctx)
	}
}
