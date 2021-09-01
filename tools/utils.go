package tools

import (
	"fmt"
	"go-application/common/log"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	var msgID string
	data, ok := c.Get("msgID")
	if !ok {
		msgID = uuid.New().String()
		c.Set("msgID", msgID)
		return msgID
	}
	msgID = cast.ToString(data)
	return msgID
}

// GetOrm 获取orm连接
func GetOrm(c *gin.Context) (*gorm.DB, error) {
	msgID := GenerateMsgIDFromContext(c)
	idb, exist := c.Get("db")
	if !exist {
		return nil, fmt.Errorf("msgID[%s], db connect not exist", msgID)
	}
	switch idb := idb.(type) {
	case *gorm.DB:
		return idb, nil
	default:
		return nil, fmt.Errorf("msgID[%s], db connect not exist", msgID)
	}
}

// CompareHashAndPassword 比较两个hash值是是否一样
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		_, file, line, _ := runtime.Caller(1)
		log.Info("%s:%v error: %#v", file, line, err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
