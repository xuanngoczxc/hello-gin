package models

import "time"

// Teacher represents a teacher in the system
type Teacher struct {
	ID          int       `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	TeacherCode *string   `json:"teacher_code" example:"GV001"`
	TeacherName *string   `json:"teacher_name" example:"Nguyen Thi B"`
}

// TableName overrides the table name used by Teacher to `Teacher`
func (Teacher) TableName() string {
	return "Teacher"
}
