package middlewares

import (
	"github.com/gin-gonic/gin"
)

// todo: 全局异常捕获处理
func Recovery(c *gin.Context){
	defer func() {
		return
	}()
}
