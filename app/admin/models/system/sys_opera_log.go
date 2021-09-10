package system

import (
	"time"

	"go-application/app/admin/models"
)

type SysOperaLog struct {
	models.BaseModel
	models.ControlBy

	Title         string    `json:"title" gorm:"type:varchar(255);comment:操作模块"`                   //
	BusinessType  string    `json:"business_type" gorm:"type:varchar(128);comment:操作类型"`           //
	BusinessTypes string    `json:"business_types" gorm:"type:varchar(128);comment:BusinessTypes"` //
	Method        string    `json:"method" gorm:"type:varchar(128);comment:函数"`                    //
	RequestMethod string    `json:"request_method" gorm:"type:varchar(128);comment:请求方式"`          //
	OperatorType  string    `json:"operator_type" gorm:"type:varchar(128);comment:操作类型"`           //
	OperName      string    `json:"oper_name" gorm:"type:varchar(128);comment:操作者"`                //
	DeptName      string    `json:"dept_name" gorm:"type:varchar(128);comment:部门名称"`               //
	OperUrl       string    `json:"oper_url" gorm:"type:varchar(255);comment:访问地址"`                //
	OperIp        string    `json:"oper_ip" gorm:"type:varchar(128);comment:客户端ip"`                // 客户端ip
	OperLocation  string    `json:"oper_location" gorm:"type:varchar(128);comment:访问位置"`           //
	OperParam     string    `json:"oper_param" gorm:"type:varchar(255);comment:请求参数"`              //
	Status        string    `json:"status" gorm:"type:varchar(4);comment:操作状态"`                    //
	OperTime      time.Time `json:"oper_time" gorm:"type:timestamp;comment:操作时间"`                  //
	JsonResult    string    `json:"json_result" gorm:"type:varchar(255);comment:返回数据"`             //
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:备注"`                    //
	LatencyTime   string    `json:"latency_time" gorm:"type:varchar(128);comment:耗时"`              //
	UserAgent     string    `json:"user_agent" gorm:"type:varchar(255);comment:ua"`                //
}

func (SysOperaLog) TableName() string {
	return "sys_opera_log"
}
