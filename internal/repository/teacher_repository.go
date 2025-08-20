package repository

import (
	"hello-gin/config"
	"hello-gin/internal/models"
)

func GetAllTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher
	result := config.DB.Find(&teachers)
	return teachers, result.Error
}

func GetTeacherByID(id int) (*models.Teacher, error) {
	var teacher models.Teacher
	result := config.DB.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

func CreateTeacher(teacher *models.Teacher) error {
	result := config.DB.Create(teacher)
	return result.Error
}
