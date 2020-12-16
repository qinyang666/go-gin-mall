/*
@Desc:
@Author: qy
@Time : 2020/12/16 9:48
@File : core
@Software: GoLand
*/
package core

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 0
	ERROR = 100
)


type response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func RenderResponse(c *gin.Context, code int, data interface{}, msg string){
	res := response{
		Code: code,
		Data: data,
		Message: msg,
	}
	c.JSON(200, res)
}

func Success(c *gin.Context) {
	RenderResponse(c, SUCCESS, nil, "请求成功")
}

func Fail(c *gin.Context, msg string) {
	RenderResponse(c, ERROR, nil, msg)
}

func SuccessWithData(c *gin.Context, data interface{}){
	RenderResponse(c, SUCCESS, data, "请求成功")
}
