// Package dao 用户数据层
package dao

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/util"
	. "Go-Management-System/pkg/db"
	"time"
)

// SysAdminDetail 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// GetSysAdminByUsername 根据用户名查询用户
func GetSysAdminByUsername(username string) (sysAdmin entity.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// CreateSysAdmin 新增用户
func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	sysAdminByUsername := GetSysAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		return false
	}
	sysAdmin := entity.SysAdmin{
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   util.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&sysAdmin)
	sysAdminExist := GetSysAdminByUsername(dto.Username)
	var e entity.SysAdminRole
	e.AdminId = sysAdminExist.ID
	e.RoleId = sysAdminExist.ID
	Db.Create(&e)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysAdminInfo 查询用户详情
func GetSysAdminInfo(Id int) (sysAdminInfo entity.SysAdminInfo) {
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, Id)
	return sysAdminInfo
}
