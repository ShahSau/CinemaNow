package controllers

import "github.com/gin-gonic/gin"

// GetAllTheaters is a function that returns all the theaters
func GetAllTheaters(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All theaters",
	})
}

// GetTheater is a function that returns a single theater
func GetTheater(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single theater",
	})
}

// CreateTheater is a function that creates a theater
func CreateTheater(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create theater",
	})
}

// UpdateTheater is a function that updates a theater
func UpdateTheater(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update theater",
	})
}

// DeleteTheater is a function that deletes a theater
func DeleteTheater(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete theater",
	})
}

// GetTheaterByName is a function that returns all the theaters by name
func GetTheaterByName(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get theater by name",
	})
}

// GetTheaterByLocation is a function that returns all the theaters by location
func GetTheaterByLocation(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get theater by location",
	})
}

// GetTheaterBySeatingCapacity is a function that returns all the theaters by seating capacity
func GetTheaterBySeatingCapacity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get theater by seating capacity",
	})
}
