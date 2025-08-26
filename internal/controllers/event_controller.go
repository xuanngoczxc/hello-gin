package controllers

import (
	"hello-gin/internal/interfaces"
	"hello-gin/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService interfaces.EventServiceInterface
}

func NewEventController(eventService interfaces.EventServiceInterface) *EventController {
	return &EventController{
		eventService: eventService,
	}
}

// GetEvents retrieves all events
// @Summary Get all events
// @Description Get a list of all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events [get]
func (c *EventController) GetEvents(ctx *gin.Context) {
	events, err := c.eventService.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve events",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Events retrieved successfully",
		"data":    events,
		"count":   len(events),
	})
}

// GetEventByID retrieves an event by ID
// @Summary Get event by ID
// @Description Get a single event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 404 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/{id} [get]
func (c *EventController) GetEventByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"message": err.Error(),
		})
		return
	}

	event, err := c.eventService.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Event not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event retrieved successfully",
		"data":    event,
	})
}

// GetEventWithSessions retrieves an event by ID with its sessions
// @Summary Get event with sessions
// @Description Get a single event by its ID including all attendance sessions
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 404 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/{id}/sessions [get]
func (c *EventController) GetEventWithSessions(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"message": err.Error(),
		})
		return
	}

	event, err := c.eventService.GetEventByIDWithSessions(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Event not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event with sessions retrieved successfully",
		"data":    event,
	})
}

// CreateEvent creates a new event
// @Summary Create a new event
// @Description Create a new event with the provided information
// @Tags events
// @Accept json
// @Produce json
// @Param event body models.CreateEventRequest true "Event data"
// @Success 201 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events [post]
func (c *EventController) CreateEvent(ctx *gin.Context) {
	var req models.CreateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	event, err := c.eventService.CreateEvent(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create event",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}

// UpdateEvent updates an existing event
// @Summary Update an event
// @Description Update an existing event with the provided information
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body models.CreateEventRequest true "Event data"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 404 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/{id} [put]
func (c *EventController) UpdateEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"message": err.Error(),
		})
		return
	}

	var req models.CreateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	event, err := c.eventService.UpdateEvent(uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update event",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"data":    event,
	})
}

// DeleteEvent deletes an event
// @Summary Delete an event
// @Description Delete an event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/{id} [delete]
func (c *EventController) DeleteEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"message": err.Error(),
		})
		return
	}

	err = c.eventService.DeleteEvent(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete event",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
