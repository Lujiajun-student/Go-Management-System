// Package controller 图片上传controller
package controller

import (
	"Go-Management-System/api/service"

	"github.com/gin-gonic/gin"
)

// Upload 单图片上传
// @Summary 单图片上传接口
// @Produce json
// @Description 单图片上传接口
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} result.Result
// @Router /api/upload [post]
func Upload(c *gin.Context) {
	service.UploadService().Upload(c)
}
