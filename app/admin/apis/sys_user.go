package apis

import (
	"fmt"
	"go-application/app/admin/service"
	"go-application/common/apis"
	"go-application/common/log"

	"github.com/gin-gonic/gin"
)

type ApiSysUser struct {
	apis.Api
}

func (e *ApiSysUser) GetAll(c *gin.Context) {
	sysUserParam := new(service.SerSysUser)
	err := c.Bind(sysUserParam)
	if err != nil {
		log.Error("参数错误", err)
		e.Error(c, 404, nil, "参数错误")
		return
	}

	list, err := sysUserParam.GetAllUserList(c)
	if err != nil {
		log.Error("获取数据失败", err)
		e.Error(c, 404, nil, "获取数据失败")
		return
	}
	fmt.Println(c.Get("db"))
	e.Ok(c, list, "")
}
