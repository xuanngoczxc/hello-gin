package models

import "time"

// CreateStudentRequest represents the data needed to create a new student
type CreateStudentRequest struct {
	StudentCode *string    `json:"student_code" example:"SV001"`
	StudentName *string    `json:"student_name" example:"Nguyen Van A"`
	ClassID     *uint      `json:"class_id" example:"1"`
	Phone       *string    `json:"phone" example:"0123456789"`
	Email       *string    `json:"email" example:"user@example.com"`
	WorkUnit    *string    `json:"work_unit" example:"Công ty ABC"`
	DateOfBirth *time.Time `json:"date_of_birth" example:"2000-01-01T00:00:00Z"`
}

// CreateTeacherRequest represents the data needed to create a new teacher
type CreateTeacherRequest struct {
	TeacherCode *string    `json:"teacher_code" example:"GV001"`
	TeacherName *string    `json:"teacher_name" example:"Nguyen Thi B"`
	Phone       *string    `json:"phone" example:"0123456789"`
	Email       *string    `json:"email" example:"teacher@example.com"`
	WorkUnit    *string    `json:"work_unit" example:"Trường Đại học ABC"`
	DateOfBirth *time.Time `json:"date_of_birth" example:"1980-01-01T00:00:00Z"`
}

// CreateAttendanceSessionRequest represents the data needed to create a new attendance session
type CreateAttendanceSessionRequest struct {
	EventID     *uint   `json:"event_id" example:"1"`
	ClassID     *uint   `json:"class_id" example:"1"`
	TeacherID   *string `json:"teacher_id,omitempty" example:"1"`
	SessionDate *string `json:"session_date" example:"2025-08-20T08:31:46.121Z"`
}

// CreateEventRequest represents the data needed to create a new event
type CreateEventRequest struct {
	EventName   *string    `json:"event_name" example:"Workshop AI"`
	Description *string    `json:"description" example:"Workshop về trí tuệ nhân tạo"`
	StartDate   *time.Time `json:"start_date" example:"2023-01-01T00:00:00Z"`
}

// CreateClassRequest represents the data needed to create a new class
type CreateClassRequest struct {
	ClassCode string `json:"class_code" binding:"required" example:"LOP001"`
	ClassName string `json:"class_name" binding:"required" example:"Lớp Khoa học máy tính K65"`
}

// CreateAttendanceRequest represents the data needed to create a new attendance
type CreateAttendanceRequest struct {
	SessionID       uint   `json:"session_id" binding:"required" example:"1"`
	StudentName     string `json:"student_name" binding:"required" example:"Nguyen Van A"`
	Email           string `json:"email" binding:"required" example:"student@example.com"`
	Phone           string `json:"phone" binding:"required" example:"0123456789"`
	WorkUnit        string `json:"work_unit" binding:"required" example:"Công ty ABC"`
	WorkUnitAddress string `json:"work_unit_address" binding:"required" example:"123 Đường ABC, Quận 1, TP.HCM"`
}
