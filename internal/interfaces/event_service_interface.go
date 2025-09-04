package interfaces

import "hello-gin/internal/models"

type EventServiceInterface interface {
	GetAllEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	GetEventByIDWithSessions(id uint) (*models.Event, error)
	CreateEvent(req *models.CreateEventRequest) (*models.Event, error)
	UpdateEvent(id uint, req *models.CreateEventRequest) (*models.Event, error)
	DeleteEvent(id uint) error

	GetActiveEvents() ([]models.Event, error)
	SetEventActive(id uint, isActive bool) (*models.Event, error)
}
