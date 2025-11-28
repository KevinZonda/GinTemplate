package template

import (
	"github.com/KevinZonda/GinTemplate/controller/template/handler"
	"github.com/KevinZonda/GinTemplate/controller/template/svc"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter) {
	svc.InitSvc()
	r.GET("/ping", handler.GetPing)
}
