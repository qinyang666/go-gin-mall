package main

import (
	"go-gin-mall/core"
	"go-gin-mall/global"
	"go-gin-mall/models"
	"go-gin-mall/routers"
)

func main() {
	// 初始化配置文件
	global.CONFIG = core.InitConfig()
	// 初始化日志
	global.LOG = core.InitLog()
	// 日志初始化后监控日志变化
	core.WatchConfigChange()
	// 初始化数据库连接
	global.DB = models.InitDB()
	defer global.DB.Close()
	// 路由
	r := routers.InitRouter()
	err := r.Run(":8080")
	if err != nil{
		global.LOG.Info("服务启动失败")
	}

}
