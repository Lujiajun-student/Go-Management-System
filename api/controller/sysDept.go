// Package controller 部门controller层
package controller

import (
	"Go-Management-System/api/service"

	"github.com/gin-gonic/gin"
)

var sysDeptService = service.SysDeptService()

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
