package controller

import (
	"github.com/KevinZonda/GinTemplate/controller/ping"
	"github.com/KevinZonda/GinTemplate/lib/ginger"
	"github.com/gin-gonic/gin"
)

func Init(r gin.IRouter) {
	ginger.Register(r, true,
		ping.Register,
	)
}
