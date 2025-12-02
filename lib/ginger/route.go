package ginger

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func registerHandler(r gin.IRoutes, path string, fn ...gin.HandlerFunc) {
	terms := strings.SplitN(path, " ", 2)
	if len(terms) != 2 {
		panic("invalid path: " + path)
	}
	action, path := strings.ToLower(strings.TrimSpace(terms[0])), strings.TrimSpace(terms[1])
	switch action {
	case "get":
		r.GET(path, fn...)
	case "post":
		r.POST(path, fn...)
	case "put":
		r.PUT(path, fn...)
	case "delete":
		r.DELETE(path, fn...)
	case "patch":
		r.PATCH(path, fn...)
	case "options":
		r.OPTIONS(path, fn...)
	case "head":
		r.HEAD(path, fn...)
	case "any":
		r.Any(path, fn...)
	default:
		panic("invalid action: " + action)
	}
}

func RegisterHandlers(r gin.IRoutes, pathMap map[string][]gin.HandlerFunc) {
	for path, fn := range pathMap {
		registerHandler(r, path, fn...)
	}
}
