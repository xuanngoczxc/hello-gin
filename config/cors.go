package config

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS cấu hình CORS cho phép Frontend truy cập
func SetupCORS() gin.HandlerFunc {
	// Get environment
	env := os.Getenv("GIN_MODE")

	if env == "debug" || env == "test" || env == "" {
		// Development: Allow more permissive CORS
		return cors.New(cors.Config{
			AllowOrigins: []string{
				"http://localhost:3000", // React default
				"http://localhost:3001", // React alternative
				"http://localhost:5173", // Vite default
				"http://localhost:4200", // Angular default
				"http://localhost:8080", // Vue default
				"http://localhost:8081", // Alternative port
				"http://127.0.0.1:3000",
				"http://127.0.0.1:5173",
				"http://127.0.0.1:4200",
				"http://127.0.0.1:8080",
				"http://127.0.0.1:8081",
			},
			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"PATCH",
				"DELETE",
				"HEAD",
				"OPTIONS",
			},
			AllowHeaders: []string{
				"Origin",
				"Content-Length",
				"Content-Type",
				"Authorization",
				"X-Requested-With",
				"Accept",
				"Accept-Language",
				"Accept-Encoding",
				"X-CSRF-Token",
				"X-Auth-Token",
				"Cache-Control",
			},
			ExposeHeaders: []string{
				"Content-Length",
				"Content-Type",
				"Cache-Control",
				"Content-Language",
				"Content-Location",
				"Expires",
				"Last-Modified",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
	}

	// Production: Restrict origins
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	var origins []string

	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		// Trim spaces
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
	} else {
		// Default production origins - Update these with your actual domains
		origins = []string{
			"https://yourdomain.com",
			"https://www.yourdomain.com",
		}
	}

	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}

// SetupCORSForDevelopment - CORS config cho development (cho phép tất cả)
func SetupCORSForDevelopment() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
