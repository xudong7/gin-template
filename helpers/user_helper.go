package helpers

import (
	"gin-second-fish/global"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

// update user avatar
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

// format users data for response
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

// format single user data for response
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
