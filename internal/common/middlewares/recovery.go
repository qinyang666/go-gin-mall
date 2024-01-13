package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-gin-mall/internal/common/log"
	"go-gin-mall/internal/common/response"
	"runtime"
)

// Recovery is a server middleware that recovers from any panics.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10) //nolint:gomnd
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				log.GetLogger().Errorf("%v: %+v\n%s\n", rerr, c.Params, buf)
				response.InternalServer(c)
			}
		}()
		c.Next()
	}
}
