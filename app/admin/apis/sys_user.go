package apis

import (
	"go-application/app/admin/models"
	"go-application/common/apis"
	"go-application/common/log"
	"go-application/tools"

	"github.com/gin-gonic/gin"
)

type ApiSysUser struct {
	apis.Api
}

func (e *ApiSysUser) GetAll(c *gin.Context) {
	sysUser := new(models.SysUser)
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Errorf("获取数据库控制权失败！err", err)
		return
	}
	list, err := sysUser.GetAllUserList(db)
	if err != nil {
		log.Errorf("MsgID[%s]获取数据失败, err:", msgID, err)
		e.Error(c, 404, nil, "获取数据失败")
		return
	}
	e.Ok(c, list, "")
}
