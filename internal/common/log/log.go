/*
@Desc:
@Author: qy
@Time : 2020/12/16 19:14
@File : log
@Software: GoLand
*/
package log

import (
	"github.com/sirupsen/logrus"
	"go-gin-mall/internal/config"
	"os"
)

var (
	logger *logrus.Logger
)

func NewLogger(c config.LogConf) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(c.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	log.SetLevel(logrus.Level(c.Level))
	logger = log
	return logger
}

func GetLogger() *logrus.Logger {
	if logger == nil {
		panic("logger is unset")
	}
	return logger
}
