package dao

// dao :负责数据的持久化工作，将下层存储已更简单的函数、接口形式暴露给控制层
import "go-application/app/admin/models"

type SerSysUser struct {
	Id       int `json:"id"`
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
	Req      models.SysUser
}

func (dao *SerSysUser) GetUserByPage() ([]*models.SysUser, error) {
	return dao.Req.GetAllUserList()
}
