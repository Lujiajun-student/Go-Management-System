// Package entity 角色相关实体类
package entity

import "Go-Management-System/common/util"

// SysRole 角色模型
type SysRole struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	RoleName    string     `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`
	RoleKey     string     `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`
	Status      int        `gorm:"column:status;default:1;comment:'账号启用状态：1->启用；2->禁用';NOT NULL" json:"status"`
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

// AddSysRoleDto 新增角色所需参数
type AddSysRoleDto struct {
	RoleName    string
	RoleKey     string
	Status      int
	Description string
}

// UpdateSysRoleDto 修改所需参数
type UpdateSysRoleDto struct {
	Id          uint
	RoleName    string
	RoleKey     string
	Status      int
	Description string
}
