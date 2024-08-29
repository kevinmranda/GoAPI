package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddPhoto(c *gin.Context) {
	// Retrieve the user (assuming you have some middleware that sets the user)
	user, _ := c.Get("user")

	// Parse the multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse form data",
		})
		return
	}

	// Get the file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve the file",
		})
		return
	}
	defer file.Close()

	// Define the file path
	filename := header.Filename
	filePath := fmt.Sprintf("Photos/%s", filename)

	// Create the Photos directory if it doesn't exist
	if err := os.MkdirAll("Photos", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create directory",
		})
		return
	}

	// Save the file to the Photos folder
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the file",
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the file",
		})
		return
	}

	// Retrieve other details from the request body
	title := c.PostForm("title")
	description := c.PostForm("description")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price format",
		})
		return
	}

	uploadedBy := user.(uint)

	// Create a photo instance and save to DB
	photo := models.Photo{
		Title:       title,
		Description: description,
		Filename:    filePath, // Path to the saved file
		Price:       price,
		User_id:     uploadedBy,
	}

	result := initializers.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "Record inserted successfully",
		"data":    photo,
	})
}

// get a specific photo with id
func GetPhoto(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	var photo models.Photo

	// Check if the photo exists
	initializers.DB.First(&photo, id)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    photo,
	})
}

// get all photos
func GetPhotos(c *gin.Context) {
	id := c.Param("id")
	var photos []models.Photo

	// Preload the Orders relationship to include associated orders in the results
	result := initializers.DB.Preload("Orders").Where("user_id = ?", id).Find(&photos)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "records not present",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    photos,
	})
}

// update a photo with id
func UpdatePhoto(c *gin.Context) {
	c.Get("user")
	// Get id from request
	id := c.Param("id")

	//struct for contents to be updated
	var contentForUpdate struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Filename    string  `json:"filename"` //(path to the high-quality image)
		Price       float64 `json:"price"`
		Uploaded_by uint    `json:"uploaded_by"` //Uploaded by
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&contentForUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check if the photo to be updated exists
	var photo models.Photo
	result := initializers.DB.Preload("Orders").First(&photo, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	//update photo with struct provided on request body
	if contentForUpdate.Title != "" {
		photo.Title = contentForUpdate.Title
	}

	if contentForUpdate.Description != "" {
		photo.Description = contentForUpdate.Description
	}

	if contentForUpdate.Filename != "" {
		photo.Filename = contentForUpdate.Filename
	}

	if contentForUpdate.Price != 0 {
		photo.Price = contentForUpdate.Price
	}

	if contentForUpdate.Uploaded_by != 0 {
		photo.User_id = contentForUpdate.Uploaded_by
	}

	result = initializers.DB.Save(&photo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2014,
			"error": "Failed to update the record",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    photo,
	})
}

// delete a photo with specified id
func DeletePhoto(c *gin.Context) {
	c.Get("user")
	// Get id from request
	id := c.Param("id")

	var photo models.Photo
	// Check if the photo exists
	initializers.DB.First(&photo, id)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	// Delete the photo
	result := initializers.DB.Delete(&photo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    2013,
			"error": "failed to delete record",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2012,
		"message": "record deleted successfully",
	})
}
