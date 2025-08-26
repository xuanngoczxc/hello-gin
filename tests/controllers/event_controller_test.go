package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"hello-gin/internal/controllers"
	"hello-gin/internal/models"
	"hello-gin/tests"
	mockServices "hello-gin/tests/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEvents_Success(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Create sample events
	events := []models.Event{*tests.CreateSampleEvent()}

	// Setup mock expectations
	mockService.On("GetAllEvents").Return(events, nil)

	// Setup Gin
	r := tests.SetupTestGin()
	r.GET("/events", controller.GetEvents)

	// Create request
	req, _ := http.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Events retrieved successfully", response["message"])
	assert.Equal(t, float64(1), response["count"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestGetEvents_ServiceError(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Setup mock expectations
	mockService.On("GetAllEvents").Return([]models.Event{}, errors.New("database error"))

	// Setup Gin
	r := tests.SetupTestGin()
	r.GET("/events", controller.GetEvents)

	// Create request
	req, _ := http.NewRequest("GET", "/events", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Failed to retrieve events", response["error"])
	assert.Equal(t, "database error", response["message"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestGetEventByID_Success(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Create sample event
	event := tests.CreateSampleEvent()

	// Setup mock expectations
	mockService.On("GetEventByID", uint(1)).Return(event, nil)

	// Setup Gin
	r := tests.SetupTestGin()
	r.GET("/events/:id", controller.GetEventByID)

	// Create request
	req, _ := http.NewRequest("GET", "/events/1", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Event retrieved successfully", response["message"])
	assert.NotNil(t, response["data"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestGetEventByID_InvalidID(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Setup Gin
	r := tests.SetupTestGin()
	r.GET("/events/:id", controller.GetEventByID)

	// Create request with invalid ID
	req, _ := http.NewRequest("GET", "/events/invalid", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid event ID", response["error"])
}

func TestGetEventByID_NotFound(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Setup mock expectations
	mockService.On("GetEventByID", uint(999)).Return((*models.Event)(nil), errors.New("event not found"))

	// Setup Gin
	r := tests.SetupTestGin()
	r.GET("/events/:id", controller.GetEventByID)

	// Create request
	req, _ := http.NewRequest("GET", "/events/999", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Event not found", response["error"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestCreateEvent_Success(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Create sample request and response
	request := tests.CreateSampleCreateEventRequest()
	event := tests.CreateSampleEvent()

	// Setup mock expectations
	mockService.On("CreateEvent", mock.AnythingOfType("*models.CreateEventRequest")).Return(event, nil)

	// Setup Gin
	r := tests.SetupTestGin()
	r.POST("/events", controller.CreateEvent)

	// Create request body
	requestBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Event created successfully", response["message"])
	assert.NotNil(t, response["data"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestCreateEvent_InvalidJSON(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Setup Gin
	r := tests.SetupTestGin()
	r.POST("/events", controller.CreateEvent)

	// Create request with invalid JSON
	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid request body", response["error"])
}

func TestCreateEvent_ServiceError(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Create sample request
	request := tests.CreateSampleCreateEventRequest()

	// Setup mock expectations
	mockService.On("CreateEvent", mock.AnythingOfType("*models.CreateEventRequest")).Return((*models.Event)(nil), errors.New("database error"))

	// Setup Gin
	r := tests.SetupTestGin()
	r.POST("/events", controller.CreateEvent)

	// Create request body
	requestBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Failed to create event", response["error"])
	assert.Equal(t, "database error", response["message"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}

func TestDeleteEvent_Success(t *testing.T) {
	// Setup
	mockService := new(mockServices.MockEventService)
	controller := controllers.NewEventController(mockService)

	// Setup mock expectations
	mockService.On("DeleteEvent", uint(1)).Return(nil)

	// Setup Gin
	r := tests.SetupTestGin()
	r.DELETE("/events/:id", controller.DeleteEvent)

	// Create request
	req, _ := http.NewRequest("DELETE", "/events/1", nil)
	w := httptest.NewRecorder()

	// Execute
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Event deleted successfully", response["message"])

	// Verify mock expectations
	mockService.AssertExpectations(t)
}
