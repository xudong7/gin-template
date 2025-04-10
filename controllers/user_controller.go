package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"
	"gin-second-fish/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := global.Db.Find(&users).Error; err != nil {
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

	var user models.User
	if err := global.Db.Where("id = ?", id).First(&user).Error; err != nil {
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

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to auto migrate user",
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

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	existingUser, err := helpers.SelectUserByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updateData struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if updateData.Nickname != "" {
		existingUser.Nickname = updateData.Nickname
	}

	if updateData.Email != "" {
		existingUser.Email = updateData.Email
	}

	if updateData.Phone != "" {
		existingUser.Phone = updateData.Phone
	}

	if updateData.Avatar != "" {
		existingUser.Avatar = updateData.Avatar
	}

	if updateData.Password != "" {
		hashedPassword, err := utils.HashPassword(updateData.Password)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		existingUser.Password = hashedPassword
	}

	if err := global.Db.Save(&existingUser).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "update user",
		"user":    helpers.FormatUserForResponse(existingUser),
	})
}

func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := global.Db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete user",
	})
}

func DeleteUsers(ctx *gin.Context) {
	// delete all users
	if err := global.Db.Exec("DELETE FROM users").Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete all users",
	})
}
