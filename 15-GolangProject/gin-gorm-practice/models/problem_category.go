package models

import "gorm.io/gorm"

// 问题与分类关联表
type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `gorm:"column:problem_id;type:int(11);" json:"problem_id"`   // 问题ID
	CategoryId    string         `gorm:"column:category_id;type:int(11);" json:"category_id"` // 分类ID
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`                // 关联分类的基础信息表
}

func (this *ProblemCategory) TableName() string {
	return "problem_category"
}
