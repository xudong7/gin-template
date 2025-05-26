package handler

import (
	"gin-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
	pingService *service.PingService
}

// NewPingHandler 创建新的PingHandler
func NewPingHandler() *PingHandler {
	return &PingHandler{
		pingService: service.NewPingService(),
	}
}

// Ping 处理ping请求
func (h *PingHandler) Ping(c *gin.Context) {
	response := h.pingService.Ping()
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

// Health 健康检查
func (h *PingHandler) Health(c *gin.Context) {
	status := h.pingService.HealthCheck()
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
