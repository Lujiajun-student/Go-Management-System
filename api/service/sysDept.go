// Package service 部门service层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName, DeptStatus string)
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
