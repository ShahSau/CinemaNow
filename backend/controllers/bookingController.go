package controllers

import "github.com/gin-gonic/gin"

// GetAllBookings is a function that returns all the bookings
func GetAllBookings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All bookings",
	})
}

// GetBooking is a function that returns a single booking
func GetBooking(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single booking",
	})
}

// CreateBooking is a function that creates a booking
func CreateBooking(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create booking",
	})
}

// UpdateBooking is a function that updates a booking
func UpdateBooking(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update booking",
	})
}

// DeleteBooking is a function that deletes a booking
func DeleteBooking(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete booking",
	})
}

// GetBookingByUser is a function that returns all the bookings by user
func GetBookingByUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bookings by user",
	})
}

// GetBookingByScreening is a function that returns all the bookings by screening
func GetBookingByScreening(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bookings by screening",
	})
}

// GetBookingByTransaction is a function that returns all the bookings by transaction
func GetBookingByTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bookings by transaction",
	})
}

// GetBookingByStatus is a function that returns all the bookings by status
func GetBookingByStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bookings by status",
	})
}
