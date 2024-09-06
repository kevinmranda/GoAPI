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
	c.Get("user")
	Uploaded_bystr := c.Param("id")
	Uploaded_by64, _ := strconv.ParseUint(Uploaded_bystr, 10, 32)
	Uploaded_by := uint(Uploaded_by64)
	// struct for contents to be saved
	var body struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Filename    string  `json:"filename"` //(path to the high-quality image)
		Price       float64 `json:"price"`
	}

	// Get contents from the body of the request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create a photo instance and save to DB
	photo := models.Photo{
		Title:       body.Title,
		Description: body.Description,
		Filename:    body.Filename,
		Price:       body.Price,
		User_id:     Uploaded_by,
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

func Upload(c *gin.Context) {
	c.Get("user")
	// Parse the multipart form
	err := c.Request.ParseMultipartForm(20 << 20) // 20 MB limit
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
}

func GetPhoto(c *gin.Context) {

	// Get filename from request
	filename := c.Param("filename")

	// Define the directory where photos are stored
	filepath := "Photos/" + filename

	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// If file doesn't exist, return a 404
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "file not present",
		})
		return
	}

	// Serve the file
	c.File(filepath)
}

// get all photos
func GetPhotos(c *gin.Context) {
	c.Get("user")
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

func GetAllPhotos(c *gin.Context) {
	var photos []models.Photo

	initializers.DB.Find(&photos)

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
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Filename    string `json:"filename"` //(path to the high-quality image)
		Price       string `json:"price"`
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

	if contentForUpdate.Price != "" {
		iPrice, err := strconv.ParseFloat(contentForUpdate.Price, 64)

		photo.Price = iPrice

		fmt.Println("Error converting price:", err)

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
