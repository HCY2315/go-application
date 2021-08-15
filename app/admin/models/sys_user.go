package models

import (
	"go-application/common/global"

	"gorm.io/gorm"
)

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

// 获取用户列表 GetAllUserList
func (e *SysUser) GetAllUserList() ([]*SysUser, error) {
	table := global.Eloquent.Table(e.TableName())
	var userList []*SysUser
	if err := table.Where("").Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}
