// Package controller 菜单controller层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysMenu entity.SysMenu

var sysMenuVO entity.SysMenuVO

var sysMenuIdDto entity.SysMenuIdDto

// CreateSysMenu 创建菜单
// @Summary 新增菜单接口
// @Producce json
// @Description 新增菜单接口
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/add [post]
func CreateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().CreateSysMenu(c, sysMenu)
}

// QuerySysMenuVOList 查询菜单列表
// @Summary 查询菜单列表
// @Producce json
// @Description 查询菜单列表
// @Success 200 {object} result.Result
// @router /api/menu/vo/list [get]
func QuerySysMenuVOList(c *gin.Context) {
	_ = c.BindJSON(&sysMenuVO)
	service.SysMenuService().QuerySysMenuVOList(c)
}

// GetSysMenuById 根据id查询菜单
// @Summary 根据id查询菜单
// @Producce json
// @Description 根据id查询菜单
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/menu/info [get]
func GetSysMenuById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysMenuService().GetSysMenuById(c, id)
}

// UpdateSysMenu 修改菜单
// @Summary 修改菜单
// @Producce json
// @Description 修改菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/update [put]
func UpdateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().UpdateSysMenu(c, sysMenu)
}

// DeleteSysMenuById 删除菜单
// @Summary 删除菜单
// @Producce json
// @Description 删除菜单
// @Param data body entity.SysMenuIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu/delete [delete]
func DeleteSysMenuById(c *gin.Context) {
	_ = c.BindJSON(&sysMenuIdDto)
	service.SysMenuService().DeleteSysMenuById(c, sysMenuIdDto)
}

// GetSysMenuList 查询菜单列表
// @Summary 查询菜单列表
// @Producce json
// @Description 查询菜单列表
// @Param MenuName query string false "MenuName"
// @Param MenuStatus query string false "MenuStatus"
// @Success 200 {object} result.Result
// @router /api/menu/list [get]
func GetSysMenuList(c *gin.Context) {
	MenuName := c.Query("MenuName")
	MenuStatus := c.Query("MenuStatus")
	service.SysMenuService().GetSysMenuList(c, MenuName, MenuStatus)
}
