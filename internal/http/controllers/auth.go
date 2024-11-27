package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController() AuthController {
	return AuthController{
		AuthService: services.AuthService{},
	}
}

func (ac AuthController) Login(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to authenticate!",
		})

		ctx.Abort()

		return
	}

	var request requests.LoginRequest
	json.Unmarshal(body, &request)

	response, err := ac.AuthService.Authenticate(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to authenticate!",
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": response.AccessToken,
		"expires_at": response.ExpiresAt,
	})
	
}
