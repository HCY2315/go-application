package models

import (
	"gorm.io/gorm"
)

type UserName struct {
	Username string `json:"username" gorm:"type:varchar(32);comment:用户名"`
}

type PassWord struct {
	Password string `json:"password" gorm:"type:varchar(32);comment:密码"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUser struct {
	BaseModel
	LoginM
	Photo       int `json:"password" gorm:"type:int(11);comment:手机号"`
	Sex         int `json:"sex" gorm:"type:int(1);comment:性别"`
	ClassRoomId int `json:"class_room_id" gorm:"type:int(11);comment:教室id"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 获取用户列表 GetAllUserList
func (e *SysUser) GetAllUserList(db *gorm.DB) ([]*SysUser, error) {
	var userList []*SysUser
	if err := db.Debug().Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}
