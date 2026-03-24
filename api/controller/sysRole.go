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

// UpdateSysRoleStatus 修改角色状态
// @Summary 修改角色状态
// @Produce json
// @Description 修改角色状态
// @Param data body entity.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/updateStatus [put]
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// GetSysRoleList 分页查询角色列表
// @Summary 分页查询角色列表
// @Produce json
// @Description 分页查询角色列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param roleName query string false "角色名称"
// @Param status query int false "状态：1->启用；2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/role/list [get]
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
}

// QuerySysRoleVOList 查询角色下拉列表
// @Summary 查询角色下拉列表
// @Produce json
// @Description 查询角色下拉列表
// @Success 200 {object} result.Result
// @router /api/role/vo/list [get]
func QuerySysRoleVOList(c *gin.Context) {
	service.SysRoleService().QuerySysRoleVOList(c)
}
