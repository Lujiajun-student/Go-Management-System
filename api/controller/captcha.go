package controller

import (
	"Go-Management-System/api/service"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

// Captcha 从Service获取验证码图片和ID
// @Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	// 获取验证码图片和ID
	id, base64Image := service.CreateCaptcha()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
