#!/usr/bin/env pwsh

Write-Host "🔄 Auto-generating Swagger docs..." -ForegroundColor Yellow
swag init -g cmd/main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Swagger docs generated successfully!" -ForegroundColor Green
    Write-Host "🚀 Starting server..." -ForegroundColor Blue
    go run cmd/main.go
} else {
    Write-Host "❌ Failed to generate Swagger docs" -ForegroundColor Red
    exit $LASTEXITCODE
}
