package handlers

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/models"

	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/gin-gonic/gin"
)

func GetFilms(c *gin.Context) {
	var films []models.Film
	result := database.DB.Find(&films)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}

func GetFilmByID(c *gin.Context) {
	id := c.Param("id")
	var film models.Film
	result := database.DB.First(&film, "film_id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "film not found"})
		return
	}
	c.JSON(http.StatusOK, film)
}
