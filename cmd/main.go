package main

import (
	"flag"
	"redis_test/config"
	"redis_test/http"
	"redis_test/internal/log"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

func main() {
	filePath := flag.String("c", "config.toml", "配置文件路径")

	flag.Parse()

	log.InitLog()

	cfg := &config.Config{}
	if _, err := toml.DecodeFile(*filePath, cfg); err != nil {
		log.Sugar.Fatal("解析配置文件错误", zap.Error(err))
		return
	}

	http.InitHttp(cfg)
}
