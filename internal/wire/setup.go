package wire

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-mall/internal/config"
	"go-gin-mall/internal/handler"
)

func Setup(r *gin.Engine, c *config.Config, logger *logrus.Logger) {
	userHandler := wireUserHandler(c.DBConf, logger)
	handler.RegisterUserHandler(r, userHandler)
	addr := fmt.Sprintf("%s:%d", c.ServerConf.Host, c.ServerConf.Port)
	fmt.Printf("start server: %s ...", addr)
	if err := r.Run(addr); err != nil {
		logger.Printf("stop server: %s ...", addr)
	}
}
