package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

func GetAttendances() ([]models.Attendance, error) {
	return repository.GetAllAttendances()
}

func GetAttendanceByID(id int) (*models.Attendance, error) {
	return repository.GetAttendanceByID(id)
}

func GetAttendancesBySessionID(sessionID int) ([]models.Attendance, error) {
	return repository.GetAttendancesBySessionID(sessionID)
}

func GetAttendancesByEventID(eventID uint) ([]models.Attendance, error) {
	return repository.GetAttendancesByEventID(eventID)
}

func CreateAttendance(attendance *models.Attendance) error {
	return repository.CreateAttendance(attendance)
}
