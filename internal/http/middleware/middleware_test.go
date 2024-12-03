package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Bruno07/tasks-api/internal/http/auth"
	"github.com/Bruno07/tasks-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	
	gin.SetMode(gin.TestMode)
	t.Setenv("JWT_SECRET", "924a46c0284440a9ab1fc62763d6aa69")

	var jw = auth.JWT{}
    var auth = models.User{ID: 1, Name: "Test JWT", Email: "testjwt@email.com"}
    
	var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJydW5vLmNhbnV0b0Bob3RtYWlsLmNvbSIsImV4cCI6MTczMzE5MDc3OCwiaXNzIjoidGFza3MtYXBpIiwibmFtZSI6IkJydW5vIEZlcm5hbmRlcyBDYW51dG8iLCJwZXJtaXNzaW9ucyI6WyJDUkVBVEUiLCJVUERBVEUiLCJWSUVXIl0sInVzZXJfaWQiOjF9.H-EnRliqGjtXUQdeMJI3ISf5ZvXqP8LCabS9cGWJkAY"

    signedString, _ := jw.GenerateToken(
		&auth,
		time.Now().Add(1 * time.Hour).Unix(),
		"tasks-api",
		[]string{"tasks:create", "tasks:update", "tasks:view"},
	)

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing token",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Token not provided!"}`,
		},
		{
			name:           "Invalid token",
			authHeader:     "Bearer invalid_token",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid token!"}`,
		},
		{
			name:           "Invalid or expired token",
			authHeader:     "Bearer " + expiredToken,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid token!"}`,
		},
		{
			name:           "Valid token",
			authHeader:     "Bearer " + signedString,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(AuthMiddleware())
			router.GET("/api/tasks", func(c *gin.Context) {
				c.String(http.StatusOK, "Success")
			})

			req := httptest.NewRequest("GET", "/api/tasks", nil)
			req.Header.Set("Authorization", tt.authHeader)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}

}