package tools

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
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
