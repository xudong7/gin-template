package controllers

import (
	"gin-second-fish/utils"

	"github.com/gin-gonic/gin"
)

// / toggle favorite list
func ToggleFavorite(ctx *gin.Context) {
	itemId := ctx.Param("item_id")
	userId := ctx.Param("user_id")

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
	itemId := ctx.Param("item_id")
	userId := ctx.Param("user_id")

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

// get top favorite items
func GetTopFavoriteItems(ctx *gin.Context) {
	// limit := 5
	limit := ctx.Query("limit")
	items, err := utils.GetTopFavoriteItems(int64(utils.ToInt(limit)))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to get top favorite items",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get top favorite items success",
		"items":   items,
	})
}
