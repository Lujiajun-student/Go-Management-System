// Package entity 用户与角色关系模型
package entity

type SysAdminRole struct {
	RoleId  int  `gorm:"column:role_id;comment:'角色ID';NOT NULL" json:"roleId"`
	AdminId uint `gorm:"column:admin_id;comment:'用户ID';NOT NULL" json:"adminId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
