package logger

import (
	"time"

	"github.com/KevinZonda/GinTemplate/lib/ginger/tracer"
	"github.com/gin-gonic/gin"
)

func GinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		traceId := tracer.GetTraceId(c)
		if reqMethod == "HEAD" {
			return
		}
		logger.Infof("%s | %3d | %13v | %15s | %s | %s | %s |",
			startTime.Format("2006-01-02 15:04:05"),
			statusCode,
			latencyTime,
			clientIP,
			traceId,
			reqMethod,
			reqUri,
		)
	}
}
