package models

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Phone string `gorm:"comment:登录手机号"`
	Password string `gorm:"密码"`
	Nickname string `gorm:"comment:昵称"`
	Avatar string `gorm:"comment:头像"`
	Address string `gorm:"comment:地址"`
	Linkman string `gorm:"comment:联系人"`
	LinkPhone string `gorm:"comment:联系电话"`
	Balance uint `gorm:"comment:余额"`
}

type BalanceLog struct {
	gorm.Model
	Code string `grom:"comment:单号"`
	CodeType string `grom:"comment:单号类型"`
	MemberId uint `grom:"comment:用户id"`
	OldBalance uint `grom:"comment:变更前余额"`
	NewBalance uint `grom:"comment:变更后余额"`
	Amount uint `grom:"comment:变化金额"`
	ChangeReason string `grom:"comment:变更原因"`
	Remark string `grom:"comment:备注"`
}