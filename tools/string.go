package tools

import (
	"encoding/json"
	"time"
)

// 获取当时间戳
func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 将结构体转化成json字符串
func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}
