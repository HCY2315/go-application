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

func (e *ApiSysUser) GetNotPage(c *gin.Context) {
	data := new(models.SysUser)
	msgID := tools.GenerateMsgIDFromContext(c)
	_, err := tools.GetOrm(c)
	if err != nil {
		log.Errorf("获取数据库控制权失败！err", err)
		return
	}

	data.UserName.UserName = c.Request.FormValue("user_name")
	list, err := data.GetUserNoPage()
	if err != nil {
		log.Errorf("MsgID[%s]获取数据失败, err:", msgID, err)
		e.Error(c, 404, nil, "获取数据失败")
		return
	}
	e.Ok(c, list, "")
}
