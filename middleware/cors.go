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
