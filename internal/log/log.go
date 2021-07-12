package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func InitLog() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("日志模块启动错误", err)
		os.Exit(1)
	}

	defer logger.Sync() // flushes buffer, if any
	Sugar = logger.Sugar()
}
