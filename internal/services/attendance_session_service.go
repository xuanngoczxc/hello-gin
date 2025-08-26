package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

func GetAttendanceSessions() ([]models.AttendanceSession, error) {
	return repository.GetAllAttendanceSessions()
}

func GetAttendanceSessionByID(id int) (*models.AttendanceSession, error) {
	return repository.GetAttendanceSessionByID(id)
}

func GetAttendanceSessionsByEventID(eventID uint) ([]models.AttendanceSession, error) {
	return repository.GetAttendanceSessionsByEventID(eventID)
}

func CreateAttendanceSession(session *models.AttendanceSession) error {
	return repository.CreateAttendanceSession(session)
}
