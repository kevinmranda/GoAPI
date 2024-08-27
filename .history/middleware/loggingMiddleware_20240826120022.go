package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

// Database connection
// var db *gorm.DB // Initialize your database connection

// LogRequestResponseMiddleware logs the details of HTTP requests and responses
func LogRequestResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request details
		startTime := time.Now()
		requestBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // Restore body for further usage

		// Log request details
		logEntry := fmt.Sprintf("Request: Method=%s, URL=%s, Body=%s", c.Request.Method, c.Request.RequestURI, string(requestBody))
		insertLog("INFO", logEntry, nil)

		// Capture response details
		responseWriter := &responseCapture{ResponseWriter: c.Writer, body: new(bytes.Buffer)}
		c.Writer = responseWriter

		// Process request
		c.Next()

		// Log response details
		duration := time.Since(startTime)
		logEntry = fmt.Sprintf("Response: Status=%d, Duration=%v, Body=%s", responseWriter.statusCode, duration, responseWriter.body.String())
		insertLog("INFO", logEntry, nil)
	}
}

// insertLog inserts a log entry into the database
func insertLog(level, message string, details map[string]interface{}) {
	initializers.DB.Create(&models.ActivityLog{
		Level:   level,
		Message: message,
		Details: details,
	})
}

type responseCapture struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rc *responseCapture) Write(b []byte) (int, error) {
	rc.body.Write(b)
	return rc.ResponseWriter.Write(b)
}

func (rc *responseCapture) WriteHeader(code int) {
	rc.statusCode = code
	rc.ResponseWriter.WriteHeader(code)
}
