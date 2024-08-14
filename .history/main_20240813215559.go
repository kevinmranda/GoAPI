package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	
	r.Run()
}
