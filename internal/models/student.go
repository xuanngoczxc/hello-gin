package models

import (
	"time"

	"gorm.io/gorm"
)

// Student represents a student in the system
type Student struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	StudentCode *string    `json:"student_code"`
	StudentName *string    `json:"student_name"`
	ClassID     *uint      `json:"class_id"`
	Phone       *string    `json:"phone"`
	Email       *string    `json:"email"`
	WorkUnit    *string    `json:"work_unit"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	// Relationships
	Class *Class `json:"class,omitempty"`
}

// TableName sets the table name for Student model
func (Student) TableName() string {
	return "students"
}
