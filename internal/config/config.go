/*
@Desc:
@Author: qy
@Time : 2020/12/17 15:29
@File : config
@Software: GoLand
*/
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	LogConf LogConf
	DBConf  DBConf
	ServerConf
}

func NewConfig(path string) *Config {
	var c Config
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	err = v.Unmarshal(&c)
	if err != nil {
		panic(fmt.Sprintf("Fatal unmarshal config: %s \n", err))
	}
	return &c
}
