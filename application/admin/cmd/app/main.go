package main

import (
	"log"

	config "github.com/ccjshop/go-mall/application/admin/config"
	"github.com/ccjshop/go-mall/application/admin/internal/app"
)

func main() {
	// 加载配置
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// 启动
	app.Run(cfg)
}
