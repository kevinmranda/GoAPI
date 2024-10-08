package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Database connection
// var db *gorm.DB // Initialize your database connection

// LogRequestResponseMiddleware logs the details of HTTP requests and responses
func LogRequestResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request details
		startTime := time.Now()
		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody)) // Restore body for further usage

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
	db.Create(&ActivityLog{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Details:   details,
	})
}

// ActivityLog represents a log entry in the database
type ActivityLog struct {
	ID        uint `gorm:"primaryKey"`
	Timestamp time.Time
	Level     string
	Message   string
	Details   map[string]interface{} `gorm:"type:jsonb"`
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
