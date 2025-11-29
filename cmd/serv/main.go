package main

import (
	"flag"

	"github.com/KevinZonda/GinTemplate/controller"
	"github.com/KevinZonda/GinTemplate/shared"
)

func main() {
	cfgPath := flag.String("cfg", "config.json", "config path")
	flag.Parse()

	shared.InitAll(*cfgPath)
	controller.Init(shared.GetGin())

	shared.RunGin()
}
