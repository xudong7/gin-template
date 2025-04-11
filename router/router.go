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

	// v1 not use middleware
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/hello", controllers.Hello)
		}
		user := v1.Group("/user")
		{
			user.GET("/", controllers.GetUsers)
			user.GET("/:id", controllers.GetUserById)
			user.POST("/", controllers.InsertUser)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUserById)
			user.DELETE("/", controllers.DeleteUsers)
		}
		item := v1.Group("/item")
		{
			item.GET("/", controllers.GetItems)
			item.GET("/:id", controllers.GetItemById)
			item.POST("/", controllers.InsertItem)
			item.PUT("/:id", controllers.UpdateItem)
			item.DELETE("/:id", controllers.DeleteItemById)
			item.DELETE("/", controllers.DeleteItems)
		}
		order := v1.Group("/order")
		{
			order.GET("/", controllers.GetOrders)
			order.GET("/:id", controllers.GetOrderById)
			order.POST("/", controllers.InsertOrder)
			order.PUT("/:id", controllers.UpdateOrder)
			order.DELETE("/:id", controllers.DeleteOrderById)
			order.DELETE("/", controllers.DeleteOrders)
		}
		category := v1.Group("/category")
		{
			category.GET("/", controllers.GetCategories)
			category.GET("/:id", controllers.GetCategoryById)
			category.POST("/", controllers.InsertCategory)
			category.PUT("/:id", controllers.UpdateCategory)
			category.DELETE("/:id", controllers.DeleteCategoryById)
			category.DELETE("/", controllers.DeleteCategories)
		}
		upload := v1.Group("/upload")
		{
			upload.POST("/file", controllers.UploadFile)
			upload.POST("/avatar/:userId", controllers.UploadAvatar)
		}
	}

	// v2 use middleware
	v2 := r.Group("/api/v2")
	v2.Use(middlewares.AuthMiddleWare())
	{

	}

	return r
}
