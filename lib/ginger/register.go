package ginger

import "github.com/gin-gonic/gin"

type IController interface {
	Init(r gin.IRouter)
}

func Register(r gin.IRouter, isolation bool, cs ...GinRegisterFx) {
	for _, c := range cs {
		g := r
		if isolation {
			g = r.Group("/")
		}
		c(g)
	}
}

type GinRegisterFx func(r gin.IRouter)
