package router

import (
	"gin-second-fish/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load HTML templates from the "templates" directory
	// r.LoadHTMLGlob("templates/*")

	// Serve static files from the "static" directory
	// r.Static("/static", "./static")

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.Front.Url},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == config.AppConfig.Front.Url
		},
		MaxAge: 12 * time.Hour,
	}))

	// Define your routes here
	// auth := r.Group("/api/auth")
	// {

	// }

	// api := r.Group("/api")
	// api.Use(middleware.AuthMiddleware())
	// {
	// api.GET("/user/:id", controller.GetUser)
	// }

	return r
}
