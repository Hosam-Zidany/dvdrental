package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hosam-Zidany/dvdrental/internal/config"
	"github.com/Hosam-Zidany/dvdrental/internal/database"
	"github.com/Hosam-Zidany/dvdrental/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email address"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}
	expHours := config.AppConfig.JWTExpHrs
	if expHours == 0 {
		expHours = 24
	}
	claim := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(user.ID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.AppConfig.JWTExpHrs) * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signed, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signed})
}
