// Package controller 角色相关controller层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var addSysRole entity.AddSysRoleDto

var updateSysRole entity.UpdateSysRoleDto

var sysRoleIdDto entity.SysRoleIdDto

// CreateSysRole 创建角色
// @Summary 新增角色接口
// @Produce json
// @Description 新增角色接口
// @Param data body entity.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/add [post]
func CreateSysRole(c *gin.Context) {
	_ = c.BindJSON(&addSysRole)
	service.SysRoleService().CreateSysRole(c, addSysRole)
}

// GetSysRoleById 根据id查询角色
// @Summary 根据id查询角色
// @Produce json
// @Description 根据id查询角色
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/role/info [get]
func GetSysRoleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, uint(id))
}

// UpdateSysRole 修改角色
// @Summary 修改角色
// @Produce json
// @Description 修改角色
// @Param data body entity.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/update [put]
func UpdateSysRole(c *gin.Context) {
	_ = c.BindJSON(&updateSysRole)
	service.SysRoleService().UpdateSysRole(c, updateSysRole)
}

// DeleteSysRoleById 删除角色
// @Summary 删除角色
// @Produce json
// @Description 删除角色
// @Param data body entity.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/delete [delete]
func DeleteSysRoleById(c *gin.Context) {
	_ = c.BindJSON(&sysRoleIdDto)
	service.SysRoleService().DeleteSysRoleById(c, sysRoleIdDto)
}
