package services

import (
	"testing"

	"github.com/Bruno07/tasks-api/internal/http/requests"
	"github.com/Bruno07/tasks-api/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {

	t.Run("Must return a user", func(t *testing.T) {

		userService := NewUserService(
			repositories.UserMockRepository{},
			repositories.ProfileMockRepository{},
			repositories.PermissionMockRepository{},
			repositories.ModuleMockRepository{},
		)

		response, err := userService.Create(requests.UserRequest{
			Name:      "Master",
			Email:     "master@email.com",
			Password:  "secret",
			ProfileId: 1,
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("Should return error", func(t *testing.T) {

		userService := NewUserService(
			repositories.UserMockRepository{},
			repositories.ProfileMockRepository{},
			repositories.PermissionMockRepository{},
			repositories.ModuleMockRepository{},
		)

		response, err := userService.Create(requests.UserRequest{
			Name:      "Master",
			Email:     "master@email.com",
			Password:  "secret",
			ProfileId: 3,
		})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

}
