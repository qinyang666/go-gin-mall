/*
@Desc:
@Author: qy
@Time : 2020/12/15 17:23
@File : product
@Software: GoLand
*/
package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-mall/core"
	"go-gin-mall/global"
	"go-gin-mall/models"
)

type validAddProductCategory struct {
	ParentId     uint `json:"parent_id"binding:"required"`
	Name         string `json:"name" binding:"required"`
	Level        uint `json:"level"`
	//ProductCount uint `json:"product_count"`
	ProductUnit  string `json:"product_unit"`
	NavStatus    bool `json:"nav_status"`
	ShowStatus   bool `json:"show_status"`
	Sort         uint `json:"sort"`
	Icon         string `json:"icon"`
	Keywords     string `json:"keywords"`
	Description  string `json:"description"`
}

func AddProductCategory(c *gin.Context) {
	var _params validAddProductCategory
	var err error
	global.LOG.Info("test")
	//是否可以改用中间件避免判断err
	if err = core.Validate(c, &_params); err != nil{
		return
	}
	category := models.ProductCategory{}
	core.ToModel(_params, &category)
	err = global.DB.Create(&category).Error
	if err != nil {
		fmt.Print(err)
		core.Fail(c, "插入数据失败")
		return
	}
	core.Success(c)
	//global.DB.


}
