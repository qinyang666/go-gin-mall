package response

import "github.com/gin-gonic/gin"

type Result struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	r := &Result{
		Code:    OK,
		Message: GetMessage(OK),
		Data:    data,
	}
	c.JSON(200, r)
}

func Fail(c *gin.Context, code uint32) {
	r := &Result{
		Code:    code,
		Message: GetMessage(code),
		Data:    nil,
	}
	c.JSON(200, r)
}

func InternalServer(c *gin.Context) {
	r := &Result{
		Code:    IntervalServer,
		Message: GetMessage(IntervalServer),
		Data:    nil,
	}
	c.JSON(200, r)
}
