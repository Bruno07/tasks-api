package services

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/http/responses"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct{}

func (jws JWTService) generateToken(responseJWT responses.JWTResponse) (*string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     responseJWT.User.ID,
		"name":        responseJWT.User.Name,
		"email":       responseJWT.User.Email,
		"permissions": responseJWT.User.Permissions,
		"exp":         responseJWT.ExpiresAt,
		"iss":         responseJWT.ISS,
	})

	signedString, err := token.SignedString([]byte(config.GetJWTSecret()))

	if err != nil {
		return nil, errors.New("Failed to sign the token!")
	}

	return &signedString, nil
}
