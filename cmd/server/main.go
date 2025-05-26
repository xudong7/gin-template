package main

import (
	"gin-template/internal/router"
	"log"
)

func main() {
	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
