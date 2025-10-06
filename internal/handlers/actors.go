package handlers

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/Hosam-Zidany/dvdrental/internal/models"
	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	var actors []models.Actor
	result := database.DB.Find(&actors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, actors)
}

func GetActorById(c *gin.Context) {
	id := c.Param("id")
	var actor models.Actor
	result := database.DB.First(&actor, "actor_id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func GetFilmsByActor(c *gin.Context) {
	id := c.Param("id")
	var films []models.Film
	result := database.DB.Raw(
		`SELECT f.film_id,f.title, f.description, f.rating
		FROM film f
		JOIN film_actor fa ON f.film_id = fa.film_id
		WHERE fa.actor_id = ?`, id).Scan(&films)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
