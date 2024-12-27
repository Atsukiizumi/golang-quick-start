package models

import "gorm.io/gorm"

type ProductBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(50)" json:"identity"` // 产品唯一标识
	Name     string `gorm:"column:name;type:varchar(50)" json:"name"`         // 产品名称
	Desc     string `gorm:"column:desc;type:varchar(50)" json:"desc"`         // 产品描述
}

func (this *ProductBasic) TableName() string {
	return "product_basic"
}
