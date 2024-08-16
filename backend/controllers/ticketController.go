package controllers

import "github.com/gin-gonic/gin"

// GetAllTickets is a function that returns all the tickets
func GetAllTickets(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All tickets",
	})
}

// GetTicket is a function that returns a single ticket
func GetTicket(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single ticket",
	})
}

// CreateTicket is a function that creates a ticket
func CreateTicket(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create ticket",
	})
}

// UpdateTicket is a function that updates a ticket
func UpdateTicket(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update ticket",
	})
}

// DeleteTicket is a function that deletes a ticket
func DeleteTicket(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Delete ticket",
	})
}

// GetTicketByBooking is a function that returns all the tickets by booking
func GetTicketByBooking(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get ticket by booking",
	})
}

// GetTicketByStatus is a function that returns all the tickets by status
func GetTicketByStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get ticket by status",
	})
}
