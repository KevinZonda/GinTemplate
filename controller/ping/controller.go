package ping

import (
	"github.com/KevinZonda/GinTemplate/controller/ping/handler"
	"github.com/KevinZonda/GinTemplate/controller/ping/svc"
	"github.com/KevinZonda/GinTemplate/lib/ginger/gHandler"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter) {
	svc.InitSvc()
	r.GET("/ping", gHandler.Resp[string](handler.GetPing))
}
