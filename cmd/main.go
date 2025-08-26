package main

import (
	"hello-gin/config"
	_ "hello-gin/docs"
	"hello-gin/internal/controllers"
	"hello-gin/internal/repository"
	"hello-gin/internal/routes"
	"hello-gin/internal/services"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Student Management API
// @version         1.0
// @description     This is a simple student management server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {
	// Kết nối DB
	config.ConnectDB()

	// Khởi tạo repositories
	eventRepo := repository.NewEventRepository(config.DB)

	// Khởi tạo services
	eventService := services.NewEventService(eventRepo)

	// Khởi tạo controllers với interface
	eventController := controllers.NewEventController(eventService)

	// Khởi tạo Gin
	r := gin.Default()

	// Cấu hình CORS - Cho phép Frontend truy cập
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // React default
			"http://localhost:3001", // React alternative
			"http://localhost:5173", // Vite default
			"http://localhost:4200", // Angular default
			"http://localhost:8080", // Vue default
			"http://127.0.0.1:3000",
			"http://127.0.0.1:5173",
			"http://127.0.0.1:4200",
			"http://127.0.0.1:8080",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "Accept", "Accept-Language", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Đăng ký routes
	routes.RegisterRoutes(r, eventController)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Lấy port từ env hoặc dùng mặc định 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	r.Run(":" + port)
}
