package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string          `gorm:"column(identities);type:varchar(36);" json:"identity"`               // 唯一标识
	ProblemIdentity string          `gorm:"column(problem_identity);type:varchar(36);" json:"problem_identity"` // 问题唯一标识
	ProblemBasic    *[]ProblemBasic `gorm:"foreignKey:identities;references:problem_identity"`                  // 关联问题基础表
	UserIdentity    string          `gorm:"column(user_identity);type:varchar(36);" json:"user_identity"`       // 用户唯一标识
	UserBasic       *UserBasic      `gorm:"foreignKey:identities;references:user_identity"`                     // 关联用户基础表
	Status          int             `gorm:"column(status);type:tinyint;" json:"status"`                         // 提交状态
	Path            string          `gorm:"column(path);type:text" json:"path"`                                 // 代码存放路径
}

func (this *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemId string, userId string, status int) *gorm.DB {
	tx := DB.Model(new(SubmitBasic)).
		Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
			return db.Omit("content")
		}).
		Preload("UserBasic", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		})
	if problemId != "" {
		tx.Where("problem_identity = ?", problemId)
	}
	if userId != "" {
		tx.Where("user_identity = ?", userId)
	}
	if status > 0 {
		tx.Where("status = ?", status)
	}

	return tx
}
