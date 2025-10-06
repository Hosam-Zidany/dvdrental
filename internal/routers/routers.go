package routers

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	api := r.Group("/api")
	{
		api.GET("/films", handlers.GetFilms)
		api.GET("/films/:id", handlers.GetFilmByID)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id/films", handlers.GetFilmsByCategory)
		api.GET("/actors", handlers.GetActors)
		api.GET("/actors/:id", handlers.GetActorById)
		api.GET("/actors/:id/films", handlers.GetFilmsByActor)
	}
	return r
}
