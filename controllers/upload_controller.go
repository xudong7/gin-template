package controllers

import (
	"gin-second-fish/config"
	"gin-second-fish/helpers"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 允许的文件类型
var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".pdf":  true,
	".doc":  true,
	".docx": true,
}

// maxFileSize 10MB
const maxFileSize = 10 << 20

// UploadFile
func UploadFile(ctx *gin.Context) {
	// get file from form data
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "No file uploaded",
		})
		return
	}

	// check file size
	if file.Size > maxFileSize {
		ctx.JSON(400, gin.H{
			"error": "File too large (max 10MB)",
		})
		return
	}

	// check file type
	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		ctx.JSON(400, gin.H{
			"error": "File type not allowed",
		})
		return
	}

	// upload to OSS
	fileURL, err := config.UploadFileToOSS(file)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to upload file: " + err.Error(),
		})
		return
	}

	// return file URL
	ctx.JSON(200, gin.H{
		"message":  "File uploaded successfully",
		"file_url": fileURL,
	})
}

// UploadAvatar
func UploadAvatar(ctx *gin.Context) {
	// get user ID from URL parameter
	userId := ctx.Param("userId")
	if userId == "" {
		ctx.JSON(400, gin.H{
			"error": "User ID is required",
		})
		return
	}

	// get file from form data
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "No avatar uploaded",
		})
		return
	}

	// check file size
	if file.Size > 2<<20 { // 2MB
		ctx.JSON(400, gin.H{
			"error": "Avatar too large (max 2MB)",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		ctx.JSON(400, gin.H{
			"error": "Only JPG, JPEG and PNG formats are allowed for avatars",
		})
		return
	}

	// upload to OSS
	fileURL, err := config.UploadFileToOSS(file)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to upload avatar: " + err.Error(),
		})
		return
	}

	// update url data in database
	if err := helpers.UpdateUserAvatar(userId, fileURL); err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to update user avatar: " + err.Error(),
		})
		return
	}

	// return url
	ctx.JSON(200, gin.H{
		"message":    "Avatar uploaded successfully",
		"avatar_url": fileURL,
	})
}
