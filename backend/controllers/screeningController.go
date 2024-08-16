package controllers

import "github.com/gin-gonic/gin"

// GetAllScreenings is a function that returns all the screenings
func GetAllScreenings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All screenings",
	})
}

// GetScreening is a function that returns a single screening
func GetScreening(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single screening",
	})
}

// CreateScreening is a function that creates a screening
func CreateScreening(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create screening",
	})
}

// UpdateScreening is a function that updates a screening
func UpdateScreening(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update screening",
	})
}

// DeleteScreening is a function that deletes a screening
func DeleteScreening(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete screening",
	})
}

// GetScreeningByMovie is a function that returns all the screenings by movie
func GetScreeningByMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by movie",
	})
}

// GetScreeningByAuditorium is a function that returns all the screenings by auditorium
func GetScreeningByAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by auditorium",
	})
}

// GetScreeningByDate is a function that returns all the screenings by date
func GetScreeningByDate(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Get screening by date",
	})
}

// GetScreeningByTime is a function that returns all the screenings by time
func GetScreeningByTime(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by time",
	})
}

// GetScreeningByStatus is a function that returns all the screenings by status
func GetScreeningByStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by status",
	})
}

// GetScreeningByMovieAndDate is a function that returns all the screenings by movie and date
func GetScreeningByMovieAndDate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by movie and date",
	})
}

// GetScreeningByAuditoriumAndDate is a function that returns all the screenings by auditorium and date
func GetScreeningByAuditoriumAndDate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by auditorium and date",
	})
}

// GetScreeningByAuditoriumAndTime is a function that returns all the screenings by auditorium and time
func GetScreeningByAuditoriumAndTime(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by auditorium and time",
	})
}

// GetScreeningByMovieAndAuditorium is a function that returns all the screenings by movie and auditorium
func GetScreeningByMovieAndAuditorium(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get screening by movie and auditorium",
	})
}
