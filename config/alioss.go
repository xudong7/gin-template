package config

import (
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OSSClient *oss.Client
var OSS_BUCKET *oss.Bucket

func InitOSS() error {
	var err error
	OSSClient, err = oss.New(
		AppConfig.OSS.Endpoint,
		AppConfig.OSS.AccessKeyId,
		AppConfig.OSS.AccessKeySecret,
	)
	if err != nil {
		return err
	}

	// get bucket
	OSS_BUCKET, err = OSSClient.Bucket(AppConfig.OSS.BucketName)
	if err != nil {
		return err
	}
	return nil
}

// UploadFileToOSS
func UploadFileToOSS(file *multipart.FileHeader) (string, error) {
	// open file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// generate unique file name (timestamp + file extension)
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	objectKey := AppConfig.OSS.Directory + fileName

	// upload file to OSS
	err = OSS_BUCKET.PutObject(objectKey, src)
	if err != nil {
		return "", err
	}

	// generate file URL
	fileURL := AppConfig.OSS.URLPrefix + objectKey
	return fileURL, nil
}

// DeleteFileFromOSS
func DeleteFileFromOSS(objectKey string) error {
	return OSS_BUCKET.DeleteObject(objectKey)
}

// GetFileFromOSS
func GetFileFromOSS(objectKey string) (io.ReadCloser, error) {
	return OSS_BUCKET.GetObject(objectKey)
}
