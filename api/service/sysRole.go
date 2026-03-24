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
	DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto)
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto)
	GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string)
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

// DeleteSysRoleById 删除角色
func (s SysRoleServiceImpl) DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto) {
	dao.DeleteSysRoleById(dto)
	result.Success(c, true)
}

// UpdateSysRoleStatus 修改角色状态
func (s SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto) {
	ok := dao.UpdateSysRoleStatus(dto)
	if !ok {
		return
	}
	result.Success(c, true)
}

// GetSysRoleList 分页查询角色
func (s SysRoleServiceImpl) GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysRole, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
	result.Success(c, map[string]any{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysRole})
}
