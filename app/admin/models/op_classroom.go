package models

type OpClassRoom struct {
	BaseModel
	ClassroomName string `json:"class_room_name" gorm:"type:varchar(32);size:128;comment:教室名称"`
}

func (OpClassRoom) TableName() string {
	return "op_classroom"
}
