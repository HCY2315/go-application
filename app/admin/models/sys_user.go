package models

import (
	orm "go-application/common/global"

	"gorm.io/gorm"
)

// User
type User struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}

type UserName struct {
	UserName string `gorm:"size:64;comment:用户名" json:"user_name"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"size:128;comment:密码" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT;comment:用户id" json:"user_id" `
}

type SysUserB struct {
	NickName string `gorm:"size:128;comment:昵称" json:"nick_name"`       // 昵称
	Phone    string `gorm:"size:11;comment:手机号" json:"phone"`           // 手机号
	RoleId   int    `gorm:"comment:角色编码" json:"role_id"`                // 角色编码
	Salt     string `gorm:"varchar(32);comment:盐;size:255" json:"salt"` // 盐
	Avatar   string `gorm:"size:255;comment:头像" json:"avatar"`          // 头像
	Sex      string `gorm:"size:255;comment:性别" json:"sex"`             // 性别
	Email    string `gorm:"size:128;comment:邮箱" json:"email"`           // 邮箱
	DeptId   int    `gorm:"comment:部门编码" json:"dept_id"`                // 部门编码
	PostId   int    `gorm:"comment:职位编码" json:"post_id"`                // 职位编码
	CreateBy string `gorm:"size:128;comment:创建人" json:"create_by"`      // 创建人
	UpdateBy string `gorm:"size:128;comment:修改人" json:"update_by"`      // 修改人
	Remark   string `gorm:"size:255;comment:备注" json:"remark"`          // 备注
	Status   string `gorm:"size:4;comment:用户状态" json:"status"`          // 用户状态
	BaseModel

	DataScope string `gorm:"-" json:"data_scope"`
	Params    string `gorm:"-" json:"params"`
}

type SysUser struct {
	SysUserId
	LoginM
	SysUserB
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
