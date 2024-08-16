package controllers

import "github.com/gin-gonic/gin"

// GetAllAuditoriums is a function that returns all the auditoriums
func GetAllAuditoriums(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All movies",
	})
}

// GetAuditorium is a function that returns a single auditorium
func GetAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single movie",
	})
}

// CreateAuditorium is a function that creates an auditorium
func CreateAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create movie",
	})
}

// UpdateAuditorium is a function that updates an auditorium
func UpdateAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update movie",
	})
}

// DeleteAuditorium is a function that deletes an auditorium
func DeleteAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete movie",
	})
}

// GetAuditoriumByName is a function that returns all the auditoriums by name
func GetAuditoriumByName(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get movie by name",
	})
}

// GetAuditoriumBySeatingCapacity is a function that returns all the auditoriums by seating capacity
func GetAuditoriumBySeatingCapacity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get movie by seating capacity",
	})
}

// GetSeatArrangement is a function that returns the seat arrangement of an auditorium
func GetSeatArrangement(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get seat arrangement",
	})
}
