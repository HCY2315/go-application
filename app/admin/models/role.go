package models

type SysRole struct {
	RoleId    int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT;comment:角色编码"` // 角色编码
	RoleName  string `json:"roleName" gorm:"size:128;comment:角色名称"`                 // 角色名称
	Status    string `json:"status" gorm:"size:4;comment:角色状态"`                     //角色状态
	RoleKey   string `json:"roleKey" gorm:"size:128;comment:角色代码"`                  //角色代码
	RoleSort  int    `json:"roleSort" gorm:"comment:角色排序"`                          //角色排序
	Flag      string `json:"flag" gorm:"size:128;comment:标志"`                       //标注
	CreateBy  string `json:"createBy" gorm:"size:128;comment:创建人"`                  //创建人
	UpdateBy  string `json:"updateBy" gorm:"size:128;comment:修改人"`                  //修改人
	Remark    string `json:"remark" gorm:"size:255;comment:备注"`                     //备注
	Admin     bool   `json:"admin" gorm:"size:4;comment:管理员权限"`                     // 1:管理员；2:非管理员
	DataScope string `json:"dataScope" gorm:"size:128;comment:数据范围"`                // 数据范围
	BaseModel

	Params  string `json:"params" gorm:"-"`
	MenuIds []int  `json:"menuIds" gorm:"-"`
	DeptIds []int  `json:"deptIds" gorm:"-"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
