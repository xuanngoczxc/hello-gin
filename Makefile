# Makefile for Windows PowerShell
.PHONY: dev build docs clean

# Development with auto swagger generation
dev:
	@echo "🔄 Generating Swagger docs..."
	@swag init -g cmd/main.go
	@echo "✅ Swagger docs generated!"
	@echo "🚀 Starting development server..."
	@go run cmd/main.go

# Build the application
build:
	@echo "🔄 Generating Swagger docs..."
	@swag init -g cmd/main.go
	@echo "🏗️ Building application..."
	@go build -o bin/main.exe cmd/main.go

# Generate swagger docs only
docs:
	@echo "📝 Generating Swagger documentation..."
	@swag init -g cmd/main.go
	@echo "✅ Documentation generated!"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@if exist bin rmdir /s /q bin
	@if exist tmp rmdir /s /q tmp
	@echo "✅ Clean complete!"

# Install development dependencies
install:
	@echo "📦 Installing dependencies..."
	@go mod tidy
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/air-verse/air@latest
	@echo "✅ Dependencies installed!"

# Run with hot reload using Air
watch:
	@echo "👁️ Starting with hot reload..."
	@air
