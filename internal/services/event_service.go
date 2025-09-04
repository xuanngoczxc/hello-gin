package services

import (
	"hello-gin/internal/models"
	"hello-gin/internal/repository"
)

type EventService struct {
	eventRepo *repository.EventRepository
}

func NewEventService(eventRepo *repository.EventRepository) *EventService {
	return &EventService{
		eventRepo: eventRepo,
	}
}

// GetAllEvents retrieves all events
func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.eventRepo.GetAll()
}

// GetEventByID retrieves an event by ID
func (s *EventService) GetEventByID(id uint) (*models.Event, error) {
	return s.eventRepo.GetByID(id)
}

// GetEventByIDWithSessions retrieves an event by ID with its sessions
func (s *EventService) GetEventByIDWithSessions(id uint) (*models.Event, error) {
	return s.eventRepo.GetByIDWithSessions(id)
}

// CreateEvent creates a new event
func (s *EventService) CreateEvent(req *models.CreateEventRequest) (*models.Event, error) {
	event := &models.Event{
		EventName:   req.EventName,
		Description: req.Description,
		StartDate:   req.StartDate,
	}

	err := s.eventRepo.Create(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// UpdateEvent updates an existing event
func (s *EventService) UpdateEvent(id uint, req *models.CreateEventRequest) (*models.Event, error) {
	event, err := s.eventRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.EventName != nil {
		event.EventName = req.EventName
	}
	if req.Description != nil {
		event.Description = req.Description
	}
	if req.StartDate != nil {
		event.StartDate = req.StartDate
	}

	err = s.eventRepo.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// DeleteEvent deletes an event by ID
func (s *EventService) DeleteEvent(id uint) error {
	return s.eventRepo.Delete(id)
}

// GetActiveEvents retrieves all active events
func (s *EventService) GetActiveEvents() ([]models.Event, error) {
	return s.eventRepo.GetActiveEvents()
}

// SetEventActive sets the active status of an event
func (s *EventService) SetEventActive(id uint, isActive bool) (*models.Event, error) {
	event, err := s.eventRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Set the specific active status
	event.IsActive = &isActive

	err = s.eventRepo.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
