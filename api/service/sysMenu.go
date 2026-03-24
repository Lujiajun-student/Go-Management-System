// Package service 菜单service层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysMenuService interface {
	CreateSysMenu(c *gin.Context, SysMenu entity.SysMenu)
	QuerySysMenuVOList(c *gin.Context)
}

type SysMenuServiceImpl struct {
}

var sysMenuService = SysMenuServiceImpl{}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}

// CreateSysMenu 创建菜单
func (s SysMenuServiceImpl) CreateSysMenu(c *gin.Context, SysMenu entity.SysMenu) {
	ok := dao.CreateSysMenu(SysMenu)
	if !ok {
		result.Failed(c, int(result.ApiCode.MENUISEXIST), result.ApiCode.GetMessage(result.ApiCode.MENUISEXIST))
		return
	}
	result.Success(c, true)
}

// QuerySysMenuVOList 查询菜单列表
func (s SysMenuServiceImpl) QuerySysMenuVOList(c *gin.Context) {
	result.Success(c, dao.QuerySysMenuVOList())
}
