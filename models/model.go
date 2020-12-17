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
	"go-gin-mall/global"
	"log"
)

func InitDB() *gorm.DB {
	mysqlConfig := global.CONFIG.GetStringMapString("mysql")
	host, _ := mysqlConfig["host"]
	port, _ := mysqlConfig["port"]
	username, _ := mysqlConfig["username"]
	password, _ := mysqlConfig["password"]
	dbName, _ := mysqlConfig["db_name"]

	db, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbName+ "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
