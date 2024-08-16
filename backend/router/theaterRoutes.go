package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TheaterRoutes is a function that handles all the routes for the theater
func (r *routes) TheaterRoutes(rg *gin.RouterGroup) {
	theater := rg.Group("/theater")
	theater.Use(cors.Default())

	// get all theaters
	theater.GET("/all", controllers.GetAllTheaters)

	// get a single theater
	theater.GET("/:id", controllers.GetTheater)

	// create a theater
	theater.POST("/create", controllers.CreateTheater)

	// update a theater
	theater.PUT("/:id", controllers.UpdateTheater)

	// delete a theater
	theater.DELETE("/:id", controllers.DeleteTheater)

	// get all theaters by name
	theater.GET("/name/:name", controllers.GetTheaterByName)

	// get all theaters by location
	theater.GET("/location/:location", controllers.GetTheaterByLocation)

	// get all theaters by seating capacity
	theater.GET("/seatingcapacity/:seatingcapacity", controllers.GetTheaterBySeatingCapacity)
}
