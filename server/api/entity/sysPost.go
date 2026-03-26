// Package entity 岗位相关模型
package entity

import "Go-Management-System/common/util"

type SysPost struct {
	ID         uint       `gorm:"column:id; comment:'主键';primaryKey;NOT NULL" json:"id"`
	PostCode   string     `gorm:"column:post_code;varchar(64);comment:'岗位编码;NOT NULL" json:"postCode"`
	PostName   string     `gorm:"column:post_name;varchar(50);comment:'岗位名称';NOT NULL" json:"postName"`
	PostStatus int        `gorm:"column:post_status;default:1;comment:'状态（1->正常 2->停用）';NOT NULL" json:"postStatus"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
	Remark     string     `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

// SysPostIdDto Id参数
type SysPostIdDto struct {
	Id uint `json:"id"`
}

func (SysPostIdDto) TableName() string {
	return "sys_post"
}

// DelSysPostDto 删除多个岗位
type DelSysPostDto struct {
	Ids []uint
}

func (DelSysPostDto) TableName() string {
	return "sys_post"
}

// UpdateSysPostStatusDto 修改状态参数
type UpdateSysPostStatusDto struct {
	Id         uint `json:"id"`
	PostStatus int  `json:"postStatus"`
}

func (UpdateSysPostStatusDto) TableName() string {
	return "sys_post"
}

// SysPostVO 返回给前端的岗位列表信息
type SysPostVO struct {
	Id       uint   `json:"id"`
	PostName string `json:"postName"`
}
