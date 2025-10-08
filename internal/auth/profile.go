package auth

import (
	"net/http"

	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/Hosam-Zidany/dvdrental/internal/models"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	val, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id missing in context"})
		return
	}
	uid := val.(uint)
	var user models.User
if err := database.DB.First(&user, "id = ?", uid).Error; err != nil {
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	return
}


	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	})
}
