// Package controller 部门controller层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysDeptService = service.SysDeptService()

var sysDept entity.SysDept

// GetSysDeptList 查询部门列表
// @Summary 查询部门列表接口
// @Produce json
// @Description 查询部门列表接口
// @Param deptName query string false "部门名称"
// @Param deptStatus query string false "部门状态"
// @Succss 200 {object} result.Result
// @router /api/dept/list [get]
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("deptName")
	DeptStatus := c.Query("deptStatus")
	sysDeptService.GetSysDeptList(c, DeptName, DeptStatus)
}

// CreateSysDept 新增部门
// @Summary 新增部门接口
// @Produce json
// @Description 新增部门接口
// @Param data body entity.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/add [post]
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// GetSysDeptById 根据id查询部门
// @Summary 根据id查询部门
// @Produce json
// @Description 根据id查询部门
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/dept/info [get]
func GetSysDeptById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	sysDept.ID = uint(id)
	service.SysDeptService().GetSysDeptById(c, sysDept)
}

// UpdateSysDept 修改部门
// @Summary 修改部门
// @Produce json
// @Description 修改部门
// @Param data body entity.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/update [put]
func UpdateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().UpdateSysDept(c, sysDept)
}
