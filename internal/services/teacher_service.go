package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

func GetTeachers() ([]models.Teacher, error) {
	return repository.GetAllTeachers()
}

func GetTeacherByID(id int) (*models.Teacher, error) {
	return repository.GetTeacherByID(id)
}

func CreateTeacher(teacher *models.Teacher) error {
	return repository.CreateTeacher(teacher)
}
