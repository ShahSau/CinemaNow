package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ShahSau/CinemaNow/backend/database"
	"github.com/ShahSau/CinemaNow/backend/models"
	"github.com/ShahSau/CinemaNow/backend/types"
)

var movieCollection *mongo.Collection = database.GetCollection(database.DB, "movies")

// GetAllMovies is a function that returns all the movies
func GetAllMovies(c *gin.Context) {
	var movies []models.Movie

	cursor, err := movieCollection.Find(c.Request.Context(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All movies",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// GetMovie is a function that returns a single movie
func GetMovie(c *gin.Context) {
	var movie models.Movie

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": id}

	err := movieCollection.FindOne(c.Request.Context(), filter).Decode(&movie)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie",
		"movie":   movie,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// CreateMovie is a function that creates a movie
func CreateMovie(c *gin.Context) {
	var reqMovie types.Movie
	if err := c.ShouldBindJSON(&reqMovie); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var movie models.Movie

	movie.ID = primitive.NewObjectID()
	movie.Type = reqMovie.Type
	movie.Adult = reqMovie.Adult
	movie.BackdropPath = reqMovie.BackdropPath
	movie.GenreIds = reqMovie.GenreIds
	movie.OriginalLanguage = reqMovie.OriginalLanguage
	movie.OriginalTitle = reqMovie.OriginalTitle
	movie.Overview = reqMovie.Overview
	movie.Popularity = reqMovie.Popularity
	movie.PosterPath = reqMovie.PosterPath
	movie.ReleaseDate = reqMovie.ReleaseDate
	movie.Title = reqMovie.Title
	movie.Video = reqMovie.Video
	movie.VoteAverage = reqMovie.VoteAverage
	movie.VoteCount = reqMovie.VoteCount
	movie.MovieID = reqMovie.MovieID
	movie.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	movie.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// movieCollection.FindOne(c.Request.Context(), bson.D{{"title", movie.Title}})

	foundTitle := movieCollection.FindOne(c.Request.Context(), bson.D{{Key: "title", Value: movie.Title}})
	if foundTitle.Err() == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message":  "Movie already exists",
			"status":   http.StatusConflict,
			"success":  false,
			"error":    true,
			"errorMsg": "Movie already exists",
		})
		return
	}
	_, err := movieCollection.InsertOne(c.Request.Context(), movie)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Movie not created",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie created",
		"movie":   movie,
		"status":  http.StatusCreated,
		"success": true,
		"error":   false,
	})
}

// UpdateMovie is a function that updates a movie
func UpdateMovie(c *gin.Context) {
	var reqMovie types.Movie
	if err := c.ShouldBindJSON(&reqMovie); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	movieId := c.Param("id")

	searchedmovie, err := movieCollection.Find(c.Request.Context(), bson.M{"mov": movieId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	if searchedmovie == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":  "Movie not found",
			"status":   http.StatusNotFound,
			"success":  false,
			"error":    true,
			"errorMsg": "Movie not found",
		})
		return
	}

	var movie models.Movie

	movie.ID = primitive.NewObjectID()
	movie.Type = reqMovie.Type
	movie.Adult = reqMovie.Adult
	movie.BackdropPath = reqMovie.BackdropPath
	movie.GenreIds = reqMovie.GenreIds
	movie.OriginalLanguage = reqMovie.OriginalLanguage
	movie.OriginalTitle = reqMovie.OriginalTitle
	movie.Overview = reqMovie.Overview
	movie.Popularity = reqMovie.Popularity
	movie.PosterPath = reqMovie.PosterPath
	movie.ReleaseDate = reqMovie.ReleaseDate
	movie.Title = reqMovie.Title
	movie.Video = reqMovie.Video
	movie.VoteAverage = reqMovie.VoteAverage
	movie.VoteCount = reqMovie.VoteCount
	movie.MovieID = reqMovie.MovieID
	movie.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	movie.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = movieCollection.UpdateOne(c.Request.Context(), bson.M{"movie_id": movieId}, bson.M{"$set": movie})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Movie not updated",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie updated",
		"movie":   movie,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})

}

// DeleteMovie is a function that deletes a movie
func DeleteMovie(c *gin.Context) {
	movieId := c.Param("id")

	_, err := movieCollection.DeleteOne(c.Request.Context(), bson.M{"movie_id": movieId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error deleting movie",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// SearchMovie is a function that searches for a movie
func SearchMovie(c *gin.Context) {
	var movies []models.Movie

	title := c.Query("title")
	originalLanguage := c.Query("original_language")
	originalTitle := c.Query("original_title")

	var filter = bson.M{"title": bson.M{"$regex": title}, "original_language": bson.M{"$regex": originalLanguage}, "original_title": bson.M{"$regex": originalTitle}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movies by search",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})

}

// GetMovieByRating is a function that returns all the movies by rating
func GetMovieByRating(c *gin.Context) {
	var movies []models.Movie

	rating := c.Param("rating")

	var filter = bson.M{"voteaverage": bson.M{"$gte": rating}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movies by rating",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// GetMovieByYear is a function that returns all the movies by year
func GetMovieByYear(c *gin.Context) {
	var movies []models.Movie

	year := c.Param("year")

	var filter = bson.M{"releasedate": bson.M{"$regex": year}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movies by year",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})

}

// get popular movies
func GetPopularMovies(c *gin.Context) {
	var movies []models.Movie

	var filter = bson.M{"type": bson.M{"$eq": "popular"}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Popular movies",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// get upcoming movies
func GetUpcomingMovies(c *gin.Context) {
	var movies []models.Movie

	var filter = bson.M{"type": bson.M{"$eq": "upcoming"}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upcoming movies",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// get now playing movies
func GetNowPlayingMovies(c *gin.Context) {
	var movies []models.Movie

	var filter = bson.M{"type": bson.M{"$eq": "now_playing"}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Now playing movies",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})

}

// get top rated movies
func GetTopRatedMovies(c *gin.Context) {
	var movies []models.Movie
	var filter = bson.M{"voteaverage": bson.M{"$gte": 6.5}, "votecount": bson.M{"$gte": 1000}}

	cursor, err := movieCollection.Find(c.Request.Context(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var movie models.Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movies",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Top rated movies",
		"movies":  movies,
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}
