package main

import (
	"hello-gin/config"
	_ "hello-gin/docs"
	"hello-gin/internal/routes"
	"os"

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

	// Khởi tạo Gin
	r := gin.Default()

	// Đăng ký routes
	routes.RegisterRoutes(r)

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
