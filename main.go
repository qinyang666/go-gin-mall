package main

import (
	"go-gin-mall/core"
	"go-gin-mall/global"
	"go-gin-mall/models"
	"go-gin-mall/routers"
)

func main() {
	global.DB = models.InitDB()
	defer global.DB.Close()
	global.LOG = core.InitLog()
	r := routers.InitRouter()
	err := r.Run(":8080")
	if err != nil{
		global.LOG.Info("服务启动失败")
	}

}
