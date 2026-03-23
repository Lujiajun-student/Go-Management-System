// Package controller 岗位控制层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysPost entity.SysPost

// CreateSysPost 新增岗位
// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
func CreateSysPost(c *gin.Context) {
	// 从请求中获取JSON数据并绑定到sysPost结构体
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// GetSysPostList 根据条件分页查询岗位
// @Summary 分页查询岗位列表
// @Produce json
// @Description 分页查询岗位列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param postName query string false "岗位名称"
// @Param postStatus query string false "状态：1-> 启用，2->停用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
func GetSysPostList(c *gin.Context) {
	// 从请求中获取参数，int类型需要通过strconv.Atoi来转换
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	// string类型可直接通过Query获取
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
}

// GetSysPostById 根据id查询岗位
// @Summary 根据id查询岗位
// @Produce json
// @Description 根据id查询岗位
// @Param id query int true "ID"
// @Success 200 {object} result.Result
// @router /api/post/info [get]
func GetSysPostById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysPostService().GetSysPostById(c, entity.SysPost{ID: uint(Id)})
}

// UpdateSysPost 修改岗位
// @Summary 修改岗位接口
// @Producce json
// @Description 修改岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/update [put]
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}

// DeleteSysPostById 根据id删除岗位
// @Summary 根据id删除岗位
// @Produce json
// @Description 根据id删除岗位
// @Param data body entity.SysPostIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/delete [delete]
func DeleteSysPostById(c *gin.Context) {
	var dto entity.SysPostIdDto
	_ = c.ShouldBindJSON(&dto)
	service.SysPostService().DeleteSysPostById(c, dto)
}

// BatchDeleteSysPost 批量删除岗位
// @Summary 批量删除岗位
// @Produce json
// @Description 批量删除岗位
// @Param data body entity.DelSysPostDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/batch/delete [delete]
func BatchDeleteSysPost(c *gin.Context) {
	var dto entity.DelSysPostDto
	_ = c.ShouldBindJSON(&dto)
	service.SysPostService().BatchDeleteSysPost(c, dto)
}
