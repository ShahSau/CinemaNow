package controllers

import "github.com/gin-gonic/gin"

// GetAllUsers is a function that returns all the users
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All users",
	})
}

// GetUser is a function that returns a single user
func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single user",
	})
}

// CreateUser is a function that creates a user
func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create user",
	})
}

// UpdateUser is a function that updates a user
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update user",
	})
}

// DeleteUser is a function that deletes a user
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete user",
	})
}
