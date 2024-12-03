package auth

import (
	"testing"
	"time"

	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestAuth_GenerateToken(t *testing.T) {

    var jw = JWT{}
    var auth = models.User{ID: 1, Name: "Test JWT", Email: "testjwt@email.com"}
    
    signedString, err := jw.generateToken(
		&auth,
		time.Now().Add(1 * time.Hour).Unix(),
		"tasks-api",
		[]string{"tasks:create", "tasks:update", "tasks:view"},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, signedString)

}