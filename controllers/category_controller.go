package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	var categories []models.Category
	if err := global.Db.Find(&categories).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to get categories",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":    "get categories success",
		"categories": categories,
	})
}

func GetCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")

	var category models.Category
	if err := global.Db.Where("id = ?", id).First(&category).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "category not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "get category success",
		"category": category,
	})
}

func InsertCategory(ctx *gin.Context) {
	var requestData map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fields := []string{"parent_id"}
	helpers.ConvertStringIdsToInt(requestData, global.CategoryFields...)
	var category models.Category
	if err := helpers.MapToStruct(requestData, &category); err != nil {
		ctx.JSON(400, gin.H{
			"error": "data type error: " + err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&models.Category{}); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to migrate category",
		})
		return
	}

	if err := global.Db.Create(&category).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to create category",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "create category success",
		"category": category,
	})
}

func UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	existingCategory, err := helpers.SelectCategoryByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "category not found",
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

	// fields := []string{"parent_id"}
	helpers.ConvertStringIdsToInt(requestData, global.CategoryFields...)
	var category models.Category
	if err := helpers.MapToStruct(requestData, &category); err != nil {
		ctx.JSON(400, gin.H{
			"error": "data type error: " + err.Error(),
		})
		return
	}

	if err := global.Db.Model(&existingCategory).Updates(category).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to update category",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "update category success",
		"category": category,
	})
}

func DeleteCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := global.Db.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to delete category",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete category success",
	})
}

func DeleteCategories(ctx *gin.Context) {
	// delete all categories
	if err := global.Db.Exec("DELETE FROM categories").Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to delete categories",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete categories success",
	})
}
