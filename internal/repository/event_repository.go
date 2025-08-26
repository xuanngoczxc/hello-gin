package repository

import (
	"hello-gin/internal/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

// GetAll retrieves all events
func (r *EventRepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

// GetByID retrieves an event by ID
func (r *EventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// GetByIDWithSessions retrieves an event by ID with its sessions
func (r *EventRepository) GetByIDWithSessions(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.Preload("Sessions").Preload("Sessions.Class").Preload("Sessions.Teacher").First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Create creates a new event
func (r *EventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

// Update updates an existing event
func (r *EventRepository) Update(event *models.Event) error {
	return r.db.Save(event).Error
}

// Delete deletes an event by ID
func (r *EventRepository) Delete(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}

// GetActiveEvents retrieves all active events
func (r *EventRepository) GetActiveEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Where("is_active = ?", true).Find(&events).Error
	return events, err
}
