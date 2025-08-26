package controllers

import (
	"hello-gin/internal/models"
	"hello-gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetClasses godoc
// @Summary Get all classes
// @Description Get all classes from the database
// @Tags classes
// @Produce json
// @Success 200 {array} models.Class
// @Failure 500 {object} map[string]interface{}
// @Router /classes [get]
func GetClasses(c *gin.Context) {
	classes, err := services.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch classes",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    classes,
		"count":   len(classes),
	})
}

// GetClassByID godoc
// @Summary Get class by ID
// @Description Get a specific class with its students
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} models.Class
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /classes/{id} [get]
func GetClassByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid class ID",
			"message": "Class ID must be a number",
		})
		return
	}

	class, err := services.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Class not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    class,
	})
}

// CreateClass godoc
// @Summary Create a new class
// @Description Create a new class in the database
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.CreateClassRequest true "Class data"
// @Success 201 {object} models.Class
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /classes [post]
func CreateClass(c *gin.Context) {
	var request models.CreateClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"message": err.Error(),
		})
		return
	}

	// Create class from DTO
	class := models.Class{
		ClassCode: &request.ClassCode,
		ClassName: &request.ClassName,
	}

	if err := services.CreateClass(&class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create class",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    class,
		"message": "Class created successfully",
	})
}
