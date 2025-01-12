package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36)" json:"identity"`                 // 唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"` // 问题唯一标识
	Input           string `gorm:"column:input;type:text" json:"input"`                              // 输入
	Output          string `gorm:"column:output;type:text" json:"output"`                            // 输出
}

func (table *TestCase) TableName() string {
	return "test_case"
}
