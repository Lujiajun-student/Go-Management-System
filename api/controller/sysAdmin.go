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

// UpdateSysAdminStatus 修改用户状态
// @Summary 修改用户状态
// @Produce json
// @Description 修改用户状态
// @param data body entity.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updateStatus [put]
func UpdateSysAdminStatus(c *gin.Context) {
	var dto entity.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminStatus(c, dto)
}

// ResetSysAdminPassword 重置密码
// @Summary 重置密码
// @Produce json
// @Description 重置密码
// @param data body entity.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePassword [put]
func ResetSysAdminPassword(c *gin.Context) {
	var dto entity.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().ResetSysAdminPassword(c, dto)
}

// GetSysAdminList 分页查询用户
// @Summary 分页查询用户
// @Produce json
// @Description 分页查询用户
// @param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名"
// @Param Status query string false "账号启用状态：1->启用，2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/admin/list [get]
func GetSysAdminList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysAdminService().GetSysAdminList(c, pageSize, pageNum, Username, Status, BeginTime, EndTime)
}

// UpdatePersonal 修改个人信息
// @Summary 修改个人信息
// @Produce json
// @Description 修改个人信息
// @param data body entity.UpdatePersonalDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonal [put]
func UpdatePersonal(c *gin.Context) {
	var dto entity.UpdatePersonalDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdatePersonal(c, dto)
}

// UpdatePersonalPassword 修改密码
// @Summary 修改密码
// @Produce json
// @Description 修改密码
// @param data body entity.UpdatePersonalPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonalPassword [put]
func UpdatePersonalPassword(c *gin.Context) {
	var dto entity.UpdatePersonalPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdatePersonalPassword(c, dto)
}
