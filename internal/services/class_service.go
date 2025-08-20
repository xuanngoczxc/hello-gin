package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

func GetClasses() ([]models.Class, error) {
	return repository.GetAllClasses()
}

func GetClassByID(id int) (*models.Class, error) {
	return repository.GetClassByID(id)
}

func CreateClass(class *models.Class) error {
	return repository.CreateClass(class)
}
