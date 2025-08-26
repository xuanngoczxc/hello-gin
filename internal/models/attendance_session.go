package models

import (
	"time"

	"gorm.io/gorm"
)

// AttendanceSession represents an attendance session
type AttendanceSession struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	EventID     *uint      `json:"event_id"`
	ClassID     *uint      `json:"class_id"`
	TeacherID   *uint      `json:"teacher_id"`
	SessionDate *time.Time `json:"session_date"`

	// Relationships
	Event       *Event       `json:"event,omitempty"`
	Class       *Class       `json:"class,omitempty"`
	Teacher     *Teacher     `json:"teacher,omitempty"`
	Attendances []Attendance `gorm:"foreignKey:SessionID" json:"attendances,omitempty"`
}

// TableName overrides the table name used by AttendanceSession
func (AttendanceSession) TableName() string {
	return "attendance_sessions"
}
