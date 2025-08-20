package repository

import (
	"hello-gin/config"
	"hello-gin/internal/models"
)

func GetAllClasses() ([]models.Class, error) {
	var classes []models.Class
	result := config.DB.Find(&classes)
	return classes, result.Error
}

func GetClassByID(id int) (*models.Class, error) {
	var class models.Class
	result := config.DB.Preload("Students").First(&class, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

func CreateClass(class *models.Class) error {
	result := config.DB.Create(class)
	return result.Error
}
