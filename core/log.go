/*
@Desc:
@Author: qy
@Time : 2020/12/16 19:14
@File : log
@Software: GoLand
*/
package core

import (
	"github.com/sirupsen/logrus"
	"os"
)


func InitLog() *logrus.Logger{
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("log/qy_mall.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil{
		log.SetOutput(file)
	}
	log.SetLevel(logrus.InfoLevel)
	return log
}