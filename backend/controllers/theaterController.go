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

var theaterCollection *mongo.Collection = database.GetCollection(database.DB, "theater")

// GetAllTheaters is a function that returns all the theaters
func GetAllTheaters(c *gin.Context) {
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
			{Key: "name", Value: 1},
			{Key: "address", Value: 1},
			{Key: "auditorium", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
		}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := theaterCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})
	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var theaters []bson.M
	if err := result.All(c.Request.Context(), &theaters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":          page,
		"recordPerPage": recordPerPage,
		"screenings":    theaters,
		"message":       "All screenings",
		"error":         false,
		"total":         len(theaters),
	})

}

// GetTheater is a function that returns a single theater
func GetTheater(c *gin.Context) {
	var theater models.Theater

	theaterID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = theaterCollection.FindOne(c.Request.Context(), bson.M{"_id": theaterID}).Decode(&theater)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while trying to find the theater", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Theater found",
		"error":   false,
		"theater": theater,
		"success": true,
	})
}

// CreateTheater is a function that creates a theater
func CreateTheater(c *gin.Context) {
	var reqtheater types.Theater

	if err := c.ShouldBindJSON(&reqtheater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	var theater models.Theater
	theater.ID = primitive.NewObjectID()
	theater.Name = reqtheater.Name
	theater.Address = reqtheater.Address
	theater.Auditorium = make([]models.Auditorium, len(reqtheater.Auditorium))
	for i, auditorium := range reqtheater.Auditorium {
		theater.Auditorium[i] = models.Auditorium{
			Name:          auditorium.Name,
			Rows:          auditorium.Rows,
			Columns:       auditorium.Columns,
			Seats:         make([]models.Seat, len(auditorium.Seats)),
			SelectedSeats: make([]models.Seat, len(auditorium.SelectedSeats)),
		}
	}
	theater.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	theater.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := theaterCollection.InsertOne(c.Request.Context(), theater)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Theater not created",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Theater created",
		"error":   false,
		"success": true,
		"theater": theater,
	})
}

// UpdateTheater is a function that updates a theater
func UpdateTheater(c *gin.Context) {
	var reqtheater types.Theater

	if err := c.ShouldBindJSON(&reqtheater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	theaterID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	searchedTheater, err := theaterCollection.Find(c.Request.Context(), bson.M{"_id": theaterID})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if searchedTheater == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Theater not found",
			"success": false,
			"error":   true,
		})
		return
	}

	var theater models.Theater
	theater.ID = theaterID
	theater.Name = reqtheater.Name
	theater.Address = reqtheater.Address
	theater.Auditorium = make([]models.Auditorium, len(reqtheater.Auditorium))
	for i, auditorium := range reqtheater.Auditorium {
		theater.Auditorium[i] = models.Auditorium{
			Name:          auditorium.Name,
			Rows:          auditorium.Rows,
			Columns:       auditorium.Columns,
			Seats:         make([]models.Seat, len(auditorium.Seats)),
			SelectedSeats: make([]models.Seat, len(auditorium.SelectedSeats)),
		}
	}
	theater.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = theaterCollection.ReplaceOne(c.Request.Context(), bson.M{"_id": theaterID}, theater)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Theater not updated",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Theater updated",
		"error":   false,
		"success": true,
		"theater": theater,
	})
}

// DeleteTheater is a function that deletes a theater
func DeleteTheater(c *gin.Context) {
	theaterID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = theaterCollection.DeleteOne(c.Request.Context(), bson.M{"_id": theaterID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while trying to delete the theater", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Theater deleted",
		"error":   false,
		"success": true,
	})
}

// GetTheaterByName is a function that returns all the theaters by name
func GetTheaterByName(c *gin.Context) {
	var theaters []models.Theater
	theaterName := c.Param("name")

	cursor, err := theaterCollection.Find(c.Request.Context(), bson.M{"name": theaterName})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var theater models.Theater
		cursor.Decode(&theater)
		theaters = append(theaters, theater)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Theater found",
		"error":    false,
		"success":  true,
		"theaters": theaters,
	})

}

// GetTheaterByLocation is a function that returns all the theaters by location
func GetTheaterByLocation(c *gin.Context) {
	var theaters []models.Theater
	theaterAddress := c.Param("address")

	cursor, err := theaterCollection.Find(c.Request.Context(), bson.M{"address": theaterAddress})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var theater models.Theater
		cursor.Decode(&theater)
		theaters = append(theaters, theater)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Theater found",
		"error":    false,
		"success":  true,
		"theaters": theaters,
	})

}

// GetTheaterBySeatingCapacity is a function that returns all the theaters by seating capacity
func GetTheaterBySeatingCapacity(c *gin.Context) {
	var theaters []models.Theater
	seatingCapacity, err := strconv.Atoi(c.Param("seatingCapacity"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cursor, err := theaterCollection.Find(c.Request.Context(), bson.M{"auditorium.seats": bson.M{"$gt": seatingCapacity}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var theater models.Theater
		cursor.Decode(&theater)
		theaters = append(theaters, theater)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the theater",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Theater found",
		"error":    false,
		"success":  true,
		"theaters": theaters,
	})
}
