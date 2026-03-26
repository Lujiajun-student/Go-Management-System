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
	GetSysDeptById(c *gin.Context, sysDept entity.SysDept)
	UpdateSysDept(c *gin.Context, sysDept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto)
	QuerySysDeptVOList(c *gin.Context)
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

// GetSysDeptById 根据id查询部门
func (s SysDeptServiceImpl) GetSysDeptById(c *gin.Context, sysDept entity.SysDept) {
	result.Success(c, dao.GetSysDeptById(sysDept))
}

// UpdateSysDept 修改部门
func (s SysDeptServiceImpl) UpdateSysDept(c *gin.Context, sysDept entity.SysDept) {
	dao.UpdateSysDept(sysDept)
	result.Success(c, sysDept)
}

// DeleteSysDeptById 删除部门
func (s SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto) {
	ok := dao.DeleteSysDeptById(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.DEPTISDISTRIBUTE), result.ApiCode.GetMessage(result.ApiCode.DEPTISDISTRIBUTE))
		return
	}
	result.Success(c, true)
}

// QuerySysDeptVOList 查询部门列表
func (s SysDeptServiceImpl) QuerySysDeptVOList(c *gin.Context) {
	result.Success(c, dao.QuerySysDeptVOList())
}
