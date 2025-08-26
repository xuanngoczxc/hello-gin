package models

import (
	"time"

	"gorm.io/gorm"
)

// Class represents a class in the system
type Class struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ClassCode *string `json:"class_code"`
	ClassName *string `json:"class_name"`

	// Relationships
	Students []Student           `gorm:"foreignKey:ClassID" json:"students,omitempty"`
	Sessions []AttendanceSession `gorm:"foreignKey:ClassID" json:"sessions,omitempty"`
}

// TableName sets the table name for Class model
func (Class) TableName() string {
	return "classes"
}
