package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

func GetStudents() ([]models.Student, error) {
	return repository.GetAllStudents()
}

func CreateStudent(student *models.Student) error {
	return repository.CreateStudent(student)
}
