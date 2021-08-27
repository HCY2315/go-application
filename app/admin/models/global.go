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
