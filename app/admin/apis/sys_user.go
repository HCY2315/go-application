package apis

import (
	"fmt"
	"go-application/common/apis"

	"github.com/gin-gonic/gin"
)

type ApiSysUser struct {
	apis.Api
}

func (e *ApiSysUser) GetAll(c *gin.Context) {
	fmt.Println(c.Get("db"))
	e.Ok(c, "", "")
}
