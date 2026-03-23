// Package dao 部门dao层
package dao

import (
	"Go-Management-System/api/entity"
	. "Go-Management-System/pkg/db"
)

// GetSysDeptList 查询部门列表
func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []entity.SysDept) {
	curDb := Db.Table("sys_dept")
	if DeptName != "" {
		curDb = curDb.Where("dept_name LIKE ?", "%"+DeptName+"%")
	}
	if DeptStatus != "" {
		curDb = curDb.Where("dept_status LIKE ?", "%"+DeptStatus+"%")
	}
	curDb.Find(&sysDept)
	return sysDept
}
