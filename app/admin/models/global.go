package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	// ID        uint           `json:"id" gorm:"primarykey;comment:唯一标识"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"update_at" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:删除时间"`
}

type ControlBy struct {
	CreateBy uint `json:"create_by" gorm:"index;comment:'创建者'"`
	UpdateBy uint `json:"update_by" gorm:"index;comment:'更新者'"`
}

func (e *ControlBy) SetCreateBy(createBy uint) {
	e.CreateBy = createBy
}

func (e *ControlBy) SetUpdateBy(updateBy uint) {
	e.UpdateBy = updateBy
}
