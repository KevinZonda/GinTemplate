package ginger

import "github.com/gin-gonic/gin"

type IController interface {
	Init(r gin.IRouter)
}

func Register(r gin.IRouter, cs ...GinRegisterFx) {
	for _, c := range cs {
		c(r)
	}
}

type GinRegisterFx func(r gin.IRouter)
