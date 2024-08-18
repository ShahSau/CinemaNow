package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// MovieDetailsRoutes is a function that handles all the routes for the movie details
func (r *routes) MovieDetailsRoutes(rg *gin.RouterGroup) {
	movieDetailsRouteGrouping := rg.Group("/movieDetails")
	movieDetailsRouteGrouping.Use(cors.Default())

	// get all movie details
	movieDetailsRouteGrouping.GET("/all", controllers.GetAllMovieDetails)
	// get a single movie detail
	movieDetailsRouteGrouping.GET("/:id", controllers.GetMovieDetails)

	// create a movie detail
	movieDetailsRouteGrouping.POST("/create", controllers.CreateMovieDetail)

	// update a movie detail
	movieDetailsRouteGrouping.PUT("/:id", controllers.UpdateMovieDetail)

	// delete a movie detail
	movieDetailsRouteGrouping.DELETE("/:id", controllers.DeleteMovieDetail)

	// search for a movie detail by imdb_id
	movieDetailsRouteGrouping.GET("/search/:imdbId", controllers.SearchMovieDetailByIMDBID)

	// search for a movie detail by genre
	movieDetailsRouteGrouping.GET("/genre/:genre", controllers.SearchMovieDetailByGenre)

	// search for a movie detail by production company
	movieDetailsRouteGrouping.GET("/productionCompany/:productionCompany", controllers.SearchMovieDetailByProductionCompany)
}
