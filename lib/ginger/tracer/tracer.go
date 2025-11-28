package tracer

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const MID_TRACE_MARK = "MID_TRACE_MARK"

func MidCreateTraceId(c *gin.Context) {
	c.Set(MID_TRACE_MARK, uuid.New().String())
	c.Next()
}

func GetTraceId(c *gin.Context) string {
	v, ok := c.Get(MID_TRACE_MARK)
	if !ok || v == nil {
		return ""
	}
	return v.(string)
}
