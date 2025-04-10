package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func GetItems(ctx *gin.Context) {
	items, err := helpers.SelectAllItems()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to get items",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get items success",
		"items":   items,
	})
}

func GetItemById(ctx *gin.Context) {
	id := ctx.Param("id")

	item, err := helpers.SelectItemByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "item not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get item success",
		"item":    item,
	})
}

func InsertItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := helpers.SelectItemByIdString(string(rune(item.ID)))
	if err == nil {
		ctx.JSON(400, gin.H{
			"error": "item already exists",
		})
		return
	}

	if err := global.Db.Create(&item).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to create item",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "item created successfully",
		"item":    item,
	})
}

func UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")
	exsitingItem, err := helpers.SelectItemByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "item not found",
		})
		return
	}

	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Model(&exsitingItem).Updates(item).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to update item",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "item updated successfully",
		"item":    item,
	})
}

func DeleteItemById(ctx *gin.Context) {
	id := ctx.Param("id")

	exsitingItem, err := helpers.SelectItemByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "item not found",
		})
		return
	}

	if err := global.Db.Delete(&exsitingItem).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to delete item",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "item deleted successfully",
	})
}

func DeleteItems(ctx *gin.Context) {
	// delete all items
	if err := global.Db.Exec("DELETE FROM items").Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to delete items",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "all items deleted successfully",
	})
}
