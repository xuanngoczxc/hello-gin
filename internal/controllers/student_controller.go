package controllers

import (
	"hello-gin/config"
	"hello-gin/internal/models"
	"hello-gin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStudents godoc
// @Summary      Get all students
// @Description  Get a list of all students
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Student
// @Failure      500  {object}  map[string]interface{}
// @Router       /students [get]
func GetStudents(c *gin.Context) {
	students, err := services.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// CreateStudent godoc
// @Summary      Create a new student
// @Description  Create a new student in the database
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body models.CreateStudentRequest true "Student data"
// @Success      201  {object}  models.Student
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
	var req models.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"message": err.Error(),
		})
		return
	}

	// Convert DTO to Student model
	student := models.Student{
		StudentCode: req.StudentCode,
		StudentName: req.StudentName,
		ClassID:     req.ClassID,
		Phone:       req.Phone,
		Email:       req.Email,
		WorkUnit:    req.WorkUnit,
		DateOfBirth: req.DateOfBirth,
	}

	if err := services.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create student",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    student,
		"message": "Student created successfully",
	})
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if the application and database are running
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	// Kiểm tra kết nối database
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Cannot get database instance",
			"error":   err.Error(),
		})
		return
	}

	// Ping database
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"message":  "✅ Database connected successfully!",
		"database": "PostgreSQL",
	})
}
