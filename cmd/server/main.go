package main

import (
	"gin-template/internal/router"
	"gin-template/pkg/logger"
)

func main() {
	// 初始化日志
	logConfig := &logger.Config{
		LogDir:    "logs",
		EnableLog: true,
	}

	log, err := logger.NewLogger(logConfig)
	if err != nil {
		panic("初始化日志失败: " + err.Error())
	}
	defer log.Close()

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	log.Info("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
