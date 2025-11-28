package gHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleReqResp[TReq any, TResp any](fn func(c *gin.Context, req TReq) (TResp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TReq
		if err := c.ShouldBindJSON(&req); err != nil {
			AbortIfError(c, err, http.StatusBadRequest, REQ_BIND_FAILED)
			return
		}
		resp, err := fn(c, req)
		simpleFinal(c, resp, err)
	}
}

func SimpleRespOnly[TResp any](fn func(c *gin.Context) (TResp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := fn(c)
		simpleFinal(c, resp, err)
	}
}

func simpleFinal(c *gin.Context, resp any, err error) {
	if err != nil {
		AbortIfError(c, err, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, resp)

}
