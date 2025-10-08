package routers

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/auth"
	"github.com/Hosam-Zidany/dvdrental/internal/handlers"
	"github.com/Hosam-Zidany/dvdrental/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	api := r.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", auth.Register)
			authGroup.POST("/login", auth.Login)
		}

		api.GET("/films", handlers.GetFilms)
		api.GET("/films/:id", handlers.GetFilmByID)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id/films", handlers.GetFilmsByCategory)
		api.GET("/actors", handlers.GetActors)
		api.GET("/actors/:id", handlers.GetActorById)
		api.GET("/actors/:id/films", handlers.GetFilmsByActor)

		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/profile", auth.GetProfile)
			// protected.POST("/rentals", handlers.CreateRental)
			// protected.GET("/payments", handlers.GetUserPayments)
		}
	}
	return r
}
