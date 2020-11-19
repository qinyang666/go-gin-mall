package models

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	Code string `grom:"comment:商品编码"`
	Title string `grom:"comment:商品名称"`
	Images string `grom:"comment:商品图片列表"`
	PriceSale uint `grom:"comment:销售价"`
	PriceOrigin uint `grom:"comment:原价"`
	Stock uint `grom:"comment:库存"`
	CategoryId uint `grom:"comment:分类id"`
	Desc string `gorm:"comment:描述"`
	IsOnSale bool `gorm:"是否上架：0:已下架 1:已上架"`
}