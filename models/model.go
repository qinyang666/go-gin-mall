/*
@Desc: db初始化及model相关公用项
@Author: qy
@Time : 2020/12/10 10:31
@File : model
@Software: GoLand
*/
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:asinking@/qy_mall?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
