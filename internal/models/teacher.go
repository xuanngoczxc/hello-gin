package models

import (
	"time"

	"gorm.io/gorm"
)

// Teacher represents a teacher in the system
type Teacher struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	TeacherCode *string    `json:"teacher_code"`
	TeacherName *string    `json:"teacher_name"`
	Phone       *string    `json:"phone"`
	Email       *string    `json:"email"`
	WorkUnit    *string    `json:"work_unit"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	// Relationships
	Sessions []AttendanceSession `gorm:"foreignKey:TeacherID" json:"sessions,omitempty"`
}

// TableName sets the table name for Teacher model
func (Teacher) TableName() string {
	return "teachers"
}
