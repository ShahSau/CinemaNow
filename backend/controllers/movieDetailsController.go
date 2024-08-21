package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ShahSau/CinemaNow/backend/database"
	"github.com/ShahSau/CinemaNow/backend/models"
	"github.com/ShahSau/CinemaNow/backend/types"
)

var movieDetailsCollection *mongo.Collection = database.GetCollection(database.DB, "moviesDetails")

// GetAllMovieDetails is a function that returns all the movie details
func GetAllMovieDetails(c *gin.Context) {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errp := strconv.Atoi(c.Query("page"))
	if errp != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, _ = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "id", Value: 1},
				{Key: "adult", Value: 1},
				{Key: "backdrop_path", Value: 1},
				{Key: "belongs_to_collection", Value: 1},
				{Key: "budget", Value: 1},
				{Key: "genres", Value: 1},
				{Key: "homepage", Value: 1},
				{Key: "imdb_id", Value: 1},
				{Key: "movie_id", Value: 1},
				{Key: "original_country", Value: 1},
				{Key: "original_language", Value: 1},
				{Key: "original_title", Value: 1},
				{Key: "overview", Value: 1},
				{Key: "popularity", Value: 1},
				{Key: "poster_path", Value: 1},
				{Key: "production_companies", Value: 1},
				{Key: "production_countries", Value: 1},
				{Key: "release_date", Value: 1},
				{Key: "revenue", Value: 1},
				{Key: "runtime", Value: 1},
				{Key: "spoken_languages", Value: 1},
				{Key: "status", Value: 1},
				{Key: "tagline", Value: 1},
				{Key: "title", Value: 1},
				{Key: "video", Value: 1},
				{Key: "vote_average", Value: 1},
				{Key: "vote_count", Value: 1},
				{Key: "created_at", Value: 1},
				{Key: "updated_at", Value: 1},
			},
		},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := movieDetailsCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var allMovieDetails []bson.M
	if err = result.All(c.Request.Context(), &allMovieDetails); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": allMovieDetails, "page": page, "recordPerPage": recordPerPage, "startIndex": startIndex})
}

// GetMovieDetails is a function that returns a single movie details by its ID
func GetMovieDetails(c *gin.Context) {
	var movieDetails models.MovieDetails

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": id}

	err := movieDetailsCollection.FindOne(c.Request.Context(), filter).Decode(&movieDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details found",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}

// CreateMovieDetail is a function that creates a new movie detail
func CreateMovieDetail(c *gin.Context) {
	var reqMovie types.MovieDetails

	if err := c.ShouldBindJSON(&reqMovie); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var movieDetails models.MovieDetails

	movieDetails.ID = primitive.NewObjectID()
	movieDetails.Adult = reqMovie.Adult
	movieDetails.BackdropPath = reqMovie.BackdropPath
	movieDetails.BelongsToCollection = reqMovie.BelongsToCollection
	movieDetails.Budget = reqMovie.Budget
	movieDetails.Genres = reqMovie.Genres
	movieDetails.Homepage = reqMovie.Homepage
	movieDetails.ImdbID = reqMovie.ImdbID
	movieDetails.MovieID = reqMovie.MovieID
	movieDetails.OriginalCountry = reqMovie.OriginalCountry
	movieDetails.OriginalLanguage = reqMovie.OriginalLanguage
	movieDetails.OriginalTitle = reqMovie.OriginalTitle
	movieDetails.Overview = reqMovie.Overview
	movieDetails.Popularity = reqMovie.Popularity
	movieDetails.PosterPath = reqMovie.PosterPath
	movieDetails.ProductionCompanies = reqMovie.ProductionCompanies
	movieDetails.ProductionCountries = reqMovie.ProductionCountries
	movieDetails.ReleaseDate = reqMovie.ReleaseDate
	movieDetails.Revenue = reqMovie.Revenue
	movieDetails.Runtime = reqMovie.Runtime
	movieDetails.SpokenLanguages = reqMovie.SpokenLanguages
	movieDetails.Status = reqMovie.Status
	movieDetails.Tagline = reqMovie.Tagline
	movieDetails.Title = reqMovie.Title
	movieDetails.Video = reqMovie.Video
	movieDetails.VoteAverage = reqMovie.VoteAverage
	movieDetails.VoteCount = reqMovie.VoteCount
	movieDetails.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	movieDetails.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	foundTitle := movieDetailsCollection.FindOne(c.Request.Context(), bson.M{"title": movieDetails.Title})
	if foundTitle.Err() == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Movie already exists",
			"status":  http.StatusBadRequest,
			"success": false,
			"error":   true,
		})
		return
	}

	_, err := movieDetailsCollection.InsertOne(c.Request.Context(), movieDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error creating movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details created",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}

// UpdateMovieDetail is a function that updates a movie detail by its ID
func UpdateMovieDetail(c *gin.Context) {
	var reqMovie types.MovieDetails

	if err := c.ShouldBindJSON(&reqMovie); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")

	searchedMovieDetails, err := movieDetailsCollection.Find(c.Request.Context(), bson.M{"movie_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	if searchedMovieDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":  "Movie details not found",
			"status":   http.StatusNotFound,
			"success":  false,
			"error":    true,
			"errorMsg": "Movie details not found",
		})
		return
	}

	var movieDetails models.MovieDetails

	movieDetails.ID = primitive.NewObjectID()
	movieDetails.Adult = reqMovie.Adult
	movieDetails.BackdropPath = reqMovie.BackdropPath
	movieDetails.BelongsToCollection = reqMovie.BelongsToCollection
	movieDetails.Budget = reqMovie.Budget
	movieDetails.Genres = reqMovie.Genres
	movieDetails.Homepage = reqMovie.Homepage
	movieDetails.ImdbID = reqMovie.ImdbID
	movieDetails.MovieID = reqMovie.MovieID
	movieDetails.OriginalCountry = reqMovie.OriginalCountry
	movieDetails.OriginalLanguage = reqMovie.OriginalLanguage
	movieDetails.OriginalTitle = reqMovie.OriginalTitle
	movieDetails.Overview = reqMovie.Overview
	movieDetails.Popularity = reqMovie.Popularity
	movieDetails.PosterPath = reqMovie.PosterPath
	movieDetails.ProductionCompanies = reqMovie.ProductionCompanies
	movieDetails.ProductionCountries = reqMovie.ProductionCountries
	movieDetails.ReleaseDate = reqMovie.ReleaseDate
	movieDetails.Revenue = reqMovie.Revenue
	movieDetails.Runtime = reqMovie.Runtime
	movieDetails.SpokenLanguages = reqMovie.SpokenLanguages
	movieDetails.Status = reqMovie.Status
	movieDetails.Tagline = reqMovie.Tagline
	movieDetails.Title = reqMovie.Title
	movieDetails.Video = reqMovie.Video
	movieDetails.VoteAverage = reqMovie.VoteAverage
	movieDetails.VoteCount = reqMovie.VoteCount
	movieDetails.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	movieDetails.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = movieDetailsCollection.InsertOne(c.Request.Context(), movieDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error updating movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details updated",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}

// DeleteMovieDetail is a function that deletes a movie detail by its ID
func DeleteMovieDetail(c *gin.Context) {
	id := c.Param("id")

	_, err := movieDetailsCollection.DeleteOne(c.Request.Context(), bson.M{"movie_id": id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error deleting movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details deleted",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
	})
}

// SearchMovieDetailByIMDBID is a function that searches for a movie detail by its IMDB ID
func SearchMovieDetailByIMDBID(c *gin.Context) {
	imdbId := c.Param("imdbId")

	var movieDetails models.MovieDetails

	err := movieDetailsCollection.FindOne(c.Request.Context(), bson.M{"imdb_id": imdbId}).Decode(&movieDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details found",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}

// SearchMovieDetailByGenre is a function that searches for a movie detail by its genre
func SearchMovieDetailByGenre(c *gin.Context) {
	genre := c.Param("genre")

	var movieDetails []models.MovieDetails

	cursor, err := movieDetailsCollection.Find(c.Request.Context(), bson.M{"genres": genre})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	if err = cursor.All(c.Request.Context(), &movieDetails); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details found",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}

// SearchMovieDetailByProductionCompany is a function that searches for a movie detail by its production company
func SearchMovieDetailByProductionCompany(c *gin.Context) {
	productionCompany := c.Param("productionCompany")

	var movieDetails []models.MovieDetails

	cursor, err := movieDetailsCollection.Find(c.Request.Context(), bson.M{"production_companies": productionCompany})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Error fetching movie details",
			"status":   http.StatusInternalServerError,
			"success":  false,
			"error":    true,
			"errorMsg": err.Error(),
		})
		return
	}

	if err = cursor.All(c.Request.Context(), &movieDetails); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie details found",
		"status":  http.StatusOK,
		"success": true,
		"error":   false,
		"data":    movieDetails,
	})
}
