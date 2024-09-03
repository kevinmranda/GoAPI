package controllers

import (
	"bytes"
	"net/http"
	"os"
)

func AuthToken() {
	authUrl := os.Getenv("AZAMPAY_AUTH")
	appName := os.Getenv("AZAMPAY_APP_NAME")
	clientID := os.Getenv("AZAMPAY_CLIENT_ID")
	clientSecret := os.Getenv("AZAMPAY_CLIENT_SECRET")
	// checkoutUrl := os.Getenv("AZAMPAY_CHECKOUT")

	jsonBody := []byte(
		`{
	"appName": %s,
	"clientId": %s,
	"clientSecret":%s
	}`, appName, clientID, clientSecret,
	)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, authUrl, bodyReader)

}

func MNOCheckout() {

}
