package models

import (
	"time"

	"gorm.io/gorm"
)

// Event represents an attendance event (e.g., workshop, seminar, course)
type Event struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	EventName   *string    `json:"event_name"`
	Description *string    `json:"description"`
	StartDate   *time.Time `json:"start_date"`

	// Relationships
	Sessions []AttendanceSession `gorm:"foreignKey:EventID" json:"sessions,omitempty"`
}

// TableName sets the table name for Event model
func (Event) TableName() string {
	return "events"
}
