package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero. ")
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

func ParseRequest(c *gin.Context) (*Context, error) {
	headerAuth := c.Request.Header.Get("Authorization")
	if len(headerAuth) == 0 {
		return &Context{}, ErrMissingHeader
	}

	secret := viper.GetString("jwt_secret")

	var token string
	fmt.Sscanf(headerAuth, "Bearer %s", &token)

	return Parse(token, secret)
}

// Parse validates the token with the specified secret.
// and returns the context if the token was valid.
func Parse(tokenString, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}
