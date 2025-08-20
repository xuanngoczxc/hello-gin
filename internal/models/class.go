package models

import "time"

// Class represents a class in the system
type Class struct {
	ID        int       `json:"id" gorm:"primaryKey" example:"1"`
	ClassCode *string   `json:"class_code" example:"IT001"`
	ClassName *string   `json:"class_name" example:"Lap trinh Web"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	Students  []Student `json:"students,omitempty" gorm:"foreignKey:ClassID"`
}

// TableName overrides the table name used by Class to `Class`
func (Class) TableName() string {
	return "Class"
}
