package models

type ProductCategory struct {
	Id uint
	ParentId uint `gorm:"comment:上级分类的编号：0表示以及分类"`
	Name string `grom:"comment:名称"`
	Level uint `grom:"comment:分类级别：0->1级; 1->2级"`
	ProductCount uint `grom:"comment:商品数量"`
	ProductUnit string `grom:"comment:商品单位"`
	NavStatus bool `grom:"comment:是否显示在导航栏：0->不显示; 1->显示"`
	ShowStatus bool `grom:"comment:显示状态：0->不显示; 1->显示d"`
	Sort uint `gorm:"comment:排序"`
	Icon string `gorm:"图标"`
	Keywords string `gorm:"关键字"`
	Description string `gorm:"描述"`
}

func (ProductCategory) TableName() string {
	return "pms_product_category"
}