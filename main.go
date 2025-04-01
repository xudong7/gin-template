package main

import (
	"fmt"
	"gin-second-fish/config"
	"gin-second-fish/router"
)

func main() {
	config.InitConfig()
	fmt.Println("*******    NAME    ******* : " + config.AppConfig.App.Name)
	fmt.Println("*******    PORT    ******* : " + config.AppConfig.App.Port)

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))

}
