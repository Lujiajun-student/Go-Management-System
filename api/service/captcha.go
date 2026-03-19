package service

import (
	"Go-Management-System/common/util"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 验证码

var store = util.RedisStore{}

// CreateCaptcha 生成验证码
func CreateCaptcha() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码图片信息
	captchaConfig := base64Captcha.DriverString{
		Height:          69,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 生成并返回结果
	lid, lb64s, _, _ := captcha.Generate()
	return lid, lb64s

}
