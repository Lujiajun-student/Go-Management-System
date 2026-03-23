// Package service 部门service层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName, DeptStatus string)
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
}

type SysDeptServiceImpl struct {
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}

// GetSysDeptList 部门列表查询
func (s SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName, DeptStatus string) {
	result.Success(c, dao.GetSysDeptList(DeptName, DeptStatus))
}

// CreateSysDept 新增部门
func (s SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept entity.SysDept) {
	ok := dao.CreateSysDept(sysDept)
	if !ok {
		result.Failed(c, int(result.ApiCode.DEPTISEXIST), result.ApiCode.GetMessage(result.ApiCode.DEPTISEXIST))
		return
	}
	result.Success(c, true)
}
