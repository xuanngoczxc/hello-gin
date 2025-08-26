package controllers

import (
	"hello-gin/internal/models"
	"hello-gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTeachers godoc
// @Summary Get all teachers
// @Description Get all teachers from the database
// @Tags teachers
// @Produce json
// @Success 200 {array} models.Teacher
// @Failure 500 {object} map[string]interface{}
// @Router /teachers [get]
func GetTeachers(c *gin.Context) {
	teachers, err := services.GetTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch teachers",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teachers,
		"count":   len(teachers),
	})
}

// GetTeacherByID godoc
// @Summary Get teacher by ID
// @Description Get a specific teacher by ID
// @Tags teachers
// @Produce json
// @Param id path int true "Teacher ID"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /teachers/{id} [get]
func GetTeacherByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid teacher ID",
			"message": "Teacher ID must be a number",
		})
		return
	}

	teacher, err := services.GetTeacherByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Teacher not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teacher,
	})
}

// CreateTeacher godoc
// @Summary Create a new teacher
// @Description Create a new teacher in the database
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher body models.CreateTeacherRequest true "Teacher data"
// @Success 201 {object} models.Teacher
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /teachers [post]
func CreateTeacher(c *gin.Context) {
	var req models.CreateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"message": err.Error(),
		})
		return
	}

	// Convert DTO to Teacher model
	teacher := models.Teacher{
		TeacherCode: req.TeacherCode,
		TeacherName: req.TeacherName,
		Phone:       req.Phone,
		Email:       req.Email,
		WorkUnit:    req.WorkUnit,
		DateOfBirth: req.DateOfBirth,
	}

	if err := services.CreateTeacher(&teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create teacher",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    teacher,
		"message": "Teacher created successfully",
	})
}
