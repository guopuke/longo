package router

import (
	"github.com/gin-gonic/gin"
	"longo/handler/sd"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
