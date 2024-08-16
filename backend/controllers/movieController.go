package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllMovies is a function that returns all the movies
func GetAllMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All movies",
	})
}

// GetMovie is a function that returns a single movie
func GetMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single movie",
	})
}

// CreateMovie is a function that creates a movie
func CreateMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create movie",
	})
}

// UpdateMovie is a function that updates a movie
func UpdateMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update movie",
	})
}

// DeleteMovie is a function that deletes a movie
func DeleteMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete movie",
	})
}

// SearchMovie is a function that searches for a movie
func SearchMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search movie",
	})
}

// GetMovieByGenre is a function that returns all the movies by genre
func GetMovieByGenre(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Movies by genre",
	})
}

// GetMovieByRating is a function that returns all the movies by rating
func GetMovieByRating(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Movies by rating",
	})
}

// GetMovieByYear is a function that returns all the movies by year
func GetMovieByYear(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Movies by year",
	})
}

// get popular movies
func GetPopularMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Popular movies",
	})
}

// get upcoming movies
func GetUpcomingMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Upcoming movies",
	})
}

// get now playing movies
func GetNowPlayingMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Now playing movies",
	})
}

// get top rated movies
func GetTopRatedMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Top rated movies",
	})
}
