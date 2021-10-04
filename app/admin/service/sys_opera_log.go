package service

import (
	"go-application/app/admin/models/system"
	"go-application/common/log"
	common "go-application/common/models"
	"go-application/common/service"
)

type SysOperaLog struct {
	service.Service
}

// InsertSysOperaLog 创建SysOperaLog对象
func (e *SysOperaLog) InsertSysOperaLog(model common.ActiveRecord) error {
	var err error
	var data system.SysOperaLog
	err = e.Orm.Model(&data).Create(model).Error
	if err != nil {
		log.Errorf("msgID[%s] db error:%s", e.MsgID, err)
		return err
	}
	return nil
}
