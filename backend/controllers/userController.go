package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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

func ComparePassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, "Invalid password"
	}
	return true, "Password is valid"
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic("Error hashing password")
	}
	return string(bytes)
}
