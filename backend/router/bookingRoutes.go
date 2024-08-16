package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// BookingRoutes is a function that handles all the routes for the booking
func (r *routes) BookingsRoutes(rg *gin.RouterGroup) {
	booking := rg.Group("/booking")
	booking.Use(cors.Default())

	// get all bookings
	booking.GET("/all", controllers.GetAllBookings)

	// get a single booking
	booking.GET("/:id", controllers.GetBooking)

	// create a booking
	booking.POST("/create", controllers.CreateBooking)

	// update a booking
	booking.PUT("/:id", controllers.UpdateBooking)

	// delete a booking
	booking.DELETE("/:id", controllers.DeleteBooking)

	// get all bookings by user
	booking.GET("/user/:id", controllers.GetBookingByUser)

	// get all bookings by screening
	booking.GET("/screening/:id", controllers.GetBookingByScreening)

	// get all bookings by transaction
	booking.GET("/transaction/:id", controllers.GetBookingByTransaction)

	// get all bookings by status
	booking.GET("/status/:status", controllers.GetBookingByStatus)
}
