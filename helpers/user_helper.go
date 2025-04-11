package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func UpdateUserAvatar(userId, fileUrl string) error {
	// update user.avatar
	user, err := SelectUserByIdString(userId)
	if err != nil {
		return err
	}

	if err := global.Db.Model(&user).Update("avatar", fileUrl).Error; err != nil {
		return err
	}

	return nil
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
			"created_at":    user.CreatedAt,
			"updated_at":    user.UpdatedAt,
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
		"created_at":    user.CreatedAt,
		"updated_at":    user.UpdatedAt,
	}
}
