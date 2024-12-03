package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/Bruno07/tasks-api/internal/requests"
	"github.com/Bruno07/tasks-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {

	userRepo := &repositories.MockUserRepository{
		MockGetByEmail: func(email string) (*models.User, error) {

			var err error
			var result = models.User{}

			var usersGroup = map[int64]models.User{}
			usersGroup[1] = models.User{Email: "master@email.com", Password: "$2a$10$6u.SUGSdniAsLvYkFEMiUe7EPvzB9/1PVoUK7ulb4AgpPMz7Afud2"}
			usersGroup[2] = models.User{Email: "tec1@email.com", Password: "$2a$10$sG6jyu0CI57WTs2aoGYiIuA9afsuQUK58iE3e/s5gtbmtkuRXZT5K"}
			usersGroup[3] = models.User{Email: "tec2@email.com", Password: "$2a$10$JvMfg6GZyPTrGaZ3pj8WxekCx06QjjPGK9Ndo4HS60ZKFA59Rd2By"}

			for _, userGroup := range usersGroup {
				if email == userGroup.Email {
					result = userGroup
				}
			}

			return &result, err
		},
	}

	service := services.NewAuthService(userRepo)
	controller := NewAuthController(*service)

	t.Run("Should return a status of 200", func(t *testing.T) {

		var request = requests.UserRequestDTO{
			Email:    "tec1@email.com",
			Password: "12345678",
		}
		body, _ := json.Marshal(request)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodDelete, "/login", bytes.NewReader(body))
		c.Request = r

		controller.Login(c)

		assert.Equal(t, http.StatusOK, w.Code)

	})

}
