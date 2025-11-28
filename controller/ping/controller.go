package ping

import (
	"github.com/KevinZonda/GinTemplate/controller/ping/handler"
	"github.com/KevinZonda/GinTemplate/controller/ping/svc"
	"github.com/KevinZonda/GinTemplate/controller/types"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

var _ types.IController = (*Controller)(nil)

func (c *Controller) Init(r gin.IRouter) {
	svc.InitSvc()
	r.GET("/ping", handler.GetPing)
}
