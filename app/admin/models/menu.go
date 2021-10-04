package models

import orm "go-application/common/global"

type Menu struct {
	MenuId     int    `json:"menuId" gorm:"primary_key;AUTO_INCREMENT"`
	MenuName   string `json:"menuName" gorm:"size:128;"`
	Title      string `json:"title" gorm:"size:128;"`
	Icon       string `json:"icon" gorm:"size:128;"`
	Path       string `json:"path" gorm:"size:128;comment:接口"` // 接口
	Paths      string `json:"paths" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Action     string `json:"action" gorm:"size:16;comment:访问方法"` // 访问方法
	Permission string `json:"permission" gorm:"size:255;"`
	ParentId   int    `json:"parentId" gorm:"size:11;"`
	NoCache    bool   `json:"noCache" gorm:"size:8;"`
	Breadcrumb string `json:"breadcrumb" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int    `json:"sort" gorm:"size:4;"`
	Visible    string `json:"visible" gorm:"size:1;"`
	CreateBy   string `json:"createBy" gorm:"size:128;"`
	UpdateBy   string `json:"updateBy" gorm:"size:128;"`
	IsFrame    string `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	DataScope  string `json:"dataScope" gorm:"-"`
	Params     string `json:"params" gorm:"-"`
	RoleId     int    `gorm:"-"`
	Children   []Menu `json:"children" gorm:"-"`
	IsSelect   bool   `json:"is_select" gorm:"-"`
	BaseModel
}

func (Menu) TableName() string {
	return "sys_menu"
}

// 通过接口和访问方式获取菜单列表
func (e *Menu) Get() (Menus []Menu, err error) {
	table := orm.Eloquent.Table(e.TableName())
	if e.Path != "" {
		table = table.Where("path = ?", e.Path)
	}
	if e.Action != "" {
		table = table.Where("action = ?", e.Action)
	}
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}
