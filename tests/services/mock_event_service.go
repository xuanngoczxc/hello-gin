package services

import (
	"hello-gin/internal/interfaces"
	"hello-gin/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockEventService is a mock implementation of EventServiceInterface
type MockEventService struct {
	mock.Mock
}

// Ensure MockEventService implements EventServiceInterface
var _ interfaces.EventServiceInterface = (*MockEventService)(nil)

func (m *MockEventService) GetAllEvents() ([]models.Event, error) {
	args := m.Called()
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventService) GetEventByID(id uint) (*models.Event, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventService) GetEventByIDWithSessions(id uint) (*models.Event, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventService) CreateEvent(req *models.CreateEventRequest) (*models.Event, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventService) UpdateEvent(id uint, req *models.CreateEventRequest) (*models.Event, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventService) DeleteEvent(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEventService) GetActiveEvents() ([]models.Event, error) {
	args := m.Called()
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventService) SetEventActive(id uint, isActive bool) (*models.Event, error) {
	args := m.Called(id, isActive)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}
