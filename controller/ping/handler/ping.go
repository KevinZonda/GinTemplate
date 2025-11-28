package handler

import "github.com/gin-gonic/gin"

func GetPing(c *gin.Context) (string, error) {
	return "pong", nil
}
