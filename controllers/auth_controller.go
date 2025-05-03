package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"
	"gin-second-fish/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dbUser models.User
	if err := global.Db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		ctx.JSON(401, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	if !utils.ComparePassword(user.Password, dbUser.Password) {
		ctx.JSON(401, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	// userId := fmt.Sprintf("%d", dbUser.ID)
	userId := utils.ToString(dbUser.ID)

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	if err := utils.SetToken(userId, token); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to set token to redis",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login success",
		"token":   token,
	})
}

func Register(ctx *gin.Context) {
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
			"error": "failed to hash password",
		})
		return
	}
	user.Password = hashedPassword

	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to create user",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "register success",
	})
}
