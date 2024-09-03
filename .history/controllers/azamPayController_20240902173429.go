package controllers

import (
	"os"
)

func AuthToken() {
	authUrl := os.Getenv("AZAMPAY_AUTH")
	// checkoutUrl := os.Getenv("AZAMPAY_CHECKOUT")

	jsonBody :={
"appName": "string",
"clientId": "string",
"clientSecret": "string"
}
	

}

func MNOCheckout() {

}
