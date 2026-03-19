package main

import (
	"Go-Management-System/common/config"
	_ "Go-Management-System/docs"
	"Go-Management-System/pkg/db"
	"Go-Management-System/pkg/log"
	"Go-Management-System/pkg/redis"
	"Go-Management-System/router"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// main 启动程序，只负责启动，不负责具体逻辑
// @title 通用后台管理系统
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// 初始化日志
	log := log.Log()

	// 确认Gin模式，当前为Debug
	gin.SetMode(config.Config.Server.Model)

	// 初始化路由，准备好了中间件
	r := router.InitRouter()

	// 将初始化的路由传给http.Server
	srv := &http.Server{
		// 定义路由的监听地址
		Addr: config.Config.Server.Address,
		// 将路由Engine传入
		Handler: r,
	}

	// 使用协程来启动服务
	go func() {
		// 通过ListenAndServe启动路由
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Info("listen: %s\n", err)
		}
		// 启动成功，打印监听地址
		log.Info("listen: %s\n", config.Config.Server.Address)
	}()
	// 创建信号通道，将关闭程序的信号装入channel中
	quit := make(chan os.Signal)
	// 监听消息
	signal.Notify(quit, os.Interrupt)
	// 一直阻塞，直到收到关闭信号
	<-quit
	log.Info("Shutdown Server ...")
	// 等待5秒，确保所有请求都处理完毕
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 尝试关闭路由
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown: ", err)
	}
	log.Info("Server exiting")

}

// 初始化连接
func init() {
	// 读取配置文件
	config.InitConfig()
	// 初始化数据库
	db.InitDB()
	// 初始化Redis
	redis.InitRedis()
}
