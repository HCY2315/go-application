package tools

import (
	"encoding/json"
	"go-application/common/log"
	"go-application/tools/config"
	"io/ioutil"
	"net/http"
)

// GetLocation 获取外网IP信息
func GetLocation(ip string) (string, error) {
	if ip == "127.0.0.1" || ip == "location" {
		return "内部IP", nil
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=" + config.ApplicationConfig.IpSourceKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("读取数据IP地址定位错误，err:", err)
		return "", err
	}

	m := make(map[string]string)
	err = json.Unmarshal(s, &m)
	if err != nil {
		log.Error("Umarshal failed:", err)
	}
	if m["province"] == "" {
		return "未知位置", nil
	}
	return m["province"] + "-" + m["city"], nil
}
