package controllers

import (
	"gin-second-fish/utils"

	"github.com/gin-gonic/gin"
)

// / toggle favorite list
// / if favorite is true, add item to favorite list
// / if favorite is false, remove item from favorite list
// / change favorite counts
func ToggleFavorite(ctx *gin.Context) {
	itemId := ctx.Param("item_id") // get item id from context
	userId := ctx.Param("user_id") // get user id from context

	// toggle favorite
	if err := utils.SetFavorite(userId, itemId); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to toggle favorite",
		})
		return
	}
}

// check if user in item favorite list
func IsFavorite(ctx *gin.Context) {
	itemId := ctx.Param("item_id") // get item id from context
	userId := ctx.Param("user_id") // get user id from context

	isMember, err := utils.IsFavorite(userId, itemId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to check favorite",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":     "check favorite success",
		"is_favorite": isMember,
	})
}
