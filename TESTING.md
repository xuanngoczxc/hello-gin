# 🧪 Unit Testing Guide cho Go Gin API

## 📋 Tổng quan

Project này sử dụng Go testing framework kết hợp với các thư viện bổ trợ để viết unit tests cho Gin REST APIs.

## 🛠️ Công cụ Testing

### Dependencies
```bash
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/mock
go get github.com/DATA-DOG/go-sqlmock
go get github.com/gin-gonic/gin
```

### Thư viện sử dụng:
- **testify/assert**: Assertions cho tests
- **testify/mock**: Mock objects cho dependency injection
- **go-sqlmock**: Mock SQL database
- **net/http/httptest**: Test HTTP requests/responses
- **gin.CreateTestContext()**: Gin test context

## 📁 Cấu trúc Test

```
tests/
├── test_helper.go                    # Helper functions chung
├── controllers/
│   └── event_controller_test.go      # Test cho Event API
└── services/
    └── mock_event_service.go         # Mock service implementations
```

## 🎯 Pattern Testing cho Gin Controllers

### 1. Setup Test Structure

```go
package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestControllerMethod_Scenario(t *testing.T) {
    // 1. Setup mock service
    mockService := new(MockEventService)
    controller := controllers.NewEventController(mockService)
    
    // 2. Setup mock expectations
    mockService.On("MethodName", args).Return(data, nil)
    
    // 3. Setup Gin router
    r := tests.SetupTestGin()
    r.POST("/events", controller.CreateEvent)
    
    // 4. Create HTTP request
    req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(requestBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    
    // 5. Execute request
    r.ServeHTTP(w, req)
    
    // 6. Assert results
    assert.Equal(t, http.StatusCreated, w.Code)
    
    // 7. Verify mock expectations
    mockService.AssertExpectations(t)
}
```

### 2. Test Helper Functions

```go
// tests/test_helper.go
func SetupTestGin() *gin.Engine {
    gin.SetMode(gin.TestMode)
    return gin.New()
}

func CreateSampleEvent() *models.Event {
    return &models.Event{
        ID:          1,
        EventName:   "Sample Event",
        Description: "Sample Description",
        ClassID:     1,
    }
}
```

### 3. Mock Service Implementation

```go
// tests/services/mock_event_service.go
type MockEventService struct {
    mock.Mock
}

func (m *MockEventService) GetAllEvents() ([]models.Event, error) {
    args := m.Called()
    return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventService) CreateEvent(req *models.CreateEventRequest) (*models.Event, error) {
    args := m.Called(req)
    return args.Get(0).(*models.Event), args.Error(1)
}
```

## 🧪 Các Test Cases đã implement

### Event Controller Tests

#### 1. GET /events
- ✅ **Success**: Trả về danh sách events
- ✅ **Service Error**: Xử lý lỗi database

#### 2. GET /events/:id  
- ✅ **Success**: Trả về event theo ID
- ✅ **Invalid ID**: ID không hợp lệ (non-numeric)
- ✅ **Not Found**: Event không tồn tại

#### 3. POST /events
- ✅ **Success**: Tạo event thành công
- ✅ **Invalid JSON**: Request body không hợp lệ
- ✅ **Service Error**: Lỗi khi tạo event

#### 4. DELETE /events/:id
- ✅ **Success**: Xóa event thành công

## 🚀 Cách chạy Tests

### Chạy tests cụ thể
```bash
# Test cho Event Controller
go test ./tests/controllers -v

# Test với pattern matching
go test ./tests/controllers -run TestCreateEvent -v
```

### Chạy toàn bộ tests
```bash
# Tất cả tests trong project
go test ./...

# Với verbose output
go test ./... -v
```

### Test Coverage
```bash
# Coverage cơ bản
go test ./tests/controllers -cover

# Chi tiết coverage
go test ./tests/controllers -coverprofile=coverage.out
go tool cover -html=coverage.out

# Coverage cho toàn bộ project
go test ./... -coverprofile=coverage.out
```

### Test với flags hữu ích
```bash
# Chạy test nhiều lần
go test ./tests/controllers -count=5

# Timeout cho tests
go test ./tests/controllers -timeout=30s

# Chạy parallel tests
go test ./tests/controllers -parallel=4

# Race condition detection
go test ./tests/controllers -race
```

## 📝 Best Practices

### 1. Naming Convention
```go
func TestMethodName_Scenario(t *testing.T) {
    // TestGetEvents_Success
    // TestCreateEvent_InvalidJSON
    // TestDeleteEvent_NotFound
}
```

### 2. Test Structure (AAA Pattern)
```go
func TestExample(t *testing.T) {
    // Arrange - Setup test data and mocks
    mockService := new(MockService)
    controller := NewController(mockService)
    
    // Act - Execute the code under test
    result := controller.Method(input)
    
    // Assert - Verify the results
    assert.Equal(t, expected, result)
    mockService.AssertExpectations(t)
}
```

### 3. Mock Expectations
```go
// Return specific values
mockService.On("GetByID", uint(1)).Return(data, nil)

// Return different values for different calls
mockService.On("Create", mock.Anything).Return(nil, errors.New("error"))

// Match any argument of specific type
mockService.On("Update", mock.AnythingOfType("*models.User")).Return(nil)
```

### 4. HTTP Response Testing
```go
// Test status code
assert.Equal(t, http.StatusOK, w.Code)

// Test response body
var response map[string]interface{}
err := json.Unmarshal(w.Body.Bytes(), &response)
assert.NoError(t, err)
assert.Equal(t, "Success", response["message"])
```

## 🔧 Dependency Injection cho Testing

### Interface-based Design
```go
// internal/interfaces/event_service_interface.go
type EventServiceInterface interface {
    GetAllEvents() ([]models.Event, error)
    GetEventByID(id uint) (*models.Event, error)
    CreateEvent(req *models.CreateEventRequest) (*models.Event, error)
    DeleteEvent(id uint) error
}

// Controller sử dụng interface
type EventController struct {
    EventService interfaces.EventServiceInterface
}
```

## 📊 Test Output Example

```bash
=== RUN   TestGetEvents_Success
[GIN] 2025/08/22 - 09:51:06 | 200 |       202.8µs | GET "/events"
--- PASS: TestGetEvents_Success (0.00s)

=== RUN   TestCreateEvent_InvalidJSON  
[GIN] 2025/08/22 - 09:51:06 | 400 |            0s | POST "/events"
--- PASS: TestCreateEvent_InvalidJSON (0.00s)

PASS
ok  hello-gin/tests/controllers     0.119s
```

## 🎯 Mở rộng Tests

### Để thêm tests cho controllers khác:

1. **Tạo interface** cho service
2. **Refactor controller** để sử dụng interface  
3. **Tạo mock service** implementation
4. **Viết test cases** theo pattern đã có

### Template cho Controller mới:
```go
func TestNewController_Method_Scenario(t *testing.T) {
    // Setup
    mockService := new(MockNewService)
    controller := controllers.NewController(mockService)
    
    // Mock expectations
    mockService.On("Method", args).Return(data, nil)
    
    // Setup Gin
    r := tests.SetupTestGin()
    r.GET("/endpoint", controller.Method)
    
    // Create request
    req, _ := http.NewRequest("GET", "/endpoint", nil)
    w := httptest.NewRecorder()
    
    // Execute
    r.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    mockService.AssertExpectations(t)
}
```

## 🐛 Debugging Tests

### Test Failure Example

Khi test fail, Go sẽ hiển thị thông tin chi tiết về lỗi:

```bash
=== RUN   TestCreateEvent_InvalidJSON
[GIN] 2025/08/22 - 09:57:04 | 400 |       508.1µs | POST "/events"
    event_controller_test.go:239:
            Error Trace:    event_controller_test.go:239
            Error:          Not equal:
                            expected: 200
                            actual  : 400
            Test:           TestCreateEvent_InvalidJSON 

    event_controller_test.go:245:
            Error Trace:    event_controller_test.go:245
            Error:          Not equal:
                            expected: "Success message" 
                            actual  : "Invalid request body"

                            Diff:
                            --- Expected
                            +++ Actual
                            @@ -1 +1 @@
                            -Success message
                            +Invalid request body       
--- FAIL: TestCreateEvent_InvalidJSON (0.00s)
```

### Cách đọc Test Failure:

1. **Error Trace**: Chỉ ra dòng code bị lỗi
2. **Expected vs Actual**: So sánh giá trị mong đợi vs thực tế
3. **Diff**: Hiển thị sự khác biệt chi tiết

### Cách sửa Test Failure:

#### 1. **Status Code sai:**
```go
// ❌ Sai - expect sai status code
assert.Equal(t, http.StatusOK, w.Code) // expect 200

// ✅ Đúng - expect đúng status code  
assert.Equal(t, http.StatusBadRequest, w.Code) // expect 400
```

#### 2. **Response Message sai:**
```go
// ❌ Sai - expect sai message
assert.Equal(t, "Success message", response["error"])

// ✅ Đúng - expect đúng message
assert.Equal(t, "Invalid request body", response["error"])
```

#### 3. **Mock không match:**
```go
// ❌ Sai - mock setup không đúng arguments
mockService.On("GetEventByID", 1).Return(event, nil) // expect int

// ✅ Đúng - mock setup đúng type
mockService.On("GetEventByID", uint(1)).Return(event, nil) // expect uint
```

### Common Issues:

#### 1. Mock not set up correctly
```bash
panic: mock: Unexpected Method Call
```
**Solution**: Ensure all mocked methods are defined
```go
mockService.On("MethodName", args).Return(data, nil)
```

#### 2. Wrong assertion type
```bash
Error: expected string, got float64
```
**Solution**: Check JSON unmarshaling types
```go
// JSON numbers become float64 by default
assert.Equal(t, float64(1), response["id"])
// Or convert to specific type
assert.Equal(t, 1, int(response["id"].(float64)))
```

#### 3. Missing mock expectations
```bash
Error: mock: The method was expected to be called with arguments...
```
**Solution**: Verify all mock calls are set up
```go
mockService.AssertExpectations(t) // Fails if not all mocks called
```

### Verbose Test Output:
```bash
go test ./tests/controllers -v -run TestSpecificTest
```

### Test with Debug Info:
```go
t.Logf("Response: %+v", response)
t.Logf("Status Code: %d", w.Code)
t.Logf("Response Body: %s", w.Body.String())
```

### Test Specific Scenario:
```bash
# Chạy chỉ test cases có "Success" trong tên
go test ./tests/controllers -run Success -v

# Chạy test với timeout
go test ./tests/controllers -timeout=10s -v

# Chạy test với race detection  
go test ./tests/controllers -race -v
```

### Demo Test Failure và Cách Fix:

#### Ví dụ 1: Status Code sai
```go
// Test sẽ FAIL
func TestCreateEvent_Demo_StatusCodeFail(t *testing.T) {
    // ... setup code ...
    
    // ❌ WRONG: expect 200 but API returns 201
    assert.Equal(t, http.StatusOK, w.Code)
    
    // ✅ CORRECT: expect 201 for created resource
    assert.Equal(t, http.StatusCreated, w.Code)
}
```

#### Ví dụ 2: Mock Arguments không match
```go
// Test sẽ FAIL
func TestGetEvent_Demo_MockFail(t *testing.T) {
    mockService := new(MockEventService)
    
    // ❌ WRONG: setup mock với int nhưng service expect uint
    mockService.On("GetEventByID", 1).Return(event, nil)
    
    // ✅ CORRECT: sử dụng đúng type uint
    mockService.On("GetEventByID", uint(1)).Return(event, nil)
}
```

#### Output khi test FAIL:
```bash
--- FAIL: TestCreateEvent_InvalidJSON (0.00s)
    event_controller_test.go:239:
        Error Trace: event_controller_test.go:239
        Error:       Not equal:
                     expected: 200
                     actual  : 400
        Test:        TestCreateEvent_InvalidJSON
```

---

## 📚 Tài liệu tham khảo

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Gin Testing Guide](https://gin-gonic.com/docs/testing/)
- [Go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

---

**Kết quả hiện tại**: ✅ 9/9 test cases PASS cho Event API
