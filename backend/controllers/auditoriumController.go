package controllers

import (
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

var auditoriumCollection *mongo.Collection = database.GetCollection(database.DB, "movies")

// GetAllAuditoriums is a function that returns all the auditoriums
func GetAllAuditoriums(c *gin.Context) {
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
			{Key: "movie_id", Value: 1},
			{Key: "no_seats", Value: 1},
			{Key: "rows", Value: 1},
			{Key: "columns", Value: 1},
			{Key: "seats", Value: 1},
		}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := movieDetailsCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var auditorium []bson.M

	if err := result.All(c.Request.Context(), &auditorium); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditoriums",
		"data":    auditorium,
		"status":  true,
		"page":    page,
		"record":  recordPerPage,
		"total":   len(auditorium),
		"error":   false,
	})

}

// GetAuditorium is a function that returns a single auditorium
func GetAuditorium(c *gin.Context) {
	var auditorium models.Auditorium

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": id}

	err := auditoriumCollection.FindOne(c.Request.Context(), filter).Decode(&auditorium)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditorium",
		"success": true,
		"data":    auditorium,
		"error":   false,
	})

}

// CreateAuditorium is a function that creates an auditorium
func CreateAuditorium(c *gin.Context) {
	var reqAuditorium types.Auditorium

	if err := c.ShouldBindJSON(&reqAuditorium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid data",
			"success": false,
			"data":    nil,
		})
		return
	}

	var auditorium models.Auditorium

	auditorium.ID = primitive.NewObjectID()
	auditorium.Name = reqAuditorium.Name
	auditorium.MovieID = reqAuditorium.MovieID
	auditorium.NoSeats = reqAuditorium.NoSeats
	auditorium.Rows = reqAuditorium.Rows
	auditorium.Columns = reqAuditorium.Columns
	auditorium.Seats = make([]models.Seat, len(reqAuditorium.Seats))
	for i, seat := range reqAuditorium.Seats {
		auditorium.Seats[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	auditorium.SelectedSeats = make([]models.Seat, len(reqAuditorium.SelectedSeats))
	for i, seat := range reqAuditorium.SelectedSeats {
		auditorium.SelectedSeats[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	auditorium.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	auditorium.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := auditoriumCollection.InsertOne(c.Request.Context(), auditorium)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not created",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditorium created",
		"success": true,
		"data":    auditorium,
		"error":   false,
	})
}

// UpdateAuditorium is a function that updates an auditorium
func UpdateAuditorium(c *gin.Context) {
	var reqAuditorium types.Auditorium

	if err := c.ShouldBindJSON(&reqAuditorium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid data",
			"success": false,
			"data":    nil,
		})
		return
	}

	auditoriumID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	searchedAuditorium, err := auditoriumCollection.Find(c.Request.Context(), bson.M{"_id": auditoriumID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	if searchedAuditorium == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
			"error":   false,
		})
		return
	}

	var auditorium models.Auditorium

	auditorium.ID = auditoriumID
	auditorium.Name = reqAuditorium.Name
	auditorium.MovieID = reqAuditorium.MovieID
	auditorium.NoSeats = reqAuditorium.NoSeats
	auditorium.Rows = reqAuditorium.Rows
	auditorium.Columns = reqAuditorium.Columns
	auditorium.Seats = make([]models.Seat, len(reqAuditorium.Seats))
	for i, seat := range reqAuditorium.Seats {
		auditorium.Seats[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	auditorium.SelectedSeats = make([]models.Seat, len(reqAuditorium.SelectedSeats))
	for i, seat := range reqAuditorium.SelectedSeats {
		auditorium.SelectedSeats[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	auditorium.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = auditoriumCollection.UpdateOne(c.Request.Context(), bson.M{"_id": auditoriumID}, auditorium)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not updated",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditorium updated",
		"success": true,
		"data":    auditorium,
		"error":   false,
	})

}

// DeleteAuditorium is a function that deletes an auditorium
func DeleteAuditorium(c *gin.Context) {
	auditorium := c.Param("id")

	_, err := auditoriumCollection.DeleteOne(c.Request.Context(), bson.M{"_id": auditorium})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not deleted",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditorium deleted",
		"success": true,
		"error":   false,
	})
}

// GetAuditoriumByName is a function that returns all the auditoriums by name
func GetAuditoriumByName(c *gin.Context) {
	var auditoriums []models.Auditorium
	name := c.Param("name")

	cursor, err := auditoriumCollection.Find(c.Request.Context(), bson.M{"name": name})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var auditorium models.Auditorium
		cursor.Decode(&auditorium)
		auditoriums = append(auditoriums, auditorium)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Auditorium search by name",
		"success": true,
		"data":    auditoriums,
		"error":   false,
	})
}

// GetAuditoriumBySeatingCapacity is a function that returns all the auditoriums by seating capacity
func GetAuditoriumBySeatingCapacity(c *gin.Context) {
	var auditoriums []models.Auditorium

	seatingCapacity, _ := strconv.Atoi(c.Param("seatingCapacity"))

	cursor, err := auditoriumCollection.Find(c.Request.Context(), bson.M{"no_seats": seatingCapacity})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var auditorium models.Auditorium
		cursor.Decode(&auditorium)
		auditoriums = append(auditoriums, auditorium)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Auditorium search by seating capacity",
		"success": true,
		"data":    auditoriums,
		"error":   false,
	})
}

// GetSeatArrangement is a function that returns the seat arrangement of an auditorium
func GetSeatArrangement(c *gin.Context) {
	var auditorium models.Auditorium

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": id}

	err := auditoriumCollection.FindOne(c.Request.Context(), filter).Decode(&auditorium)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seat arrangement",
		"success": true,
		"data":    auditorium.Seats,
		"error":   false,
	})
}
