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

var bookingCollection *mongo.Collection = database.GetCollection(database.DB, "bookings")

// GetAllBookings is a function that returns all the bookings
func GetAllBookings(c *gin.Context) {
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
			{Key: "booking_id", Value: 1},
			{Key: "user_id", Value: 1},
			{Key: "screening_id", Value: 1},
			{Key: "transaction_id", Value: 1},
			{Key: "status", Value: 1},
			{Key: "seats", Value: 1},
			{Key: "total_price", Value: 1},
		}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := bookingCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var bookings []bson.M

	if err := result.All(c.Request.Context(), &bookings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Auditoriums",
		"data":    bookings,
		"status":  true,
		"page":    page,
		"record":  recordPerPage,
		"total":   len(bookings),
		"error":   false,
	})
}

// GetBooking is a function that returns a single booking
func GetBooking(c *gin.Context) {
	var booking models.Booking

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": id}

	err := bookingCollection.FindOne(c.Request.Context(), filter).Decode(&booking)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking",
		"success": true,
		"data":    booking,
		"error":   false,
	})
}

// CreateBooking is a function that creates a booking
func CreateBooking(c *gin.Context) {
	var reqBooking types.Booking

	if err := c.ShouldBindJSON(&reqBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid data",
			"success": false,
			"data":    nil,
		})
		return
	}

	var booking models.Booking

	booking.ID = primitive.NewObjectID()
	booking.UserId = reqBooking.UserId
	booking.TransactionId = reqBooking.TransactionId
	booking.SeatId = reqBooking.SeatId
	booking.ScreeningId = reqBooking.ScreeningId
	booking.Status = reqBooking.Status
	booking.Seat = make([]models.Seat, len(reqBooking.Seat))
	for i, seat := range reqBooking.Seat {
		booking.Seat[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	booking.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	booking.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := bookingCollection.InsertOne(c.Request.Context(), booking)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error creating booking",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created",
		"success": true,
		"data":    booking,
		"error":   false,
	})
}

// UpdateBooking is a function that updates a booking
func UpdateBooking(c *gin.Context) {
	var reqBooking types.Booking

	if err := c.ShouldBindBodyWithJSON(&reqBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid data",
			"success": false,
			"data":    nil,
		})
		return
	}

	bookingID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	searchedBooking, err := bookingCollection.Find(c.Request.Context(), bson.M{"_id": bookingID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Auditorium not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	if searchedBooking == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Booking not found",
			"success": false,
			"data":    nil,
			"error":   false,
		})
		return
	}

	var booking models.Booking
	booking.ID = bookingID
	booking.TransactionId = reqBooking.TransactionId
	booking.UserId = reqBooking.UserId
	booking.SeatId = reqBooking.SeatId
	booking.ScreeningId = reqBooking.ScreeningId
	booking.Status = reqBooking.Status
	booking.Seat = make([]models.Seat, len(reqBooking.Seat))
	for i, seat := range reqBooking.Seat {
		booking.Seat[i] = models.Seat{
			Row:       seat.Row,
			Number:    seat.Number,
			Available: seat.Available,
			Price:     seat.Price,
			Type:      seat.Type,
		}
	}
	booking.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = bookingCollection.UpdateOne(c.Request.Context(), bson.M{"_id": bookingID}, booking)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "booking not updated",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking Updated",
		"success": true,
		"data":    booking,
		"error":   false,
	})

}

// DeleteBooking is a function that deletes a booking
func DeleteBooking(c *gin.Context) {
	booking := c.Param("id")

	_, err := bookingCollection.DeleteOne(c.Request.Context(), bson.M{"_id": booking})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not deleted",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking deleted",
		"success": true,
		"error":   false,
	})
}

// GetBookingByUser is a function that returns all the bookings by user
func GetBookingByUser(c *gin.Context) {
	var bookings []models.Booking
	user := c.Param("user")
	cursor, err := bookingCollection.Find(c.Request.Context(), bson.M{"user_id": user})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Booking search by name",
		"success": true,
		"data":    bookings,
		"error":   false,
	})

}

// GetBookingByScreening is a function that returns all the bookings by screening
func GetBookingByScreening(c *gin.Context) {
	var bookings []models.Booking
	screening := c.Param("screening")
	cursor, err := bookingCollection.Find(c.Request.Context(), bson.M{"screening_id": screening})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Booking search by screening",
		"success": true,
		"data":    bookings,
		"error":   false,
	})
}

// GetBookingByTransaction is a function that returns all the bookings by transaction
func GetBookingByTransaction(c *gin.Context) {
	var bookings []models.Booking
	transaction := c.Param("transaction")
	cursor, err := bookingCollection.Find(c.Request.Context(), bson.M{"transaction_id": transaction})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Booking search by transaction",
		"success": true,
		"data":    bookings,
		"error":   false,
	})
}

// GetBookingByStatus is a function that returns all the bookings by status
func GetBookingByStatus(c *gin.Context) {
	var bookings []models.Booking
	status := c.Param("status")
	cursor, err := bookingCollection.Find(c.Request.Context(), bson.M{"status": status})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Booking not found",
			"success": false,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Booking search by status",
		"success": true,
		"data":    bookings,
		"error":   false,
	})
}
