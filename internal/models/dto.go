package models

// CreateStudentRequest represents the data needed to create a new student
type CreateStudentRequest struct {
	StudentCode *string `json:"student_code" example:"SV001"`
	StudentName *string `json:"student_name" example:"Nguyen Van A"`
	ClassID     *int    `json:"class_id" example:"1"`
	Email       *string `json:"email" example:"user@example.com"`
}

// CreateTeacherRequest represents the data needed to create a new teacher
type CreateTeacherRequest struct {
	TeacherCode *string `json:"teacher_code" example:"GV001"`
	TeacherName *string `json:"teacher_name" example:"Nguyen Thi B"`
}
