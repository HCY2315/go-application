package models

type UserName struct {
	Username string `gorm:"size:64" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"size:128" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUser struct {
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}
