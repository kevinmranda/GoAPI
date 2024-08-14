package main

import "github.com/kevinmranda/GoAPI/initializers"

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.Run()
}