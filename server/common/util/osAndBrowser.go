package util

import (
	"github.com/gin-gonic/gin"
	useragent "github.com/wenlng/go-user-agent"
)

func GetOs(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	return useragent.GetOsName(userAgent)
}

func GetBrowser(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	return useragent.GetBrowserName(userAgent)
}
