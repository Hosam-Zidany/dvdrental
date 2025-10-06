package handlers

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/Hosam-Zidany/dvdrental/internal/models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func GetFilmsByCategory(c *gin.Context) {
	id := c.Param("id")
	var films []models.Film
	result := database.DB.Raw(
		`SELECT f.film_id,f.title, f.description, f.rating
		FROM film f
		JOIN film_category fc ON f.film_id = fc.film_id
		WHERE fc.category_id = ?`, id).Scan(&films)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
