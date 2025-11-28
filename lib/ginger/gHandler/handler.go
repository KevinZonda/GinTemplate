package gHandler

import (
	"net/http"

	"github.com/KevinZonda/GinTemplate/logger"
	"github.com/gin-gonic/gin"
)

func ReqResp[TReq any, TResp any](fn func(c *gin.Context, req TReq, resp *Resp[TResp])) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TReq
		if err := c.ShouldBindJSON(&req); err != nil {
			AbortIfError(c, err, http.StatusBadRequest, REQ_BIND_FAILED)
			return
		}
		resp := NewResp[TResp]()
		fn(c, req, resp)
		final(c, resp)
	}
}

func RespOnly[TResp any](fn func(c *gin.Context, resp *Resp[TResp])) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := NewResp[TResp]()
		fn(c, resp)
		final(c, resp)
	}
}

func final[T any](c *gin.Context, resp *Resp[T]) {
	if c.IsAborted() {
		return
	}
	if resp.Error != "" {
		logger.WithTraceId(c).Error(resp.E.Error())
		Abort(c, http.StatusInternalServerError, resp.Error)
		return
	}
	c.JSON(http.StatusOK, resp.Data)
}
