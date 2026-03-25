// Package controller 用户控制层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login 登录
// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}

// CreateSysAdmin 创建用户
// @Summary 创建用户接口
// @Produce json
// @Description 创建用户接口
// @param data body entity.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/add [post]
func CreateSysAdmin(c *gin.Context) {
	var addSysAdminDto entity.AddSysAdminDto
	_ = c.BindJSON(&addSysAdminDto)
	service.SysAdminService().CreateSysAdmin(c, addSysAdminDto)
}

// GetSysAdminInfo 根据id查询用户
// @Summary 根据id查询用户
// @Produce json
// @Description 根据id查询用户
// @param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/admin/info [get]
func GetSysAdminInfo(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysAdminService().GetSysAdminInfo(c, Id)
}

// UpdateSysAdmin 修改用户
// @Summary 修改用户
// @Produce json
// @Description 修改用户
// @param data body entity.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/update [put]
func UpdateSysAdmin(c *gin.Context) {
	var dto entity.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdmin(c, dto)
}

// DeleteSysAdminById 根据id删除用户
// @Summary 根据id删除用户
// @Produce json
// @Description 根据id删除用户
// @param data body entity.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/delete [delete]
func DeleteSysAdminById(c *gin.Context) {
	var dto entity.SysAdminIdDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().DeleteSysAdminById(c, dto)
}
