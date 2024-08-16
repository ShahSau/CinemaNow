package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// AuditoriumRoutes is a function that handles all the routes for the auditorium
func (r *routes) AuditoriumRoutes(rg *gin.RouterGroup) {
	auditorium := rg.Group("/auditorium")
	auditorium.Use(cors.Default())

	// get all auditoriums
	auditorium.GET("/all", controllers.GetAllAuditoriums)

	// get a single auditorium
	auditorium.GET("/:id", controllers.GetAuditorium)

	// create an auditorium
	auditorium.POST("/create", controllers.CreateAuditorium)

	// update an auditorium
	auditorium.PUT("/:id", controllers.UpdateAuditorium)

	// delete an auditorium
	auditorium.DELETE("/:id", controllers.DeleteAuditorium)

	// get all auditoriums by name
	auditorium.GET("/name/:name", controllers.GetAuditoriumByName)

	// get all auditoriums by seating capacity
	auditorium.GET("/seatingcapacity/:seatingcapacity", controllers.GetAuditoriumBySeatingCapacity)

	// customized seat arrangement
	auditorium.GET("/seatarrangement/:id", controllers.GetSeatArrangement)
}
