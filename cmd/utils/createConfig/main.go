package main

import (
	"github.com/KevinZonda/GinTemplate/lib/utils"
	"github.com/KevinZonda/GinTemplate/shared"
	"github.com/KevinZonda/GoX/pkg/iox"
)

func main() {
	cfg := shared.Config{
		Addr:  "127.0.0.1:8080",
		Debug: false,
	}
	iox.WriteAllText("config.json", utils.JsonPrettier(cfg))
}
