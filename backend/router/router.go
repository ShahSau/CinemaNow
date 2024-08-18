package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

// ClientRoutes is a function that handles all the routes for the client side
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
	r.router.Use(cors.Default())

	// swagger TODO
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// swagger docs TODO

	v := r.router.Group((os.Getenv("API_VERSION")))

	// all v1 routes
	r.MovieRoutes(v)
	r.MovieDetailsRoutes(v)
	r.AuthRoutes(v)
	r.AuditoriumRoutes(v)
	r.UserRoutes(v)
	r.TransactionRoutes(v)
	r.ScreeningRoutes(v)
	//r.SeatRoutes(v)
	r.BookingsRoutes(v)
	r.TicketsRoutes(v)
	r.TheaterRoutes(v)

	r.router.Run(":" + os.Getenv("PORT"))
}
