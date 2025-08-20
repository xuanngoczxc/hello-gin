package controllers

import (
	"hello-gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAttendanceSessions godoc
// @Summary Get all attendance sessions
// @Description Get all attendance sessions with class and teacher information
// @Tags attendance-sessions
// @Produce json
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
