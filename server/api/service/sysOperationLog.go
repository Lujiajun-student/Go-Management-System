// Package service 操作日志service
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

type ISysOperationLogService interface {
	GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int)
	DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto)
	BatchDeleteSysOperationLog(c *gin.Context, dto entity.BatchDeleteSysOperationLogDto)
	CleanSysOperationLog(c *gin.Context)
}
type SysOperationLogServiceImpl struct{}

var sysOperationLogService = SysOperationLogServiceImpl{}

func SysOperationLogService() ISysOperationLogService {
	return &sysOperationLogService
}

// GetSysOperationLogList 分页查询操作日志
func (s SysOperationLogServiceImpl) GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysOperationLogList, count := dao.GetSysOperationLogList(Username, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]any{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysOperationLogList})
}

// DeleteSysOperationLogById 根据id删除操作日志
func (s SysOperationLogServiceImpl) DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto) {
	dao.DeleteSysOperationLogById(dto)
	result.Success(c, true)
}

// BatchDeleteSysOperationLog 批量删除操作日志
func (s SysOperationLogServiceImpl) BatchDeleteSysOperationLog(c *gin.Context, dto entity.BatchDeleteSysOperationLogDto) {
	dao.BatchDeleteSysOperationLog(dto)
	result.Success(c, true)
}

// CleanSysOperationLog 清空操作日志
func (s SysOperationLogServiceImpl) CleanSysOperationLog(c *gin.Context) {
	dao.CleanSysOperationLog()
	result.Success(c, true)
}
