package models

import (
	"go-application/common/log"
	"go-application/tools"

	"github.com/gin-gonic/gin"
)

type UserName struct {
	Username string `json:"username" gorm:"type:varchar(32);size:64;comment:用户名"`
}

type PassWord struct {
	Password string `json:"password" gorm:"type:varchar(32);size:128;comment:密码"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUser struct {
	BaseModel
	LoginM
	Photo string `json:"password" gorm:"type:int(11);size:32;comment:手机号"`
	Sex   string `json:"sex" gorm:"type:int(1);size:32;comment:性别"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 获取用户列表 GetAllUserList
func (e *SysUser) GetAllUserList(c *gin.Context) ([]*SysUser, error) {
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error("获取数据库控制权失败！err", err)
		return nil, err
	}
	var userList []*SysUser
	if err := db.Debug().Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}
