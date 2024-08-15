package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddPhoto(c *gin.Context) {
	//struct of the request body
	var body struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Filename    string  `json:"filename" binding:"required"` //(path to the high-quality image)
		Price       float64 `json:"price" binding:"required"`
		User_id     uint    `json:"user_id" binding:"required"` //Uploaded by
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//add photo to db
	photo := models.Photo{
		Title:       body.Title,
		Description: body.Description,
		Filename:    body.Filename,
		Price:       body.Price,
		User_id:     body.User_id,
	}
	result := initializers.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "record inserted successfully",
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

// update a photo with id
func UpdatePhoto(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	//struct for contents to be updated
	var contentForUpdate struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Filename    string  `json:"filename"` //(path to the high-quality image)
		Price       float64 `json:"price"`
		User_id     uint    `json:"user_id"` //Uploaded by
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&contentForUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var photo models.Photo

	// Check if the photo to be updated exists
	initializers.DB.First(&photo, id)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	//update photo with struct provided on request body
	updatedPhoto := models.Photo{
		Title:       contentForUpdate.Title,
		Description: contentForUpdate.Description,
		Filename:    contentForUpdate.Filename,
		Price:       contentForUpdate.Price,
		User_id:     contentForUpdate.User_id,
	}
	
	resultinitializers.DB.Updates(&updatedPhoto)
	if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"id":    2000,
				"error": "Failed to insert the record",
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
