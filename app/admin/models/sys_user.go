package models

import (
	orm "go-application/common/global"

	"gorm.io/gorm"
)

type UserName struct {
	UserName string `json:"user_name" gorm:"type:varchar(32);comment:用户名"`
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
	ClassroomId int `json:"class_room_id" gorm:"type:int(11);comment:教室id"`
}

type GetUser struct {
	BaseModel
	LoginM
	Photo         int    `json:"password" gorm:"column:photo"`
	Sex           int    `json:"sex" gorm:"column:sex"`
	ClassroomName string `json:"classroom_name" gorm:"column:classroom_name"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) GetUserNoPage() (GetUser []GetUser, err error) {
	table := orm.Eloquent.Table(e.TableName()).Select([]string{"sys_user.*", "op_classroom.classroom_name"})
	table = table.Debug().Joins("left join op_classroom on sys_user.classroom_id=op_classroom.id")

	if e.UserName.UserName != "" {
		table.Where("user_name = ?", e.UserName.UserName)
	}

	// TODO: 添加其他查询条件
	if err = table.Find(&GetUser).Error; err != nil {
		return
	}
	return
}

// 获取用户列表 GetAllUserList
func (e *SysUser) GetAllUserList(db *gorm.DB) ([]*SysUser, error) {
	var userList []*SysUser
	if err := db.Debug().Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}
