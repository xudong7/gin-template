package router

import (
	"gin-second-fish/config"
	"gin-second-fish/controllers"
	"gin-second-fish/docs"
	"gin-second-fish/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Swagger setting
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/hello", controllers.Hello)
		}
		need_auth := v1.Group("/need_auth")
		need_auth.Use(middlewares.AuthMiddleWare())
		{
			// apis which need auth
		}
	}

	return r
}
