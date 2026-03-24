package router

import (
	"Go-Management-System/api/controller"
	"Go-Management-System/common/config"
	"Go-Management-System/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 创建初始的Engine
	router := gin.New()
	// 配置中间件
	// 1. 配置Recovery
	router.Use(gin.Recovery())
	// 2. 配置跨域请求
	router.Use(middleware.Cors())
	// 3. 配置文件上传服务
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	// 4. 配置日志中间件
	router.Use(middleware.Logger())

	// 5. 注册路由
	register(router)

	return router
}

// register 路由注册
func register(router *gin.Engine) {
	// todo 添加接口url
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)
	router.POST("/api/post/add", controller.CreateSysPost)
	router.GET("/api/post/list", controller.GetSysPostList)
	router.GET("/api/post/info", controller.GetSysPostById)
	router.PUT("/api/post/update", controller.UpdateSysPost)
	router.DELETE("/api/post/delete", controller.DeleteSysPostById)
	router.DELETE("/api/post/batch/delete", controller.BatchDeleteSysPost)
	router.PUT("/api/post/updateStatus", controller.UpdateSysPostStatus)
	router.GET("/api/post/vo/list", controller.QuerySysPostVOList)
	router.GET("/api/dept/list", controller.GetSysDeptList)
	router.POST("/api/dept/add", controller.CreateSysDept)
	router.GET("/api/dept/info", controller.GetSysDeptById)
	router.PUT("/api/dept/update", controller.UpdateSysDept)
	router.DELETE("/api/dept/delete", controller.DeleteSysDeptById)
	router.GET("/api/dept/vo/list", controller.QuerySysDeptVOList)
	router.POST("/api/menu/add", controller.CreateSysMenu)
	router.GET("/api/menu/list", controller.QuerySysMenuVOList)
	router.GET("/api/menu/info", controller.GetSysMenuById)
	router.PUT("/api/menu/update", controller.UpdateSysMenu)
}
