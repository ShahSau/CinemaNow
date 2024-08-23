package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ShahSau/CinemaNow/backend/database"
	"github.com/ShahSau/CinemaNow/backend/models"
	"github.com/ShahSau/CinemaNow/backend/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var screeningsCollection *mongo.Collection = database.GetCollection(database.DB, "screenings")

// GetAllScreenings is a function that returns all the screenings
func GetAllScreenings(c *gin.Context) {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errp := strconv.Atoi(c.Query("page"))
	if errp != nil || page < 1 {
		page = 1
	}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "auditorium_id", Value: 1},
			{Key: "cinema_id", Value: 1},
			{Key: "movie_id", Value: 1},
			{Key: "start_time", Value: 1},
			{Key: "auditorium", Value: 1},
			{Key: "theater", Value: 1},
			{Key: "movie", Value: 1},
			{Key: "bookable", Value: 1},
		}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := screeningsCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var screenings []bson.M

	if err = result.All(c.Request.Context(), &screenings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":          page,
		"recordPerPage": recordPerPage,
		"screenings":    screenings,
		"message":       "All screenings",
		"error":         false,
		"total":         len(screenings),
	})

}

// GetScreening is a function that returns a single screening
func GetScreening(c *gin.Context) {
	var screening models.Screening

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}

	err := screeningsCollection.FindOne(c.Request.Context(), filter).Decode(&screening)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screening": screening,
		"message":   "Screening found",
		"success":   true,
		"error":     false,
	})

}

// CreateScreening is a function that creates a screening
func CreateScreening(c *gin.Context) {
	var reqscreening types.Screening

	if err := c.ShouldBindJSON(&reqscreening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	var screening models.Screening

	screening.ID = primitive.NewObjectID()
	screening.AuditoriumId = reqscreening.AuditoriumId
	screening.CinemaId = reqscreening.CinemaId
	screening.MovieId = reqscreening.MovieId
	screening.StartTime = reqscreening.StartTime
	screening.Bookable = reqscreening.Bookable
	screening.Auditorium = make([]models.Auditorium, len(reqscreening.Auditorium))
	for i, auditorium := range reqscreening.Auditorium {
		screening.Auditorium[i] = models.Auditorium{
			Name:          auditorium.Name,
			Rows:          auditorium.Rows,
			Columns:       auditorium.Columns,
			Seats:         make([]models.Seat, len(auditorium.Seats)),
			SelectedSeats: make([]models.Seat, len(auditorium.SelectedSeats)),
		}
	}
	screening.Theater = make([]models.Theater, len(reqscreening.Theater))
	for i, theater := range reqscreening.Theater {
		screening.Theater[i] = models.Theater{
			Name:    theater.Name,
			Address: theater.Address,
		}
	}
	screening.Movie = make([]models.Movie, len(reqscreening.Movie))
	for i, movie := range reqscreening.Movie {
		screening.Movie[i] = models.Movie{
			MovieID:          movie.MovieID,
			Type:             movie.Type,
			Adult:            movie.Adult,
			BackdropPath:     movie.BackdropPath,
			GenreIds:         movie.GenreIds,
			OriginalLanguage: movie.OriginalLanguage,
			OriginalTitle:    movie.OriginalTitle,
			Overview:         movie.Overview,
			Popularity:       movie.Popularity,
			PosterPath:       movie.PosterPath,
			ReleaseDate:      movie.ReleaseDate,
			Title:            movie.Title,
			Video:            movie.Video,
			VoteAverage:      movie.VoteAverage,
			VoteCount:        movie.VoteCount,
		}
	}
	screening.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	screening.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := screeningsCollection.InsertOne(c.Request.Context(), screening)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not created",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screening": screening,
		"message":   "Screening created",
		"success":   true,
		"error":     false,
	})

}

// UpdateScreening is a function that updates a screening
func UpdateScreening(c *gin.Context) {
	var reqscreening types.Screening

	if err := c.ShouldBindJSON(&reqscreening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	screening, _ := primitive.ObjectIDFromHex(c.Param("id"))
	searchedScreening, err := screeningsCollection.Find(c.Request.Context(), bson.M{"_id": screening})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
		})
		return
	}

	if searchedScreening == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Screening not found",
			"success": false,
			"error":   true,
		})
		return
	}

	var screeningUpdated models.Screening
	screeningUpdated.ID = screening
	screeningUpdated.AuditoriumId = reqscreening.AuditoriumId
	screeningUpdated.CinemaId = reqscreening.CinemaId
	screeningUpdated.MovieId = reqscreening.MovieId
	screeningUpdated.StartTime = reqscreening.StartTime
	screeningUpdated.Bookable = reqscreening.Bookable
	screeningUpdated.Auditorium = make([]models.Auditorium, len(reqscreening.Auditorium))
	for i, auditorium := range reqscreening.Auditorium {
		screeningUpdated.Auditorium[i] = models.Auditorium{
			Name:          auditorium.Name,
			Rows:          auditorium.Rows,
			Columns:       auditorium.Columns,
			Seats:         make([]models.Seat, len(auditorium.Seats)),
			SelectedSeats: make([]models.Seat, len(auditorium.SelectedSeats)),
		}
	}
	screeningUpdated.Theater = make([]models.Theater, len(reqscreening.Theater))
	for i, theater := range reqscreening.Theater {
		screeningUpdated.Theater[i] = models.Theater{
			Name:    theater.Name,
			Address: theater.Address,
		}
	}
	screeningUpdated.Movie = make([]models.Movie, len(reqscreening.Movie))
	for i, movie := range reqscreening.Movie {
		screeningUpdated.Movie[i] = models.Movie{
			MovieID:          movie.MovieID,
			Type:             movie.Type,
			Adult:            movie.Adult,
			BackdropPath:     movie.BackdropPath,
			GenreIds:         movie.GenreIds,
			OriginalLanguage: movie.OriginalLanguage,
			OriginalTitle:    movie.OriginalTitle,
			Overview:         movie.Overview,
			Popularity:       movie.Popularity,
			PosterPath:       movie.PosterPath,
			ReleaseDate:      movie.ReleaseDate,
			Title:            movie.Title,
			Video:            movie.Video,
			VoteAverage:      movie.VoteAverage,
			VoteCount:        movie.VoteCount,
		}
	}

	_, err = screeningsCollection.UpdateOne(c.Request.Context(), bson.M{"_id": screening}, screeningUpdated)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not updated",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screening": screeningUpdated,
		"message":   "Screening updated",
		"success":   true,
		"error":     false,
	})

}

// DeleteScreening is a function that deletes a screening
func DeleteScreening(c *gin.Context) {
	screening := c.Param("id")

	_, err := screeningsCollection.DeleteOne(c.Request.Context(), bson.M{"_id": screening})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not deleted",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Screening deleted",
		"success": true,
		"error":   false,
	})
}

// GetScreeningByMovie is a function that returns all the screenings by movie
func GetScreeningByMovie(c *gin.Context) {
	var screenings []models.Screening
	movie := c.Param("movie")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"movie": movie})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByAuditorium is a function that returns all the screenings by auditorium
func GetScreeningByAuditorium(c *gin.Context) {
	var screenings []models.Screening
	auditorium := c.Param("auditorium")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"auditorium": auditorium})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByDate is a function that returns all the screenings by date
func GetScreeningByDate(c *gin.Context) {
	var screenings []models.Screening
	date := c.Param("date")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"start_time": date})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByTime is a function that returns all the screenings by time
func GetScreeningByTime(c *gin.Context) {
	var screenings []models.Screening
	time := c.Param("time")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"start_time": time})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByStatus is a function that returns all the screenings by status
func GetScreeningByStatus(c *gin.Context) {
	var screenings []models.Screening
	status := c.Param("status")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"bookable": status})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByMovieAndDate is a function that returns all the screenings by movie and date
func GetScreeningByMovieAndDate(c *gin.Context) {
	var screenings []models.Screening
	movie := c.Param("movie")
	date := c.Param("date")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"movie": movie, "start_time": date})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByAuditoriumAndDate is a function that returns all the screenings by auditorium and date
func GetScreeningByAuditoriumAndDate(c *gin.Context) {
	var screenings []models.Screening
	auditorium := c.Param("auditorium")
	date := c.Param("date")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"auditorium": auditorium, "start_time": date})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByAuditoriumAndTime is a function that returns all the screenings by auditorium and time
func GetScreeningByAuditoriumAndTime(c *gin.Context) {
	var screenings []models.Screening
	auditorium := c.Param("auditorium")
	time := c.Param("time")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"auditorium": auditorium, "start_time": time})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}

// GetScreeningByMovieAndAuditorium is a function that returns all the screenings by movie and auditorium
func GetScreeningByMovieAndAuditorium(c *gin.Context) {
	var screenings []models.Screening
	movie := c.Param("movie")
	auditorium := c.Param("auditorium")

	cursor, err := screeningsCollection.Find(c.Request.Context(), bson.M{"movie": movie, "auditorium": auditorium})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var screening models.Screening
		cursor.Decode(&screening)
		screenings = append(screenings, screening)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Screening not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
		"message":    "Screening found",
		"success":    true,
		"error":      false,
	})
}
