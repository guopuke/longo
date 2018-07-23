package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

// Context is the context of the JSON web token.
type Context struct {
	ID       uint64
	Username string
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
