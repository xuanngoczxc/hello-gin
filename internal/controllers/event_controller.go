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

// GetActiveEvents retrieves all active events
// @Summary Get all active events
// @Description Get a list of all active events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/active [get]
func (c *EventController) GetActiveEvents(ctx *gin.Context) {
	events, err := c.eventService.GetActiveEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve active events",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Active events retrieved successfully",
		"data":    events,
		"count":   len(events),
	})
}

// ToggleEventActive sets the active status of an event
// @Summary Set event active status
// @Description Set the active status of an event by its ID (1 = active, 0 = inactive)
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param active body object{active=int} true "Active status (1 = active, 0 = inactive)"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} map[string]interface{} "error"
// @Failure 404 {object} map[string]interface{} "error"
// @Failure 500 {object} map[string]interface{} "error"
// @Router /events/{id}/active [put]
func (c *EventController) EventActive(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"message": err.Error(),
		})
		return
	}

	// Parse request body to get active status
	var request struct {
		Active *int `json:"active"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body format",
			"message": err.Error(),
		})
		return
	}

	// Check if active field is provided
	if request.Active == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing required field 'active'",
			"message": "Request body must include 'active' field with value 0 or 1",
		})
		return
	}

	// Validate active value (must be 0 or 1)
	if *request.Active != 0 && *request.Active != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid active value. Must be 0 (inactive) or 1 (active)",
			"message": "Active value must be either 0 or 1",
		})
		return
	}

	// Convert to boolean
	isActive := *request.Active == 1

	event, err := c.eventService.SetEventActive(uint(id), isActive)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Event not found or failed to update active status",
			"message": err.Error(),
		})
		return
	}

	status := "deactivated"
	if event.IsActive != nil && *event.IsActive {
		status = "activated"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event " + status + " successfully",
		"data":    event,
	})
}
