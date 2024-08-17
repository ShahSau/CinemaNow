package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ScreeningRoutes is a function that handles all the routes for the screening
func (r *routes) ScreeningRoutes(rg *gin.RouterGroup) {
	screening := rg.Group("/screening")
	screening.Use(cors.Default())

	// get all screenings
	screening.GET("/all", controllers.GetAllScreenings)

	// get a single screening
	screening.GET("/:id", controllers.GetScreening)

	// create a screening
	screening.POST("/create", controllers.CreateScreening)

	// update a screening
	screening.PUT("/:id", controllers.UpdateScreening)

	// delete a screening
	screening.DELETE("/:id", controllers.DeleteScreening)

	// get all screenings by movie
	screening.GET("/movie/:id", controllers.GetScreeningByMovie)

	// get all screenings by auditorium
	screening.GET("/auditorium/:id", controllers.GetScreeningByAuditorium)

	// get all screenings by date
	screening.GET("/date/:date", controllers.GetScreeningByDate)

	// get all screenings by time
	screening.GET("/time/:time", controllers.GetScreeningByTime)

	// get all screenings by status
	screening.GET("/status/:status", controllers.GetScreeningByStatus)

	// get all screenings by movie and date
	// screening.GET("/movie/:movieid/date/:date", controllers.GetScreeningByMovieAndDate)

	// get all screenings by auditorium and date
	// screening.GET("/auditorium/:auditoriumid/date/:date", controllers.GetScreeningByAuditoriumAndDate)

	// get all screenings by auditorium and time
	// screening.GET("/auditorium/:auditoriumid/time/:time", controllers.GetScreeningByAuditoriumAndTime)

	// get all screenings by movie and auditorium
	// screening.GET("/movie/:movieid/auditorium/:auditoriumid", controllers.GetScreeningByMovieAndAuditorium)
}
