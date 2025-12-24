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
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
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

func G(c *gin.Context) *logrus.Entry {
	if c == nil {
		return logger.WithField("trace_id", "unk")
	}
	e := logger.WithField("trace_id", tracer.GetTraceId(c))
	e = e.WithField("uri", c.Request.RequestURI)
	return e
}

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Print(v ...interface{}) {
	logger.Print(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func RecordIfError(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func RecordIfErrorWithTrace(c *gin.Context, err error) {
	if err != nil {
		G(c).Error(err)
	}
}
