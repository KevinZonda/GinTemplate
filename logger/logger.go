package logger

import (
	"os"

	"github.com/KevinZonda/GinTemplate/lib/ginger/tracer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func InitLogger() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func StdLogger() *logrus.Logger {
	return logger
}

func WithTraceId(c *gin.Context) *logrus.Entry {
	if c == nil {
		return logger.WithField("trace_id", "unk")
	}
	return logger.WithField("trace_id", tracer.GetTraceId(c))
}
