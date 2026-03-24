// Package entity 菜单模型
package entity

import "Go-Management-System/common/util"

// SysMenu 菜单模型
type SysMenu struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	ParentId   uint       `gorm:"column:parent_id;comment:'父菜单id'" json:"parentId"`
	MenuName   string     `gorm:"column:menu_name;varchar(100);comment:'菜单名称'" json:"menuName"`
	Icon       string     `gorm:"column:icon;varchar(100);comment:'菜单图标'" json:"icon"`
	Value      string     `gorm:"column:value;varchar(100);comment:'权限值'" json:"value"`
	MenuType   uint       `gorm:"column:menu_type;comment:'菜单类型：1->目录；2->菜单；3->按钮'" json:"menuType"`
	Url        string     `gorm:"column:url;varchar(100);comment:'菜单URL'" json:"url"`
	MenuStatus uint       `gorm:"column:menu_status;comment:'启动状态：1->启用；2->禁用'" json:"menuStatus"`
	Sort       uint       `gorm:"column:sort;comment:'排序'" json:"sort"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间'" json:"createTime"`
	Children   []SysMenu  `gorm:"-" json:"children"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
