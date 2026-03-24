// Package dao 部门dao层
package dao

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/util"
	. "Go-Management-System/pkg/db"
	"time"
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

// GetSysDeptByName 根据部门名称查询
func GetSysDeptByName(deptName string) (sysDept entity.SysDept) {
	Db.Where("dept_name = ?", deptName).First(&sysDept)
	return sysDept
}

// CreateSysDept 新增部门
func CreateSysDept(sysDept entity.SysDept) bool {
	// 查重
	sysDeptByName := GetSysDeptByName(sysDept.DeptName)
	if sysDeptByName.ID > 0 {
		return false
	}
	if sysDept.DeptType == 1 {
		sysDept := entity.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   0,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept)
		return true
	} else {
		sysDept := entity.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   sysDept.ParentId,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept)
		return true
	}
}

// GetSysDeptById 根据id查询部门
func GetSysDeptById(sysDept entity.SysDept) entity.SysDept {
	Db.Where("id = ?", sysDept.ID).First(&sysDept)
	return sysDept
}

// UpdateSysDept 修改部门
func UpdateSysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	Db.First(&sysDept, dept.ID)
	sysDept.ParentId = dept.ParentId
	sysDept.DeptName = dept.DeptName
	sysDept.DeptType = dept.DeptType
	sysDept.DeptStatus = dept.DeptStatus
	Db.Save(&sysDept)
	return sysDept
}

// GetSysAdminDept 查询部门是否有人
func GetSysAdminDept(id int) (sysAdmin entity.SysAdmin) {
	Db.Where("dept_id = ?", id).First(&sysAdmin)
	return sysAdmin
}

// DeleteSysDeptById 删除部门
func DeleteSysDeptById(dto entity.SysDeptIdDto) bool {
	sysAdmin := GetSysAdminDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysDept{})
	Db.Delete(&entity.SysDept{}, dto.Id)
	return true
}
