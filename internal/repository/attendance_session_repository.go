package repository

import (
	"hello-gin/config"
	"hello-gin/internal/models"
)

func GetAllAttendanceSessions() ([]models.AttendanceSession, error) {
	var sessions []models.AttendanceSession
	result := config.DB.Preload("Class").Preload("Teacher").Preload("Attendances").Find(&sessions)
	return sessions, result.Error
}

func GetAttendanceSessionByID(id int) (*models.AttendanceSession, error) {
	var session models.AttendanceSession
	result := config.DB.Preload("Class").Preload("Teacher").Preload("Attendances").First(&session, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func CreateAttendanceSession(session *models.AttendanceSession) error {
	result := config.DB.Create(session)
	return result.Error
}
