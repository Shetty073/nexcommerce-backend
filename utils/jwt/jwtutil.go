package jwt

import (
	"nexcommerce/utils/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a new JWT token with the provided claims
func GenerateToken(user_id string, expiration int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expiry := time.Hour * time.Duration(expiration)
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(expiry).Unix()

	tokenString, err := token.SignedString([]byte(config.Configs.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
