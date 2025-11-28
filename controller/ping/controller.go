package ping

import (
	"github.com/KevinZonda/GinTemplate/controller/ping/handler"
	"github.com/KevinZonda/GinTemplate/controller/ping/svc"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter) {
	svc.InitSvc()
	r.GET("/ping", handler.GetPing)
}
