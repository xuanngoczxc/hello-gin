package models

import "time"

// AttendanceSession represents an attendance session in the system
type AttendanceSession struct {
	ID          int          `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt   time.Time    `json:"created_at" example:"2023-01-01T00:00:00Z"`
	ClassID     *int         `json:"class_id" example:"1"`
	TeacherID   *int         `json:"teacher_id" example:"1"`
	SessionDate *time.Time   `json:"session_date" example:"2023-01-01T09:00:00Z"`
	Class       *Class       `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Teacher     *Teacher     `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
	Attendances []Attendance `json:"attendances,omitempty" gorm:"foreignKey:SessionID"`
}

// TableName overrides the table name used by AttendanceSession
func (AttendanceSession) TableName() string {
	return "AttendanceSession"
}
