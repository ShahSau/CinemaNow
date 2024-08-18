package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TicketsRoutes is a function that handles all the routes for the tickets
func (r *routes) TicketsRoutes(rg *gin.RouterGroup) {
	tickets := rg.Group("/tickets")
	tickets.Use(cors.Default())

	// get all tickets
	tickets.GET("/all", controllers.GetAllTickets)

	// get a single ticket
	tickets.GET("/:id", controllers.GetTicket)

	// create a ticket
	tickets.POST("/create", controllers.CreateTicket)

	// update a ticket
	tickets.PUT("/:id", controllers.UpdateTicket)

	// delete a ticket
	tickets.DELETE("/:id", controllers.DeleteTicket)

	// get all tickets by booking
	tickets.GET("/booking/:id", controllers.GetTicketByBooking)

}
