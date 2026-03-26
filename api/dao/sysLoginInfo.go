// Package dao 登陆日志dao层
package dao

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/util"
	. "Go-Management-System/pkg/db"
	"time"
)

// CreateSysLoginInfo 新增登录日志
func CreateSysLoginInfo(username, ipAddress, loginLocation, browser, os, message string, loginStatus int) {
	sysLoginInfo := entity.SysLoginInfo{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     util.HTime{Time: time.Now()},
	}
	Db.Save(&sysLoginInfo)
}

// GetSysLoginInfoList 分页获取登陆日志列表
func GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) (sysLoginInfo []entity.SysLoginInfo, count int64) {
	curDb := Db.Table("sys_login_info")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}
	if LoginStatus != "" {
		curDb = curDb.Where("login_status = ?", LoginStatus)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`login_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("`login_time` desc").Find(&sysLoginInfo)
	return sysLoginInfo, count
}

// DeleteSysLoginInfoById 根据id删除日志
func DeleteSysLoginInfoById(dto entity.SysLoginInfoIdDto) {
	Db.Delete(&entity.SysLoginInfo{}, dto.Id)
}

// BatchDeleteSysLoginInfo 批量删除日志
func BatchDeleteSysLoginInfo(dto entity.DelSysLoginInfoDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysLoginInfo{})
}

// CleanSysLoginInfo 清空登录日志
func CleanSysLoginInfo() {
	Db.Exec("TRUNCATE TABLE sys_login_info")
}
