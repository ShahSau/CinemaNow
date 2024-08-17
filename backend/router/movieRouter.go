package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// MovieRoutes is a function that handles all the routes for the movie
func (r *routes) MovieRoutes(rg *gin.RouterGroup) {
	movieRouteGrouping := rg.Group("/movies")
	movieRouteGrouping.Use(cors.Default())

	// get all movies
	movieRouteGrouping.GET("/all", controllers.GetAllMovies)

	// get a single movie
	movieRouteGrouping.GET("/:id", controllers.GetMovie)

	// create a movie
	movieRouteGrouping.POST("/create", controllers.CreateMovie)

	// update a movie
	movieRouteGrouping.PUT("/:id", controllers.UpdateMovie)

	// delete a movie
	movieRouteGrouping.DELETE("/:id", controllers.DeleteMovie)

	// search for a movie
	movieRouteGrouping.GET("/search", controllers.SearchMovie)

	// get all movies by rating
	movieRouteGrouping.GET("/rating/:rating", controllers.GetMovieByRating)

	// get all movies by year
	movieRouteGrouping.GET("/year/:year", controllers.GetMovieByYear)

	// get popular movies
	movieRouteGrouping.GET("/popular", controllers.GetPopularMovies)

	// get upcoming movies
	movieRouteGrouping.GET("/upcoming", controllers.GetUpcomingMovies)

	// get now playing movies
	movieRouteGrouping.GET("/nowplaying", controllers.GetNowPlayingMovies)

	// get top rated movies
	movieRouteGrouping.GET("/toprated", controllers.GetTopRatedMovies)
}
