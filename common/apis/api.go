package apis

import (
	"go-application/common/models"
	"go-application/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
}

//获取orm连接 GetOrm
func (e *Api) GetOrm(c *gin.Context) (*gorm.DB, error) {
	return tools.GetOrm(c)
}

// Error 通常错误数据处理
func (e *Api) Error(c *gin.Context, code int, err error, msg string) {
	var res models.Response
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	res.RequestId = tools.GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}

func (e *Api) Ok(c *gin.Context, data interface{}, msg string) {
	var res models.Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	res.RequestId = tools.GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnOK())
}

func (e *Api) PageOk(c *gin.Context, result interface{}, count int, pageNum int, pageSize int, msg string) {
	var res models.Page
	res.List = result
	res.Count = count
	res.PageNum = pageNum
	res.PageSize = pageSize
	e.Ok(c, res, msg)
}
