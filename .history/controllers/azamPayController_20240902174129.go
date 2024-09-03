package controllers

import (
	"bytes"
	"net/http"
	"os"
)

func AuthToken() {
	authUrl := os.Getenv("AZAMPAY_AUTH")
	// checkoutUrl := os.Getenv("AZAMPAY_CHECKOUT")

	jsonBody := []byte(
		`{
	"appName": "string",
	"clientId": "string",
	"clientSecret": "string"
	}`,
	)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, authUrl, bodyReader)

}

func MNOCheckout() {

}
