package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to authenticate!",
		})

		c.Abort()

		return
	}

	var request = &requests.UserRequestDTO{}
	json.Unmarshal(body, request)

	token, expiresAt, err := ac.authService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to authenticate!",
		})

		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"expires_at": expiresAt,
	})

}
