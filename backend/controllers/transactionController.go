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

var transactionCollection *mongo.Collection = database.GetCollection(database.DB, "transaction")

// Get all transactions
func GetAllTransactions(c *gin.Context) {
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
			{Key: "ticket_id", Value: 1},
			{Key: "user_id", Value: 1},
			{Key: "quantity", Value: 1},
			{Key: "total", Value: 1},
			{Key: "paid", Value: 1},
			{Key: "ticket", Value: 1},
			{Key: "booking", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
		}},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := transactionCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})
	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var transactions []bson.M
	if err := result.All(c.Request.Context(), &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":          page,
		"recordPerPage": recordPerPage,
		"transactions":  transactions,
		"message":       "All transactions",
		"error":         false,
		"total":         len(transactions),
	})
}

// Get a single transaction
func GetTransaction(c *gin.Context) {
	var transaction models.Transaction

	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = transactionCollection.FindOne(c.Request.Context(), bson.M{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction found",
		"error":       false,
		"transaction": transaction,
		"success":     true,
	})
}

// Create a transaction
func CreateTransaction(c *gin.Context) {
	var reqtransaction types.Transaction

	if err := c.ShouldBindJSON(&reqtransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	var transaction models.Transaction
	transaction.ID = primitive.NewObjectID()
	transaction.TicketID = reqtransaction.TicketID
	transaction.UserID = reqtransaction.UserID
	transaction.Quantity = reqtransaction.Quantity
	transaction.Total = reqtransaction.Total
	transaction.Paid = reqtransaction.Paid
	transaction.Ticket = make([]models.Ticket, len(reqtransaction.Ticket))
	for i, ticket := range reqtransaction.Ticket {
		transaction.Ticket[i] = models.Ticket{
			ID:       ticket.ID,
			Name:     ticket.Name,
			Date:     ticket.Date,
			Time:     ticket.Time,
			Day:      ticket.Day,
			Price:    ticket.Price,
			Location: ticket.Location,
			Seats:    ticket.Seats,
			Row:      ticket.Row,
			Theatre:  ticket.Theatre,
		}
	}
	transaction.Booking = make([]models.Booking, len(reqtransaction.Booking))
	for i, booking := range reqtransaction.Booking {
		transaction.Booking[i] = models.Booking{

			TransactionId: booking.TransactionId,
			UserId:        booking.UserId,
			SeatId:        booking.SeatId,
			ScreeningId:   booking.ScreeningId,
			Status:        booking.Status,
		}
	}

	transaction.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	transaction.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := transactionCollection.InsertOne(c.Request.Context(), transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Transaction not created",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction created",
		"error":       false,
		"success":     true,
		"transaction": transaction,
	})

}

// Update a transaction
func UpdateTransaction(c *gin.Context) {
	var reqtransaction types.Transaction

	if err := c.ShouldBindJSON(&reqtransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request",
			"success": false,
		})
		return
	}

	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transaction models.Transaction
	transaction.ID = transactionID
	transaction.TicketID = reqtransaction.TicketID
	transaction.UserID = reqtransaction.UserID
	transaction.Quantity = reqtransaction.Quantity
	transaction.Total = reqtransaction.Total
	transaction.Paid = reqtransaction.Paid
	transaction.Ticket = make([]models.Ticket, len(reqtransaction.Ticket))
	for i, ticket := range reqtransaction.Ticket {
		transaction.Ticket[i] = models.Ticket{
			Name:     ticket.Name,
			Date:     ticket.Date,
			Time:     ticket.Time,
			Day:      ticket.Day,
			Price:    ticket.Price,
			Location: ticket.Location,
			Seats:    ticket.Seats,
			Row:      ticket.Row,
			Theatre:  ticket.Theatre,
		}
	}
	transaction.Booking = make([]models.Booking, len(reqtransaction.Booking))
	for i, booking := range reqtransaction.Booking {
		transaction.Booking[i] = models.Booking{

			TransactionId: booking.TransactionId,
			UserId:        booking.UserId,
			SeatId:        booking.SeatId,
			ScreeningId:   booking.ScreeningId,
			Status:        booking.Status,
		}
	}

	transaction.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = transactionCollection.ReplaceOne(c.Request.Context(), bson.M{"_id": transactionID}, transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Transaction not updated",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction updated",
		"error":       false,
		"success":     true,
		"transaction": transaction,
	})
}

// Delete a transaction
func DeleteTransaction(c *gin.Context) {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = transactionCollection.DeleteOne(c.Request.Context(), bson.M{"_id": transactionID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction deleted",
		"error":   false,
		"success": true,
	})
}

// Get all transactions by user
func GetTransactionByUser(c *gin.Context) {
	var transactions []models.Transaction
	transactionUser := c.Param("user_id")

	cursor, err := transactionCollection.Find(c.Request.Context(), bson.M{"user_id": transactionUser})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the transaction",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var transaction models.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the transaction",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transaction found",
		"error":        false,
		"success":      true,
		"transactions": transactions,
	})
}

// Confirm a transaction
func ConfirmTransaction(c *gin.Context) {
	var transaction models.Transaction

	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = transactionCollection.FindOne(c.Request.Context(), bson.M{"_id": transactionID}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction.Paid = true

	_, err = transactionCollection.UpdateOne(c.Request.Context(), bson.M{"_id": transactionID}, transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction confirmed",
		"error":       false,
		"success":     true,
		"transaction": transaction,
	})
}

// Cancel a transaction
func CancelTransaction(c *gin.Context) {
	var transaction models.Transaction

	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = transactionCollection.FindOne(c.Request.Context(), bson.M{"_id": transactionID}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transaction.Paid = false

	_, err = transactionCollection.UpdateOne(c.Request.Context(), bson.M{"_id": transactionID}, transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction cancelled",
		"error":       false,
		"success":     true,
		"transaction": transaction,
	})
}

// Get all transactions by status
func GetTransactionByStatus(c *gin.Context) {
	var transactions []models.Transaction

	transactionStatus := c.Param("status")

	cursor, err := transactionCollection.Find(c.Request.Context(), bson.M{"paid": transactionStatus})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the transaction",
			"success": false,
			"data":    nil,
		})
		return
	}

	defer cursor.Close(c.Request.Context())

	for cursor.Next(c.Request.Context()) {
		var transaction models.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error occurred while trying to find the transaction",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transaction found",
		"error":        false,
		"success":      true,
		"transactions": transactions,
	})
}
