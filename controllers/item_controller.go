package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func GetItems(ctx *gin.Context) {
	var items []models.Item

	if err := global.Db.Find(&items).Error; err != nil {
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

	var item models.Item
	if err := global.Db.Where("id = ?", id).First(&item).Error; err != nil {
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
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fields := []string{"seller_id", "category_id"}
	helpers.ConvertStringIdsToInt(requestData, global.ItemFields...)
	var item models.Item
	if err := helpers.MapToStruct(requestData, &item); err != nil {
		ctx.JSON(400, gin.H{
			"error": "data type error: " + err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&models.Item{}); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to auto migrate item",
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
	existingItem, err := helpers.SelectItemByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "item not found",
		})
		return
	}

	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fields := []string{"seller_id", "category_id"}
	helpers.ConvertStringIdsToInt(requestData, global.ItemFields...)
	var item models.Item
	if err := helpers.MapToStruct(requestData, &item); err != nil {
		ctx.JSON(400, gin.H{
			"error": "data type error: " + err.Error(),
		})
		return
	}

	if err := global.Db.Model(&existingItem).Updates(item).Error; err != nil {
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

	if err := global.Db.Where("id = ?", id).Delete(&models.Item{}).Error; err != nil {
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

	// delete data from database need to add always true condition like "1=1"
	// if err := global.Db.Unscoped().Where("1=1").Delete(&models.Item{}).Error; err != nil {
	// 	ctx.JSON(500, gin.H{
	// 		"error": "failed to permanently delete items",
	// 	})
	// 	return
	// }

	ctx.JSON(200, gin.H{
		"message": "all items deleted successfully",
	})
}
