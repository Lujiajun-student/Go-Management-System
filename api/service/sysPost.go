// Package service 岗位服务层
package service

import (
	"Go-Management-System/api/dao"
	"Go-Management-System/api/entity"
	"Go-Management-System/common/result"

	"github.com/gin-gonic/gin"
)

// ISysPostService 岗位相关接口
type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	GetSysPostById(c *gin.Context, post entity.SysPost)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)
}

// SysPostServiceImpl 岗位的实现类
type SysPostServiceImpl struct {
}

// sysPostService 实现类实例
var sysPostService = SysPostServiceImpl{}

// SysPostService 工厂函数，向外提供service层实例
func SysPostService() ISysPostService {
	return &sysPostService
}

// CreateSysPost 新增岗位
func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	ok := dao.CreateSysPost(sysPost)
	if !ok {
		result.Failed(c, int(result.ApiCode.POSTALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

// GetSysPostList 分页查询岗位列表
func (s SysPostServiceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {
	// 未设置页面大小，给出默认值
	if PageSize < 1 {
		PageSize = 10
	}
	// 未设置页数，给出首页
	if PageNum < 1 {
		PageNum = 1
	}
	// 调用dao层方法获取特定的岗位列表和总数
	sysPost, count := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	// 返回结果
	result.Success(c, map[string]any{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysPost})
}

// GetSysPostById 根据id查询岗位
func (s SysPostServiceImpl) GetSysPostById(c *gin.Context, post entity.SysPost) {
	result.Success(c, dao.GetSysPostById(post))
}

// UpdateSysPost 修改岗位
func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, post entity.SysPost) {
	result.Success(c, dao.UpdateSysPost(post))
}

// DeleteSysPostById 根据id删除岗位
func (s SysPostServiceImpl) DeleteSysPostById(c *gin.Context, post entity.SysPostIdDto) {
	dao.DeleteSysPostById(post)
	result.Success(c, true)
}

// BatchDeleteSysPost 批量删除岗位
func (s SysPostServiceImpl) BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto) {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}
