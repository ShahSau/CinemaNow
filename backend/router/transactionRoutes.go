package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TransactionRoutes is a function that handles all the routes for the transaction
func (r *routes) TransactionRoutes(rg *gin.RouterGroup) {
	transaction := rg.Group("/auditorium")
	transaction.Use(cors.Default())

	// get all transactions
	transaction.GET("/all", controllers.GetAllTransactions)

	// get a single transaction
	transaction.GET("/:id", controllers.GetTransaction)

	// create a transaction
	transaction.POST("/create", controllers.CreateTransaction)

	// update a transaction
	transaction.PUT("/:id", controllers.UpdateTransaction)

	// delete a transaction
	transaction.DELETE("/:id", controllers.DeleteTransaction)

	// get all transactions by user
	transaction.GET("/user/:id", controllers.GetTransactionByUser)

	// confirm a transaction
	transaction.PUT("/confirm/:id", controllers.ConfirmTransaction)

	// cancel a transaction
	transaction.PUT("/cancel/:id", controllers.CancelTransaction)

	// get all transactions by status
	transaction.GET("/status/:status", controllers.GetTransactionByStatus)

	// get all transactions by date
	transaction.GET("/date/:date", controllers.GetTransactionByDate)

}
