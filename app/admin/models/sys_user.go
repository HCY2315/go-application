package models

import "gorm.io/gorm"

type UserName struct {
	Username string `json:"username" gorm:"size:64;comment:用户名"`
}

type PassWord struct {
	Password string `json:"password" gorm:"size:128;comment:密码"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUser struct {
	gorm.Model
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}
