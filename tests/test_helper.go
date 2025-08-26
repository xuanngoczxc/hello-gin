package tests

import (
	"hello-gin/internal/models"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupTestGin initializes Gin in test mode
func SetupTestGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

// SetupMockDB creates a mock database for testing
func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// CreateSampleEvent creates a sample event for testing
func CreateSampleEvent() *models.Event {
	eventName := "Test Event"
	description := "Test Description"
	startDate := time.Now()

	return &models.Event{
		ID:          1,
		EventName:   &eventName,
		Description: &description,
		StartDate:   &startDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// CreateSampleCreateEventRequest creates a sample request for testing
func CreateSampleCreateEventRequest() *models.CreateEventRequest {
	eventName := "Test Event"
	description := "Test Description"
	startDate := time.Now()

	return &models.CreateEventRequest{
		EventName:   &eventName,
		Description: &description,
		StartDate:   &startDate,
	}
}
