package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func SelectUserById(id int) (models.User, error) {
	var user models.User

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return user, err
	}

	if err := global.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func SelectUserByIdString(id string) (models.User, error) {
	var user models.User

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return user, err
	}

	if err := global.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func SelectUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return user, err
	}

	if err := global.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func SelectAllUsers() ([]models.User, error) {
	var users []models.User

	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	if err := global.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func FormatUsersForResponse(users []models.User) []gin.H {
	var formattedUsers []gin.H

	for _, user := range users {
		formattedUsers = append(formattedUsers, gin.H{
			"id":            user.ID,
			"nickname":      user.Nickname,
			"username":      user.Username,
			"email":         user.Email,
			"phone":         user.Phone,
			"avatar":        user.Avatar,
			"status":        user.Status,
			"credit_score":  user.CreditScore,
			"last_login_at": user.LastLoginAt,
			// "create_at": user.CreateAt,
			// "update_at": user.UpdateAt,
		})
	}

	return formattedUsers
}

func FormatUserForResponse(user models.User) gin.H {
	return gin.H{
		"id":            user.ID,
		"nickname":      user.Nickname,
		"username":      user.Username,
		"email":         user.Email,
		"phone":         user.Phone,
		"avatar":        user.Avatar,
		"status":        user.Status,
		"credit_score":  user.CreditScore,
		"last_login_at": user.LastLoginAt,
		// "create_at": user.CreateAt,
		// "update_at": user.UpdateAt,
	}
}
