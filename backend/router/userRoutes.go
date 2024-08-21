package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that handles all the routes for the user
func (r *routes) UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	user.Use(cors.Default())

	// get all users
	user.GET("/all", controllers.GetAllUsers)

	// get a single user
	user.GET("/:id", controllers.GetUser)

	// create a user
	//user.POST("/create", controllers.CreateUser)

	// update a user
	user.PUT("/:id", controllers.UpdateUser)

	// delete a user
	user.DELETE("/:id", controllers.DeleteUser)

}
