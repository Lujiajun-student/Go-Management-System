// Package service 角色相关service层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
	GetSysRoleById(c *gin.Context, id uint)
	UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto)
}

type SysRoleServiceImpl struct {
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}

// CreateSysRole 创建角色
func (s SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto) {
	ok := dao.CreateSysRole(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.ROLENAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.ROLENAMEALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

// GetSysRoleById 根据id查询角色
func (s SysRoleServiceImpl) GetSysRoleById(c *gin.Context, id uint) {
	result.Success(c, dao.GetSysRoleById(id))
}

// UpdateSysRole 修改角色
func (s SysRoleServiceImpl) UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto) {
	result.Success(c, dao.UpdateSysRole(dto))
}
