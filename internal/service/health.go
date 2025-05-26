package service

import (
	"gin-template/pkg/utils"
)

type PingService struct{}

// NewPingService 创建新的PingService
func NewPingService() *PingService {
	return &PingService{}
}

// Ping 返回pong响应
func (s *PingService) Ping() string {
	timestamp := utils.GetCurrentTime()
	return "pong at " + timestamp
}

// HealthCheck 健康检查
func (s *PingService) HealthCheck() string {
	return "healthy"
}
