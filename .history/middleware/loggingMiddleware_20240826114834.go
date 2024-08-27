package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LogRequestResponseMiddleware logs the details of HTTP requests and responses
func LogRequestResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request details
		startTime := time.Now()
		requestBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody)) // Restore body for further usage

		// Log request details
		log.Printf("Request: Method=%s, URL=%s, Body=%s", c.Request.Method, c.Request.RequestURI, string(requestBody))

		// Capture response details
		responseWriter := &responseCapture{ResponseWriter: c.Writer, body: new(bytes.Buffer)}
		c.Writer = responseWriter

		// Process request
		c.Next()

		// Log response details
		duration := time.Since(startTime)
		log.Printf("Response: Status=%d, Duration=%v, Body=%s", responseWriter.statusCode, duration, responseWriter.body.String())
	}
}

// responseCapture is a custom ResponseWriter that captures the response body
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
