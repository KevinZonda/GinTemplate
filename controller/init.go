package controller

import (
	"github.com/KevinZonda/GinTemplate/controller/ping"
	"github.com/KevinZonda/GinTemplate/controller/types"
	"github.com/KevinZonda/GinTemplate/share"
	"github.com/gin-gonic/gin"
)

func Init(engine gin.IRouter) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(204, "pong")
	})

	register(&ping.Controller{})
}

func register(cs ...types.IController) {
	for _, c := range cs {
		c.Init(share.Engine)
	}
}
