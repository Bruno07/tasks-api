package auth

import (
	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

func (j *JWT) GenerateToken(
	auth *models.User,
	expiresAt int64,
	iss string,
	permissions []string,
) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     auth.ID,
		"name":        auth.Name,
		"email":       auth.Email,
		"permissions": permissions,
		"exp":         expiresAt,
		"iss":         iss,
	})

	signedString, err := token.SignedString([]byte(config.GetJWTSecret()))

	return signedString, err

}
