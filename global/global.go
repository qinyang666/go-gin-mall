/*
@Desc:
@Author: qy
@Time : 2020/12/15 16:21
@File : global
@Software: GoLand
*/
package global

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
	LOG *logrus.Logger
)
