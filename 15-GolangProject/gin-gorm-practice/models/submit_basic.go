package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string `gorm:"column(identity);type:varchar(36);" json:"identity"`                 // 唯一标识
	ProblemIdentity string `gorm:"column(problem_identity);type:varchar(36);" json:"problem_identity"` // 问题唯一标识
	UserIdentity    string `gorm:"column(user_identity);type:varchar(36);" json:"user_identity"`       // 用户唯一标识
	Status          int    `gorm:"column(status);type:tinyint;" json:"status"`                         // 提交状态
	Path            string `gorm:"column(path);type:text" json:"path"`                                 // 代码存放路径
}

func (this *SubmitBasic) TableName() string {
	return "submit_basic"
}
