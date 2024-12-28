package main

import (
	"github.com/KevinZonda/GinTemplate/controller"
	"github.com/KevinZonda/GinTemplate/share"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/gin-gonic/gin"
)

func initCfg() {
	bs, err := iox.ReadAllByte("config.json")
	panicx.NotNilErr(err)
	share.LoadConfig(bs)
}

func main() {
	initCfg()
	share.Engine = gin.Default()
	controller.Init(share.Engine)
	share.Engine.Run(share.GetConfig().Addr)
}
