package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

type MNOCheckoutResponse struct {
	TransactionId string `json:"transactionId"`
	Message       string `json:"message"`
	Success       bool   `json:"success"`
}

func AuthToken() string {
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

	}

	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Authentication failed with status code: %d", resp.StatusCode)
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into the struct
	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Return the parsed auth token
	return string(authResp.Data.AccessToken)
}

func MNOCheckout(accountNumber, amount, currency, externalId, provider string) bool {
	// Generate the token
	token := AuthToken()

	// Create JSON body
	jsonBody := fmt.Sprintf(`{
		"accountNumber": "%s",
		"additionalProperties":{
			"property1": null,
			"property2": null
		},
		"amount": "%s",
		"currency": "%s",
		"externaId": "%s",
		"provider": "%s"
	}`, accountNumber, amount, currency, externalId, provider)

	checkoutUrl := os.Getenv("AZAMPAY_CHECKOUT")
	bodyReader := bytes.NewReader([]byte(jsonBody))

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, checkoutUrl, bodyReader)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set headers separately
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Handle response (optional but recommended)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Printf("Response: %s\n", body)

	// Unmarshal the response into the struct
	var checkoutRes MNOCheckoutResponse
	if err := json.Unmarshal(body, &checkoutRes); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Return the parsed auth token
	return checkoutRes.Success
}
