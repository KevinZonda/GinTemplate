package gHandler

import (
	"github.com/KevinZonda/GinTemplate/lib/ginger/tracer"
	"github.com/KevinZonda/GinTemplate/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AbortIfError(c *gin.Context, err error, code int, msg string) bool {
	if err != nil {
		Abort(c, code, msg)
		return true
	}
	return false
}

func Abort(c *gin.Context, code int, msg string) {
	traceId := tracer.GetTraceId(c)
	logger.WithTraceId(c).WithFields(logrus.Fields{
		"code": code,
		"err":  msg,
	}).Error("Gin Error Detected")
	c.JSON(code, gin.H{"error": msg, "trace_id": traceId})
	c.Abort()
}
