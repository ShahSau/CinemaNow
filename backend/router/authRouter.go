package router

import (
	"github.com/ShahSau/CinemaNow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r *routes) AuthRoutes(rg *gin.RouterGroup) {
	authRouteGrouping := rg.Group("/auth")
	authRouteGrouping.Use(cors.Default())

	// register
	authRouteGrouping.POST("/register", controllers.Register)

	// login
	authRouteGrouping.POST("/login", controllers.Login)

	// logout
	authRouteGrouping.POST("/logout", controllers.Logout)

	// forgot password
	authRouteGrouping.POST("/forgotpassword", controllers.ForgotPassword)

	// reset password
	authRouteGrouping.POST("/resetpassword", controllers.ResetPassword)
}
