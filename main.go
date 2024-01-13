package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-gin-mall/internal/common/log"
	"go-gin-mall/internal/common/middlewares"
	"go-gin-mall/internal/config"
	"go-gin-mall/internal/wire"
)

var configFile = flag.String("f", "etc", "the config file")

func main() {
	c := config.NewConfig(*configFile)
	gin.SetMode(c.Mode)
	r := gin.Default()
	r.Use(middlewares.Recovery())
	logger := log.NewLogger(c.LogConf)
	wire.Setup(r, c, logger)
}
