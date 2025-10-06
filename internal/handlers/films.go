package handlers

import (
	"net/http"
	"strconv"

	"github.com/Hosam-Zidany/dvdrental/internal/models"

	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/gin-gonic/gin"
)

func GetFilms(c *gin.Context) {
	var films []models.Film
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	limit, _ := strconv.Atoi((c.DefaultQuery("limit", "10")))
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Film{})

	category := c.Query("category")
	language := c.Query("language")

	if language != "" {
		query.Joins("JOIN language l ON l.language_id = film.language_id").
			Where("l.name = ?", language)
	}
	if category != "" {
		query.Joins("JOIN film_category fc ON fc.film_id = film.film_id").
			Joins("JOIN category c ON c.category_id = fc.category_id").
			Where("c.name = ?", category)
	}
	// result := database.DB.Find(&films)
	result := query.Offset(offset).Limit(limit).Find(&films)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"count": result.RowsAffected,
		"data":  films,
	})
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
