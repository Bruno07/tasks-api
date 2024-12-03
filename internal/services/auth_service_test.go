package services

import (
	"testing"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/Bruno07/tasks-api/internal/repositories"
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

	service := NewAuthService(userRepo)
	
	t.Run("Must return a token with an expiration", func(t *testing.T) {

		token, expiresAt, err := service.Login("tec1@email.com", "12345678")

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.NotEmpty(t, expiresAt)

	})

	t.Run("Should not return a token", func(t *testing.T) {

		token, expiresAt, err := service.Login("tec1@email.com", "123456789")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Empty(t, expiresAt)

	})

}