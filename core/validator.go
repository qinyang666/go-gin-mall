/*
@Desc:
@Author: qy
@Time : 2020/12/16 15:12
@File : validator
@Software: GoLand
*/
package core

import (
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context, obj interface{})error{
	err := c.ShouldBindJSON(obj)
	// todo: 封装根据错误返回对应响应信息或者调用第三方库
	if err != nil{
		Fail(c, "参数错误")
	}
	return err
}
