package controllers

import (
	"microblog/application/ports/user"
	"microblog/interface/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related requests.
type UserController struct {
	FollowUserUseCase user.FollowUserPort
}

// FollowUser allows a user to follow another user.
func (uc *UserController) FollowUser(c *gin.Context) {
	var request dto.FollowUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.FollowUserUseCase.Execute(request.UserID, request.TargetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
}
