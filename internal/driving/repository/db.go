package repository

import (
	"fmt"
	"go-gin-mall/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(conf config.DBConf) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)
	var err error
	c := &gorm.Config{}
	if conf.Debug {
		c.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dsn), c)
	if err != nil {
		panic(fmt.Sprintf("create mysql conn err: %+v", err))
	}
	return db
}
