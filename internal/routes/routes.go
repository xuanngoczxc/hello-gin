package routes

import (
	"hello-gin/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Student routes
		api.GET("/students", controllers.GetStudents)
		api.POST("/students", controllers.CreateStudent)

		// Class routes
		api.GET("/classes", controllers.GetClasses)
		api.GET("/classes/:id", controllers.GetClassByID)
		api.POST("/classes", controllers.CreateClass)

		// Teacher routes
		api.GET("/teachers", controllers.GetTeachers)
		api.GET("/teachers/:id", controllers.GetTeacherByID)
		api.POST("/teachers", controllers.CreateTeacher)

		// Attendance Session routes
		api.GET("/attendance-sessions", controllers.GetAttendanceSessions)
		api.GET("/attendance-sessions/:id", controllers.GetAttendanceSessionByID)

		// Attendance routes
		api.GET("/attendances", controllers.GetAttendances)
		api.GET("/attendances/:id", controllers.GetAttendanceByID)
		api.GET("/sessions/:sessionId/attendances", controllers.GetAttendancesBySessionID)

		// Health check
		api.GET("/health", controllers.HealthCheck)
	}
}
