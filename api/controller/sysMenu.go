// Package controller 菜单controller层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"

	"github.com/gin-gonic/gin"
)

var sysMenu entity.SysMenu

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
