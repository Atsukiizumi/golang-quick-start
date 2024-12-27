package models

type UserBasic struct {
	Id       int    `json:"id"`
	Identity string `json:"identity"` // 用户唯一标识
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
	Email    string `json:"email"`    // 邮箱
}

func (this *UserBasic) TableName() string {
	return "user_basic"
}
