// Package controller 登陆日志controller
package controller

import (
	"Go-Management-System/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSysLoginInfo 分页获取登录日志
// @Summary 分页获取登录日志接口
// @Produce json
// @Description 分页获取登录日志接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param loginStatus query string false "登录状态：1 ->成功 2 ->失败"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/list [get]
// @Security ApiKeyAuth
func GetSysLoginInfo(c *gin.Context) {
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	Username := c.Query("username")
	LoginStatus := c.Query("loginStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysLoginInfoService().GetSysLoginInfoList(c, Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
}
