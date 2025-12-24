package gHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReqResp[TReq any, TResp any](fn func(c *gin.Context, req TReq) (TResp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TReq
		if err := c.ShouldBindJSON(&req); err != nil {
			AbortIfError(c, err, http.StatusBadRequest, REQ_BIND_FAILED)
			return
		}
		resp, err := fn(c, req)
		final(c, resp, err)
	}
}

func Req[TReq any](fn func(c *gin.Context, req TReq) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TReq
		if err := c.ShouldBindJSON(&req); err != nil {
			AbortIfError(c, err, http.StatusBadRequest, REQ_BIND_FAILED)
			return
		}
		err := fn(c, req)
		finalEmpty(c, err)
	}
}

func ReqDefaultResp[TReq any](fn func(c *gin.Context, req TReq) error, defaultResp any) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TReq
		if err := c.ShouldBindJSON(&req); err != nil {
			AbortIfError(c, err, http.StatusBadRequest, REQ_BIND_FAILED)
			return
		}
		err := fn(c, req)
		final(c, defaultResp, err)
	}
}

func Resp[TResp any](fn func(c *gin.Context) (TResp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := fn(c)
		final(c, resp, err)
	}
}

func NoReqResp(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		finalEmpty(c, err)
	}
}

func TextResp(fn func(c *gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := fn(c)
		finalString(c, resp, err)
	}
}

func final(c *gin.Context, resp any, err error) {
	if !handleError(c, err) {
		c.JSON(200, resp)
	}
}

func finalString(c *gin.Context, resp string, err error) {
	if !handleError(c, err) {
		c.String(200, resp)
	}
}

func finalEmpty(c *gin.Context, err error) {
	if !handleError(c, err) {
		c.Status(200)
	}
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		AbortIfError(c, err, http.StatusInternalServerError, err.Error())
		return true
	}
	return false
}

func EmptyMap() map[string]any {
	return map[string]any{}
}
