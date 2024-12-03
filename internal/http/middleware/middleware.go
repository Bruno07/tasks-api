package middleware

import (
	"net/http"
	"strings"

	"github.com/Bruno07/tasks-api/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Token not provided!"})
			
			ctx.Abort()
			
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token!"})
			
			ctx.Abort()
			
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(config.GetJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token!"})
			
			ctx.Abort()
			
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token!"})
			
			ctx.Abort()
			
			return
		}

		ctx.Set("user_id", claims["user_id"])
		ctx.Set("permissions", claims["permissions"])

		ctx.Next()
	}
}
