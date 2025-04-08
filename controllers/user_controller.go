package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"
	"gin-second-fish/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	users, err := helpers.SelectAllUsers()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to get users",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get users success",
		"users":   helpers.FormatUsersForResponse(users),
	})
}

func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := helpers.SelectUserByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get user success",
		"user":    helpers.FormatUserForResponse(user),
	})
}

func InsertUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if the username already exists
	_, err := helpers.SelectUserByUsername(user.Username)
	if err == nil {
		ctx.JSON(400, gin.H{
			"error": "username already exists",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = hashedPassword
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, helpers.FormatUserForResponse(user))
}
