// Package service 用户服务层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"
	"Go-Management-System/common/util"
	"Go-Management-System/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ISysAdminService 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, Id int)
}

type SysAdminServiceImpl struct {
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}

// Login 用户登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验，根据结构体的validate标签校验属性值是否合法
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	// 验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	//校验密码
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	// 判断用户是否被禁用
	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]any{
		"token":    tokenString,
		"sysAdmin": sysAdmin,
	})

}

// CreateSysAdmin 新增用户
func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	ok := dao.CreateSysAdmin(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.USERNAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

// GetSysAdminInfo 根据id查询用户
func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysAdminInfo(Id))
}
