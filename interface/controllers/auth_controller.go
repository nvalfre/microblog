package controllers

import (
	"microblog/interface/dto"
	"microblog/security/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related requests.
type AuthController struct{}

// GetToken generates a signed JWT token for the specified user ID.
func (ac *AuthController) GetToken(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
		return
	}

	token, err := auth.GenerateSignedToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := dto.TokenResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}
