package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`   //问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`                 //关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`        //问题标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`            //问题正文
	Max_Men           int                `gorm:"column:max_men;type:int(11);" json:"max_men"`         //最大运行内存
	Max_runtime       int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` //最大运行时间
}

func (this *ProblemBasic) TableName() string {
	return "problem_basic"
}

/*func GetProblemList(keyword string) *gorm.DB {
	return GetProblemList(keyword, "")
}*/

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx = tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = {SELECT cb.id FROM category_basic cb WHERE cb.identity = ?}}", categoryIdentity)
	}
	return tx
}
