package main

import (
	"github.com/KevinZonda/GinTemplate/controller"
	"github.com/KevinZonda/GinTemplate/share"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initCfg() {
	bs, err := iox.ReadAllByte("config.json")
	panicx.NotNilErr(err)
	panicx.NotNilErr(share.LoadConfig(bs))
}

func main() {
	initCfg()

	if share.GetConfig().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.Default()
	share.Engine = g

	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(cors.Default()) //allow all origins

	controller.Init(g)

	g.Run(share.GetConfig().Addr)
}
