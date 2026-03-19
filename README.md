# 通用基础管理系统项目

这里做一个关于Go的管理系统。

# 1. 项目初始化

## 1.1 项目搭建

首先做好项目的目录搭建。

![image-20260319110318948](README_Picture/image-20260319110318948.png)

## 1.2 项目依赖

然后安装依赖。

```cmd
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/sirupsen/logrus
go get github.com/lestrrat-go/file-rotatelogs
go get github.com/rifflock/lfshook
go get github.com/go-redis/redis/v8@v8.11.5
go get github.com/mojocn/base64Captcha@v1.3.6
go get github.com/dgrijalva/jwt-go
go get gopkg.in/yaml.v3
go get -u github.com/wenlng/go-user-agent
go get github.com/gogf/gf
go get github.com/swaggo/files
go get github.com/swaggo/gin-swagger
```

## 1.3 项目配置文件

然后写配置文件`config.yaml`进行初始配置。

```yaml
server:
  address:  :2000
  # debug模式
  model: debug
  # release模式
  # model: release
```

## 1.4 配置文件读取的初始化

在common中创建`config/config.go`，用来读取配置文件。

```go
package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// config 全局配置文件
type config struct {
	Server server `yaml:"server"`
}

// server 初始服务器配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

var Config config

// InitConfig 配置初始化，读取初始配置
func InitConfig() {

	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	// 绑定配置文件与结构体
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
```

## 1.5 主入口初始化配置文件

然后在主入口进行初始化。

```go
package main

import (
	"Go-Management-System/common/config"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

// main 启动程序，只负责启动，不负责具体逻辑
func main() {
	r := gin.Default()
	config.InitConfig()
	err := r.Run(config.Config.Server.Address)
	if err != nil {
		logger.Error("gin run error: %v", err)
	}
}
```

## 1.6 数据库配置

在`config.yaml`中配置数据库信息。

```yaml
db:
  dialects: "mysql"
  host: "localhost"
  port: "3307"
  dbName: "admin-go-api"
  username: "root"
  password: ${passowrd}
  charset: "utf8mb4"
  maxIdle: 10
  maxOpen: 100
```

然后创建这个数据库。

```mysql
create database `admin-go-api`;
```

在`config.go`中添加这个数据库的结构体并配置。

```go
// config 全局配置文件
type config struct {
	Server server `yaml:"server"`
	DB     db     `yaml:"db"`
}

// db 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}
```

## 1.7 Redis配置

接下来在`config.yaml`中配置redis。

```yaml
# redis 配置
redis:
  address: 127.0.0.1:6379
  password:
  db: 0
```

这里password留空，是因为docker上的redis没有配置密码。

然后在`config.go`中添加对应的结构体。

```go
type config struct {
	Server server `yaml:"server"`
	DB     db     `yaml:"db"`
	Redis  redis  `yaml:"redis"`
}

// redis 配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
    DB       int    `yaml:"db"`
}
```

## 1.8 图片上传地址

```yaml
# 图片上传配置
imageSettings:
  # 本地磁盘地址
  uploadDir: ./uploads/
  # 图片访问地址
  imageHost: http://localhost:2002
```

```go
// imageSettings 图片上传配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}
// config 全局配置文件
type config struct {
	Server server `yaml:"server"`
	DB     db     `yaml:"db"`
	Redis  redis  `yaml:"redis"`
	ImageSettings imageSettings `yaml:"image_settings"`
}
```

## 1.9 日志log配置

```yaml
# 日志配置
log:
  path: ./log
  name: sys
  # 输出到控制台
  model: console
  # 输出到文件
#  model: file
```

```go
// log 日志配置
type log struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
	Model string `yaml:"model"`
}

// config 全局配置文件
type config struct {
	Server        server        `yaml:"server"`
	DB            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	ImageSettings imageSettings `yaml:"image_settings"`
	Log           log           `yaml:"log"`
}
```

# 2. 完善基础配置

## 2.1 数据库配置

在pkg的db下创建`db.go`，用来初始化数据库。

```go
package db

import (
	"Go-Management-System/common/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 数据库初始化
func InitDB() {
	var err error
	// 获取config里的DB
	var dbConfig = config.Config.DB
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset,
	)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if db.Error != nil {
		panic(db.Error)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
}
```

然后在主入口中使用这个方法初始化数据库。

## 2.2 Redis配置

接下来在pkg中创建`redis/redis.go`，初始化Redis客户端。

```go
package redis

import (
	"Go-Management-System/common/config"
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisClient 得到的Redis客户端
var (
	RedisClient *redis.Client
)

// InitRedis 初始化Redis
func InitRedis() {

	// 获取config里的Redis
	redisConfig := config.Config.Redis
	// 创建Redis客户端
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	// 测试连接
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
```

然后在主入口进行初始化。

```go
package main

import (
	"Go-Management-System/common/config"
	"Go-Management-System/pkg/db"
	"Go-Management-System/pkg/redis"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

// main 启动程序，只负责启动，不负责具体逻辑
func main() {

	r := gin.Default()

	// 初始化配置文件
	config.InitConfig()

	// 初始化数据库
	db.InitDB()
	
	// 初始化Redis
	redis.InitRedis()

	err := r.Run(config.Config.Server.Address)
	if err != nil {
		logger.Error("gin run error: %v", err)
	}
}
```

## 2.3 跨域中间件

在middleware中写下面的中间件`cors.go`。

```go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 允许来自任何域名、任何端口的客户端发起的跨域请求
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许客户端发送的请求头字段
		c.Header("Access-Control-Allow-Headers", "Content-type, AccessToken, X-CSRF-Token, Authorization, Token")
		// 限制跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// JS代码获取跨域响应本来只能获取基本响应头，但这里设置了允许前端代码读取到的响应头额外信息
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// 允许在跨域请求中携带凭证
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			// OPTIONS方法预检请求，直接返回204状态码
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
```

这个是解决了Go的跨域访问的问题。

## 2.4 通用返回结构

由于存在不同的方法，后端返回给前端的数据可能是各式各样的。为了方便前端获取数据，即通过一个通用的方法来获取数据，后端需要封装一个通用的返回结构。

首先，需要实现状态码。在common的`result/code.go`中实现。

```go
package result

// Codes 定义状态
type Codes struct {
	SUCCESS uint
	FAILED  uint
	Message map[uint]string
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS: 200,
	FAILED:  501,
}

// init 初始化状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED:  "失败",
	}
}

// GetMessage 供外部使用的获取数据方法，提供状态码到状态信息的映射
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
```

这里就能实现从状态码到状态信息的映射。

然后在result下的`result.go`中实现通用信息的结构体和返回的方法。

```go
package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 消息结构体
type Result struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Data    any    `json:"data"`    // 返回的数据
}

// Success 返回成功
func Success(c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

// Failed 返回失败
func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
```

这样，根据返回的状态码和信息就能通过Success和Failed方法来返回数据。

## 2.5 鉴权中间件

这里需要验证请求的登陆用户。在common的constant下创建`constant.go`来维护系统常量。

```go
package constant

const (
	ContextKeyUserObj = "authUserObj"
)
```

然后在middleware下创建`auth.go`来进行鉴权。

```go
package middleware

import (
	"Go-Management-System/common/constant"
	"Go-Management-System/common/result"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// 未授权
			result.Failed(c, int(result.ApiCode.NOAUTHORIZATION), result.ApiCode.GetMessage(result.ApiCode.NOAUTHORIZATION))
			c.Abort()
			return
		}
		// 长度不等于2，格式错误
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHORIZATIONFORMATERROR), result.ApiCode.GetMessage(result.ApiCode.AUTHORIZATIONFORMATERROR))
			c.Abort()
			return
		}
		// todo 校验token
		var token = "token"
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
```

这里主要是从请求头中进行鉴权。未授权就使用定义好的通用错误结构来进行返回，如果成功，则将授权的token传入上下文，这样这个上下文就获取了能用的token。其中，鉴权会出现两种错误：未授权和格式错误，这两种错误需要在`code.go`中进行定义，方便直接调用。

```go
package result

// Codes 定义状态
type Codes struct {
	SUCCESS                  uint
	FAILED                   uint
	NOAUTHORIZATION          uint
	AUTHORIZATIONFORMATERROR uint
	Message                  map[uint]string
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:                  200,
	FAILED:                   501,
	NOAUTHORIZATION:          403,
	AUTHORIZATIONFORMATERROR: 405,
}

// init 初始化状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                  "成功",
		ApiCode.FAILED:                   "失败",
		ApiCode.NOAUTHORIZATION:          "未授权",
		ApiCode.AUTHORIZATIONFORMATERROR: "授权格式错误",
	}
}

// GetMessage 供外部使用的获取数据方法，提供状态码到状态信息的映射
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
```

## 2.6 日志中间件

现在需要配置日志的中间件。首先在pkg的`log/logger.go`初始化日志。

```go
package log

import (
	"Go-Management-System/common/config"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

// setLogFile 设置日志文件名
func setLogFile(file string) {
	loggerFile = file
}

// init 初始化，从配置文件中读取日志的配置信息
func init() {
	setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
}

// Log 使用日志
func Log() *logrus.Logger {
	// 如果配置文件中 Log.Model == "file"，使用文件日志
	if config.Config.Log.Model == "file" {
		// 设置日志输入到文件中
		return logFile()
	}

	// 设置日志输入到控制台
	if log == nil {
		log = logrus.New()
		log.Out = os.Stdout
		log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}

// logFile 日志方法
func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名
			loggerFile+"_%Y%m%d.log",
			// 设置最大保存时间
			rotatelogs.WithMaxAge(30*24*time.Hour),
			// 设置日志切割时间间隔
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		// 设置时间格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		// 新增Hook
		logToFile.AddHook(lfHook)

	}
	return logToFile
}
```

然后在middleware中配置日志中间件`logger.go`。

```go
package middleware

import (
	"Go-Management-System/pkg/log"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger 将请求日志打印到日志文件中
func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		header := c.Request.Header
		proto := c.Request.Proto
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		err := c.Err()
		body, _ := ioutil.ReadAll(c.Request.Body)
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
```

