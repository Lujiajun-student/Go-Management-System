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
	GetSysMenuById(c *gin.Context, id int)
	UpdateSysMenu(c *gin.Context, menu entity.SysMenu)
	DeleteSysMenuById(c *gin.Context, dto entity.SysMenuIdDto)
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

// GetSysMenuById 根据id查询菜单
func (s SysMenuServiceImpl) GetSysMenuById(c *gin.Context, id int) {
	result.Success(c, dao.GetSysMenuById(id))
}

// UpdateSysMenu 更新菜单
func (s SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, menu entity.SysMenu) {
	result.Success(c, dao.UpdateSysMenu(menu))
}

// DeleteSysMenuById 删除菜单
func (s SysMenuServiceImpl) DeleteSysMenuById(c *gin.Context, dto entity.SysMenuIdDto) {
	ok := dao.DeleteSysMenuById(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.DELSYSMENUFAILED), result.ApiCode.GetMessage(result.ApiCode.DELSYSMENUFAILED))
		return
	}
	result.Success(c, true)
}
