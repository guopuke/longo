package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoCache is a middleware function that appends headers to prevent the client from caching the HTTP response.
// 强制浏览器不使用缓存
func NoCache(c *gin.Context) {
	c.Header("Cache-control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request
// 浏览器跨域 OPTIONS 设置
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure is a middleware function that append security
// and resource access headers
// 安全设置
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Context-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	// Also consider adding Content-Security-Policy(来源白名单) headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://website.com")

	if c.Request.TLS != nil {
		// 1. 在接下来的 31536000 秒（即一年）中，浏览器向 example.com 或其子域名发送 HTTP 请求时，
		// 必须采用 HTTPS 来发起连接。比如，用户点击超链接或在地址栏输入 http://www.example.com/ ，
		// 浏览器应当自动将 http 转写成 https，然后直接向 https://www.example.com/ 发送请求。
		// 2. 在接下来的一年中，如果 example.com 服务器发送的 TLS 证书无效，用户不能忽略浏览器警告继续访问网站。
		c.Header("Strict-Transport-Security", "Max-age=31536000")
	}
}
