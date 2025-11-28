package main

import (
	"github.com/KevinZonda/GinTemplate/controller"
	"github.com/KevinZonda/GinTemplate/shared"
)

func main() {
	shared.InitAll("config.json")
	controller.Init(shared.GetGin())

	shared.RunGin()
}
