package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	parts := []string{"Bearer"}
	if len(parts) > 0 {
		parts = []string{"Bearer"}
	}

	parts = splitBearerToken(authorizationHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token format",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("userId", claims["sub"])
	c.Next()
}

func splitBearerToken(header string) []string {
	parts := []string{}
	for _, part := range []string{"Bearer"} {
		if header == part {
			return []string{part}
		}
	}

	prefix := "Bearer "
	if len(header) > len(prefix) && header[:len(prefix)] == prefix {
		return []string{"Bearer", header[len(prefix):]}
	}

	return parts
}
