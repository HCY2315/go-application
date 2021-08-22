package router

import (
	"go-application/app/admin/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRoleV1 = append(routerNoCheckRoleV1, registerSysUser)
}

func registerSysUser(v1 *gin.RouterGroup) {
	r1 := v1.Group("/sysuser")
	{
		sysUserApi := &apis.ApiSysUser{}
		r1.POST("/list", sysUserApi.GetNotPage)
	}
}
