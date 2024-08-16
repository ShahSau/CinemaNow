package controllers

import "github.com/gin-gonic/gin"

// Get all transactions
func GetAllTransactions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "All transactions",
	})
}

// Get a single transaction
func GetTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Single transaction",
	})
}

// Create a transaction
func CreateTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create transaction",
	})
}

// Update a transaction
func UpdateTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update transaction",
	})
}

// Delete a transaction
func DeleteTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete transaction",
	})
}

// Get all transactions by user
func GetTransactionByUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transactions by user",
	})
}

// Confirm a transaction
func ConfirmTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Confirm transaction",
	})
}

// Cancel a transaction
func CancelTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Cancel transaction",
	})
}

// Get all transactions by status
func GetTransactionByStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transactions by status",
	})
}

// Get all transactions by date
func GetTransactionByDate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transactions by date",
	})
}
