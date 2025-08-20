package controllers

import (
	"hello-gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAttendances godoc
// @Summary Get all attendances
// @Description Get all attendance records with session information
// @Tags attendances
// @Produce json
// @Success 200 {array} models.Attendance
// @Failure 500 {object} map[string]interface{}
// @Router /attendances [get]
func GetAttendances(c *gin.Context) {
	attendances, err := services.GetAttendances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch attendances",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    attendances,
		"count":   len(attendances),
	})
}

// GetAttendanceByID godoc
// @Summary Get attendance by ID
// @Description Get a specific attendance record with session information
// @Tags attendances
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} models.Attendance
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendances/{id} [get]
func GetAttendanceByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid attendance ID",
			"message": "Attendance ID must be a number",
		})
		return
	}

	attendance, err := services.GetAttendanceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Attendance not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    attendance,
	})
}

// GetAttendancesBySessionID godoc
// @Summary Get attendances by session ID
// @Description Get all attendance records for a specific session
// @Tags attendances
// @Produce json
// @Param sessionId path int true "Session ID"
// @Success 200 {array} models.Attendance
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sessions/{sessionId}/attendances [get]
func GetAttendancesBySessionID(c *gin.Context) {
	sessionIdParam := c.Param("sessionId")
	sessionId, err := strconv.Atoi(sessionIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid session ID",
			"message": "Session ID must be a number",
		})
		return
	}

	attendances, err := services.GetAttendancesBySessionID(sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch attendances",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    attendances,
		"count":   len(attendances),
	})
}
