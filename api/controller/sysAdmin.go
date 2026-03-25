// Package controller 用户控制层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"

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
