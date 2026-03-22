package dao

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/util"
	"Go-Management-System/pkg/db"
	"time"
)

// GetSysPostByCode 根据postCode查询岗位
func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	db.Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

// GetSysPostByName 根据岗位名称查询
func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	db.Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

// CreateSysPost 新增岗位
func CreateSysPost(sysPost entity.SysPost) bool {
	// 查看postCode是否重复
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}
	// 查看postName是否重复
	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}
	// 创建新增岗位实例
	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     sysPost.Remark,
	}
	// 保存到数据库的sys_post表中
	tx := db.Db.Save(&addSysPost)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
