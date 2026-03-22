// Package service 岗位服务层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
}

type SysPostServiceImpl struct {
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}

// CreateSysPost 新增岗位
func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	ok := dao.CreateSysPost(sysPost)
	if !ok {
		result.Failed(c, int(result.ApiCode.POSTALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}
