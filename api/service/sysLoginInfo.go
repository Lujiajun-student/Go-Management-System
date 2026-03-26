// Package service 登陆日志service
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysLoginInfoService interface {
	GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, beginTime, EndTime string, PageSize, PageNum int)
}

type SysLoginInfoServiceImpl struct {
}

var sysLoginInfoService = SysLoginInfoServiceImpl{}

func SysLoginInfoService() ISysLoginInfoService {
	return &sysLoginInfoService
}

// GetSysLoginInfoList 分页获取登录日志
func (s SysLoginInfoServiceImpl) GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysLoginInfo, count := dao.GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]any{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysLoginInfo})
}
