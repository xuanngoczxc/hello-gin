package models

import "time"

// Attendance represents an attendance record in the system
type Attendance struct {
	ID              int                `json:"id" gorm:"primaryKey" example:"1"`
	CreatedAt       time.Time          `json:"created_at" example:"2023-01-01T00:00:00Z"`
	SessionID       *int               `json:"session_id" example:"1"`
	CheckedInAt     *time.Time         `json:"checked_in_at" example:"2023-01-01T09:15:00Z"`
	StudentName     *string            `json:"student_name" example:"Nguyen Van A"`
	Email           *string            `json:"email" example:"student@example.com"`
	Phone           *string            `json:"phone" example:"0123456789"`
	WorkUnit        *string            `json:"work_unit" example:"Phong IT"`
	WorkUnitAddress *string            `json:"work_unit_address" example:"123 Nguyen Trai, Q1, HCM"`
	Session         *AttendanceSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
}

// TableName overrides the table name used by Attendance
func (Attendance) TableName() string {
	return "Attendance"
}
