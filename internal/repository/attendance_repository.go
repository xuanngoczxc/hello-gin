package repository

import (
	"hello-gin/config"
	"hello-gin/internal/models"
)

func GetAllAttendances() ([]models.Attendance, error) {
	var attendances []models.Attendance
	result := config.DB.Preload("Session").Preload("Session.Class").Preload("Session.Teacher").Find(&attendances)
	return attendances, result.Error
}

func GetAttendanceByID(id int) (*models.Attendance, error) {
	var attendance models.Attendance
	result := config.DB.Preload("Session").Preload("Session.Class").Preload("Session.Teacher").First(&attendance, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &attendance, nil
}

func GetAttendancesBySessionID(sessionID int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	result := config.DB.Where("session_id = ?", sessionID).Find(&attendances)
	return attendances, result.Error
}

func GetAttendancesByEventID(eventID uint) ([]models.Attendance, error) {
	var attendances []models.Attendance
	result := config.DB.
		Preload("Session").
		Preload("Session.Class").
		Preload("Session.Teacher").
		Preload("Session.Event").
		Joins("JOIN attendance_sessions ON attendances.session_id = attendance_sessions.id").
		Where("attendance_sessions.event_id = ?", eventID).
		Find(&attendances)
	return attendances, result.Error
}

func CreateAttendance(attendance *models.Attendance) error {
	result := config.DB.Create(attendance)
	return result.Error
}
