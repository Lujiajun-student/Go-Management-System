// Package entity 用户与角色关系模型
package entity

type SysAdminRole struct {
	AdminId uint `gorm:"column:admin_id;comment:'用户ID';NOT NULL" json:"adminId"`
	RoleId  uint `gorm:"column:role_id;comment:'角色ID';NOT NULL" json:"roleId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
