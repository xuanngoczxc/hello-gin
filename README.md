# Hello Gin API

## Setup Environment

1. Copy file `.env.example` thành `.env`:
   ```bash
   cp .env.example .env
   ```

2. Cập nhật thông tin trong file `.env`:
   ```env
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_actual_password
   DB_NAME=your_database_name
   DB_PORT=5432
   PORT=8080
   GIN_MODE=debug
   JWT_SECRET=your-actual-secret-key
   ```

## Required Environment Variables

- `DB_PASSWORD`: Database password (bắt buộc)
- `DB_NAME`: Database name (bắt buộc)

## Run Application

```bash
go run cmd/main.go
```

## API Endpoints

- `GET /api/health` - Check database connection
- `GET /api/students` - Get all students

## Security Notes

- ⚠️ **NEVER** commit file `.env` to git
- ✅ File `.env` đã được thêm vào `.gitignore`
- ✅ Chỉ share file `.env.example` cho team members
