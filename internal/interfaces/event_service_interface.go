package interfaces

import "hello-gin/internal/models"

// EventServiceInterface định nghĩa interface cho EventService
type EventServiceInterface interface {
	GetAllEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	GetEventByIDWithSessions(id uint) (*models.Event, error)
	CreateEvent(req *models.CreateEventRequest) (*models.Event, error)
	UpdateEvent(id uint, req *models.CreateEventRequest) (*models.Event, error)
	DeleteEvent(id uint) error
}
