package models

import (
	"time"

	"gorm.io/gorm"
)

// Attendance represents an attendance record in the system
type Attendance struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	SessionID       *uint      `json:"session_id"`
	CheckedInAt     *time.Time `json:"checked_in_at"`
	StudentName     *string    `json:"student_name"`
	Email           *string    `json:"email"`
	Phone           *string    `json:"phone"`
	WorkUnit        *string    `json:"work_unit"`
	WorkUnitAddress *string    `json:"work_unit_address"`

	// Relationships
	Session *AttendanceSession `json:"session,omitempty"`
}

// TableName overrides the table name used by Attendance
func (Attendance) TableName() string {
	return "attendances"
}
