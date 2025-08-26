package controllers

import (
	"hello-gin/internal/models"
	"hello-gin/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAttendanceSessions godoc
// @Summary Get all attendance sessions
// @Description Get all attendance sessions with class and teacher information
// @Tags attendance-sessions
// @Produce json
// @Param event_id query int false "Filter by Event ID"
// @Success 200 {array} models.AttendanceSession
// @Failure 500 {object} map[string]interface{}
// @Router /attendance-sessions [get]
func GetAttendanceSessions(c *gin.Context) {
	sessions, err := services.GetAttendanceSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch attendance sessions",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    sessions,
		"count":   len(sessions),
	})
}

// GetAttendanceSessionByID godoc
// @Summary Get attendance session by ID
// @Description Get a specific attendance session with class, teacher, and attendances
// @Tags attendance-sessions
// @Produce json
// @Param id path int true "Attendance Session ID"
// @Success 200 {object} models.AttendanceSession
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance-sessions/{id} [get]
func GetAttendanceSessionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid attendance session ID",
			"message": "Attendance session ID must be a number",
		})
		return
	}

	session, err := services.GetAttendanceSessionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Attendance session not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    session,
	})
}

// CreateAttendanceSession godoc
// @Summary Create a new attendance session
// @Description Create a new attendance session in the database
// @Tags attendance-sessions
// @Accept json
// @Produce json
// @Param session body models.CreateAttendanceSessionRequest true "Attendance session data"
// @Success 201 {object} models.AttendanceSession
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance-sessions [post]
func CreateAttendanceSession(c *gin.Context) {
	var req models.CreateAttendanceSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"message": err.Error(),
		})
		return
	}

	// Convert DTO to AttendanceSession model
	// TeacherID is optional - can be null if not provided or empty string
	var teacherID *uint
	if req.TeacherID != nil && *req.TeacherID != "" {
		if id, err := strconv.ParseUint(*req.TeacherID, 10, 32); err == nil {
			teacherIDUint := uint(id)
			teacherID = &teacherIDUint
		}
	}

	// Parse session date
	var sessionDate *time.Time
	if req.SessionDate != nil && *req.SessionDate != "" {
		if parsedTime, err := time.Parse(time.RFC3339, *req.SessionDate); err == nil {
			sessionDate = &parsedTime
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid session date format",
				"message": "Please use RFC3339 format (2006-01-02T15:04:05Z07:00)",
			})
			return
		}
	}

	session := models.AttendanceSession{
		EventID:     req.EventID,
		ClassID:     req.ClassID,
		TeacherID:   teacherID,
		SessionDate: sessionDate,
	}

	if err := services.CreateAttendanceSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create attendance session",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    session,
		"message": "Attendance session created successfully",
	})
}
