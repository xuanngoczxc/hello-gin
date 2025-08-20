package repository

import (
	"hello-gin/config"
	"hello-gin/internal/models"
)

func GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	result := config.DB.Find(&students)
	return students, result.Error
}

func CreateStudent(student *models.Student) error {
	result := config.DB.Create(student)
	return result.Error
}
