// Package entity 部门相关模型
package entity

import "Go-Management-System/common/util"

// SysDept 部门模型
type SysDept struct {
	ID         uint       `gorm:"column:id;comment:'主键';primary_key;NOT NULL" json:"id"`
	ParentId   uint       `gorm:"column:parent_id;comment:'父id';NOT NULL" json:"parentId"`
	DeptType   uint       `gorm:"column:dept_type;comment:'部门类型（1->公司，2->中心，3->部门）';NOT NULL" json:"deptType"`
	DeptName   string     `gorm:"column:dept_name;varchar(30);comment:'部门名称';NOT NULL" json:"deptName"`
	DeptStatus int        `gorm:"column:dept_status;default:1;comment:'部门状态（1->正常，2->停用）'" json:"deptStatus"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
	Children   []SysDept  `gorm:"-" json:"children"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}

// SysDeptIdDto 接收id参数执行删除
type SysDeptIdDto struct {
	Id int `json:"id"`
}
