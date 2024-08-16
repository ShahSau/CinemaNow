package controllers

import (
	"github.com/gin-gonic/gin"
)

// Register is a function that handles the registration of a user
func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register",
	})
}

// Login is a function that handles the login of a user
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}

// Logout is a function that handles the logout of a user
func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout",
	})
}

// ForgotPassword is a function that handles the forgot password of a user
func ForgotPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forgot password",
	})
}

// ResetPassword is a function that handles the reset password of a user
func ResetPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Reset password",
	})
}
