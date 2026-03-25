// Package service 图片上传service
package service

import (
	"Go-Management-System/common/config"
	"Go-Management-System/common/result"
	"Go-Management-System/common/util"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUploadService interface {
	Upload(c *gin.Context)
}

type UploadServiceImpl struct {
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}

// Upload 图片上传
func (u *UploadServiceImpl) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FileUploadError), result.ApiCode.GetMessage(result.ApiCode.FileUploadError))
		return
	}
	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s", config.Config.ImageSettings.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	err = util.CreateDir(filePath)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FileUploadError), result.ApiCode.GetMessage(result.ApiCode.FileUploadError))
		return
	}
	fullPath := filePath + "/" + fileName
	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FileUploadError), result.ApiCode.GetMessage(result.ApiCode.FileUploadError))
		return
	}
	result.Success(c, config.Config.ImageSettings.ImageHost+fullPath)
}
