package main

import (
	"github.com/KevinZonda/GinTemplate/controller"
	"github.com/KevinZonda/GinTemplate/share"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func initCfg() {
	bs, err := iox.ReadAllByte("config.json")
	panicx.NotNilErr(err)
	panicx.NotNilErr(share.LoadConfig(bs))
}

func main() {
	initCfg()
	share.Init()

	controller.Init(share.Engine)

	share.RunGin()
}
