package router

import (
	"gin-template/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 创建处理器实例
	pingHandler := handler.NewPingHandler()

	// 注册路由
	r.GET("/ping", pingHandler.Ping)
	r.GET("/health", pingHandler.Health)

	return r
}
