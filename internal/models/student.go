package models

import "time"

// Student represents a student in the system
type Student struct {
	ID          int        `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt   time.Time  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	StudentCode *string    `json:"student_code" example:"SV001"`
	StudentName *string    `json:"student_name" example:"Nguyen Van A"`
	ClassID     *int       `json:"class_id" example:"1"`
	Phone       *string    `json:"phone" example:"0123456789"`
	Email       *string    `json:"email" example:"student@example.com"`
	SchoolName  *string    `json:"school_name" example:"Dai hoc Bach Khoa"`
	DateOfBirth *time.Time `json:"date_of_birth" example:"2000-01-01T00:00:00Z"`
	Class       *Class     `json:"class,omitempty" gorm:"foreignKey:ClassID"`
}

// TableName overrides the table name used by Student to `Student`
func (Student) TableName() string {
	return "Student"
}
