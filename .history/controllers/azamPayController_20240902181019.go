package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthResponse struct {
	Data struct {
		AccessToken json.RawMessage `json:"accessToken"`
		Expire      json.RawMessage `json:"expire"`
	} `json:"data"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
}

// checkoutUrl := os.Getenv("AZAMPAY_CHECKOUT")

func AuthToken(c *gin.Context) string {
	authUrl := os.Getenv("AZAMPAY_AUTH")
	appName := os.Getenv("AZAMPAY_APP_NAME")
	clientID := os.Getenv("AZAMPAY_CLIENT_ID")
	clientSecret := os.Getenv("AZAMPAY_CLIENT_SECRET")

	// Create JSON body
	jsonBody := fmt.Sprintf(`{
		"appName": "%s",
		"clientId": "%s",
		"clientSecret": "%s"
	}`, appName, clientID, clientSecret)

	bodyReader := bytes.NewReader([]byte(jsonBody))

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, authUrl, bodyReader)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
	}

	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Authentication failed with status code: %d", resp.StatusCode)
		c.JSON(resp.StatusCode, gin.H{"error": "Authentication failed"})
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
	}

	// Unmarshal the response into the struct
	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
	}

	// Return the parsed data as a JSON response
	c.JSON(http.StatusOK, gin.H{
		"message":     authResp.Message,
		"accessToken": authResp.Data.AccessToken,
		"expires":     authResp.Data.Expire,
	})
	return string(authResp.Data.AccessToken)
}

func MNOCheckout() {

}
