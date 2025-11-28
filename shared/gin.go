package shared

import (
	"github.com/KevinZonda/GinTemplate/lib/ginger/tracer"
	"github.com/KevinZonda/GinTemplate/logger"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var _engine *gin.Engine

func InitGin() {
	if GetConfig().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	_engine = gin.New()
	_engine.Use(
		gin.Recovery(),
		logger.GinHandler(),
		cors.Default(),
		tracer.MidCreateTraceId,
	)
}

func GetGin() *gin.Engine {
	return _engine
}

func RunGin() {
	err := _engine.Run(GetConfig().Addr)
	panicx.NotNilErr(err)
}
