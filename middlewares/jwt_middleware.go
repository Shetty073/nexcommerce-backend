package middlewares

import (
	responses "nexcommerce/common"
	"nexcommerce/utils/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Unauthorized(c, "Missing auth header", "Authorization header is missing")
			c.Abort()
			return
		}

		// authParts := strings.Split(authHeader, " ")
		// if len(authParts) != 2 || authParts[0] != "Bearer" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString := authParts[1]

		tokenString := authHeader

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.Configs.Jwt.Secret), nil
		})

		if err != nil || !token.Valid {
			responses.Unauthorized(c, "Invalid token", "Authorization token is invalid")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			responses.Unauthorized(c, "Invalid token claims", "Authorization token claims are invalid")
			c.Abort()
			return
		}

		// You can access the claims here if needed
		username, ok := claims["username"].(string)
		if !ok {
			responses.Unauthorized(c, "Invalid username claim", "Username claim in token is invalid")
			c.Abort()
			return
		}

		// Optionally, set the claims or user information in the context for downstream handlers
		c.Set("username", username)

		// Continue processing the request
		c.Next()
	}
}
