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

var ticketCollention *mongo.Collection = database.GetCollection(database.DB, "ticket")

// GetAllTickets is a function that returns all the tickets
func GetAllTickets(c *gin.Context) {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errp := strconv.Atoi(c.Query("page"))
	if errp != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{
		{Key: "name", Value: 1},
		{Key: "date", Value: 1},
		{Key: "time", Value: 1},
		{Key: "day", Value: 1},
		{Key: "price", Value: 1},
		{Key: "location", Value: 1},
		{Key: "seats", Value: 1},
		{Key: "row", Value: 1},
		{Key: "theatre", Value: 1},
		{Key: "created_at", Value: 1},
		{Key: "updated_at", Value: 1},
	}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := movieDetailsCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var allTickets []bson.M
	if err = result.All(c.Request.Context(), &allTickets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":        allTickets,
		"message":       "All tickets",
		"page":          page,
		"recordPerPage": recordPerPage,
	})

}

// GetTicket is a function that returns a single ticket
func GetTicket(c *gin.Context) {
	var ticket models.Ticket

	ticketID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ticketCollention.FindOne(c.Request.Context(), bson.M{"_id": ticketID}).Decode(&ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  ticket,
		"message": "Ticket found",
	})
}

// CreateTicket is a function that creates a ticket
func CreateTicket(c *gin.Context) {
	var reqTicket types.Ticket

	if err := c.ShouldBindJSON(&reqTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ticket models.Ticket

	ticket.ID = primitive.NewObjectID()
	ticket.Name = reqTicket.Name
	ticket.Date = reqTicket.Date
	ticket.Time = reqTicket.Time
	ticket.Day = reqTicket.Day
	ticket.Price = reqTicket.Price
	ticket.Location = reqTicket.Location
	ticket.Seats = reqTicket.Seats
	ticket.Row = reqTicket.Row
	ticket.Theatre = reqTicket.Theatre
	ticket.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	ticket.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := ticketCollention.InsertOne(c.Request.Context(), ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  ticket,
		"message": "Ticket created",
	})

}

// UpdateTicket is a function that updates a ticket
func UpdateTicket(c *gin.Context) {
	var reqTicket types.Ticket

	if err := c.ShouldBindJSON(&reqTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketId := c.Param("id")

	searchedTicket, err := ticketCollention.Find(c.Request.Context(), bson.M{"id": ticketId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ticket"})
		return
	}

	if searchedTicket == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	var ticket models.Ticket

	ticket.ID = primitive.NewObjectID()
	ticket.Name = reqTicket.Name
	ticket.Date = reqTicket.Date
	ticket.Time = reqTicket.Time
	ticket.Day = reqTicket.Day
	ticket.Price = reqTicket.Price
	ticket.Location = reqTicket.Location
	ticket.Seats = reqTicket.Seats
	ticket.Row = reqTicket.Row
	ticket.Theatre = reqTicket.Theatre
	ticket.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	ticket.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = ticketCollention.UpdateOne(c.Request.Context(), bson.M{"id": ticketId}, bson.M{"$set": ticket})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  ticket,
		"message": "Ticket updated",
	})
}

// DeleteTicket is a function that deletes a ticket
func DeleteTicket(c *gin.Context) {
	ticketId := c.Param("id")

	_, err := ticketCollention.DeleteOne(c.Request.Context(), bson.M{"id": ticketId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket deleted",
	})
}

// GetTicketByBooking is a function that returns all the tickets by booking
func GetTicketByBooking(c *gin.Context) {
	var tickets []models.Ticket

	bookingID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cursor, err := ticketCollention.Find(c.Request.Context(), bson.M{"booking_id": bookingID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tickets"})
		return
	}

	if err = cursor.All(c.Request.Context(), &tickets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tickets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  tickets,
		"message": "Tickets found",
	})
}
