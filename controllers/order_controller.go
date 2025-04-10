package controllers

import (
	"gin-second-fish/global"
	"gin-second-fish/helpers"
	"gin-second-fish/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
	var orders []models.Order
	if err := global.Db.Find(&orders).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to get orders",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get orders success",
		"orders":  orders,
	})
}

func GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")

	var order models.Order
	if err := global.Db.Where("id = ?", id).First(&order).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "order not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get order success",
		"order":   order,
	})
}

// bind user and item -> create order -> return order
func InsertOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&models.Order{}); err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to auto migrate order",
		})
		return
	}

	if err := global.Db.Create(&order).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to create order",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "create order success",
		"order":   order,
	})
}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	existingOrder, err := helpers.SelectOrderByIdString(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "order not found",
		})
		return
	}

	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Model(&existingOrder).Updates(order).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to update order",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "update order success",
		"order":   order,
	})
}

func DeleteOrders(ctx *gin.Context) {
	// delete all orders
	if err := global.Db.Exec("DELETE FROM orders").Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to delete orders",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete orders success",
	})
}

func DeleteOrderById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := global.Db.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "order not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "delete order success",
	})
}
