package routes

import (
	"hello-gin/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, eventController *controllers.EventController) {
	api := r.Group("/api")
	{
		// Event routes
		api.GET("/events", eventController.GetEvents)
		api.GET("/events/:id", eventController.GetEventByID)
		api.GET("/events/:id/sessions", eventController.GetEventWithSessions)
		api.GET("/events/:id/attendances", controllers.GetAttendancesByEventID)
		api.POST("/events", eventController.CreateEvent)
		api.PUT("/events/:id", eventController.UpdateEvent)
		api.DELETE("/events/:id", eventController.DeleteEvent)

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
		api.POST("/attendance-sessions", controllers.CreateAttendanceSession)

		// Attendance routes
		api.GET("/attendances", controllers.GetAttendances)
		api.GET("/attendances/:id", controllers.GetAttendanceByID)
		api.POST("/attendances", controllers.CreateAttendance)
		api.GET("/sessions/:sessionId/attendances", controllers.GetAttendancesBySessionID)

		// Health check
		api.GET("/health", controllers.HealthCheck)
	}
}
