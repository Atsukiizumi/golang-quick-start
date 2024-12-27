package models

type UserBasic struct {
	Identity string `gorm:"column:identity;type:varchar(50)" json:"identity"` // 用户唯一标识
	Name     string `gorm:"column:name;type:varchar(50)" json:"name"`         // 用户名
	Password string `gorm:"column:password;type:varchar(50)" json:"password"` // 用户密码
}

func (this *UserBasic) TableName() string {
	return "user_basic"
}
