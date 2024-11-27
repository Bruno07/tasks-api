package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController struct {
	UserService services.UserService
	Validator   *validator.Validate
}

func NewUserController() *UserController {
	return &UserController{
		UserService: services.UserService{},
		Validator:   validator.New(),
	}
}

func (uc UserController) Create(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create!",
		})

		ctx.Abort()

		return
	}

	var userRequest requests.UserRequest
	json.Unmarshal(body, &userRequest)

	if err := uc.Validator.Struct(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	response, err := uc.UserService.Create(userRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create!",
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    response,
		"message": "Successfully created!",
	})

}
