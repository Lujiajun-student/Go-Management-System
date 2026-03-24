// Package dao 角色相关dao层
package dao

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/util"
	. "Go-Management-System/pkg/db"
	"time"
)

// GetSysRoleByName 根据角色名称获取角色
func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// GetSysRoleByKey 根据角色权限字符串获取角色
func GetSysRoleByKey(roleKey string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// CreateSysRole 创建角色
func CreateSysRole(dto entity.AddSysRoleDto) bool {
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		return false
	}
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		return false
	}
	addSysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysRoleById 根据id查询角色
func GetSysRoleById(Id uint) (sysRole entity.SysRole) {
	Db.First(&sysRole, Id)
	return sysRole
}

// UpdateSysRole 修改角色
func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	return sysRole
}

// DeleteSysRoleById 删除角色
func DeleteSysRoleById(dto entity.SysRoleIdDto) {
	Db.Table("sys_role").Delete(&entity.SysRole{}, dto.Id)
	Db.Table("sys_role_menu").Where("role_id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
}

// UpdateSysRoleStatus 角色状态更新
func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatusDto) bool {
	var sysRole entity.SysRole
	Db.First(&sysRole, dto.Id)
	sysRole.Status = dto.Status
	tx := Db.Save(&sysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysRoleList 分页查询角色列表
func GetSysRoleList(PageNum, PageSize int, RoleName, status, BeginTime, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name like ?", "%"+RoleName+"%")
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if status != "" {
		curDb = curDb.Where("status = ?", status)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysRole)
	return sysRole, count
}

// QuerySysRoleVOList 角色下拉查询
func QuerySysRoleVOList() (sysRoleVO []entity.SysRoleVO) {
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVO)
	return sysRoleVO
}
