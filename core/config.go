/*
@Desc:
@Author: qy
@Time : 2020/12/17 15:29
@File : config
@Software: GoLand
*/
package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-gin-mall/global"
)

func InitConfig()*viper.Viper{
	_viper := viper.New()
	_viper.SetConfigName("config")
	_viper.SetConfigType("yaml")
	_viper.AddConfigPath(".")
	err := _viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return _viper
}

// 需在日志初始化后执行
func WatchConfigChange(){
	// 监控配置变化
	global.CONFIG.WatchConfig()
	global.CONFIG.OnConfigChange(func (e fsnotify.Event){
		global.LOG.Warn("配置文件发生改变:" + e.Name)
	})
}