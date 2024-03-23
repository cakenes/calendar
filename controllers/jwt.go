package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("supersecretkey") // TODO: Maybe hide me ?
var expiration = [1]time.Duration{15 * time.Minute}

type Claims struct {
	Username string `json:"username"`
	Access   string `json:"access"`
	jwt.RegisteredClaims
}

func generateJwt(username string, access string) ([1]string, error) {
	tokens := [1]string{}

	for i, exp := range expiration {
		expirationTime := time.Now().Add(exp)
		claims := &Claims{
			Username: username,
			Access:   access,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokens[i], _ = token.SignedString(key)
	}
	return tokens, nil
}

func validateUser(c *gin.Context) bool {
	tokenString := c.Request.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	c.JSON(http.StatusOK, gin.H{"token": token.Claims})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}
