package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} hello world
// @Router /example/hello [get]
func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello world")
}
