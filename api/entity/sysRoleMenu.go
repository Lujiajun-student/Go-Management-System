// Package entity 角色菜单模型
package entity

// SysRoleMenu 角色与菜单关系模型
type SysRoleMenu struct {
	RoleId uint `gorm:"column:role_id;comment:'角色ID';NOT NULL" json:"roleId"`
	MenuId uint `gorm:"column:menu_id;comment:'用户id';NOT NULL" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
