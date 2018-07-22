package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/lexkong/log"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exits
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		// 32 位的 UUID，用于唯一标识一次 HTTP 请求
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)
		log.Infof("Header Request Id: %s", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
