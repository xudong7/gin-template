package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetItems(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "get items",
	})
}

func GetItemById(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "get item",
	})
}

func InsertItem(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "insert item",
	})
}

func UpdateItem(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "update item",
	})
}
