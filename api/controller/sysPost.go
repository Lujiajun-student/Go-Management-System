// Package controller 岗位控制层
package controller

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/api/service"

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
