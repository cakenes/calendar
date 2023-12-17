package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("supersecretkey")
var expiration = [2]time.Duration{15 * time.Minute, 24 * time.Hour}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateJwt(username string) ([2]string, error) {
	tokens := [2]string{}

	for i, exp := range expiration {
		expirationTime := time.Now().Add(exp)
		claims := &Claims{
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokens[i], _ = token.SignedString(key)
	}
	return tokens, nil
}

func validateUser(c *gin.Context) string {
	tokenString := c.Request.Header.Get("Authorization")
	tokenSplit := strings.Split(tokenString, " ")

	if len(tokenSplit) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token."})
		return ""
	}

	token, err := jwt.Parse(tokenSplit[0], func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	_, err1 := jwt.Parse(tokenSplit[1], func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return ""
	}

	//TODO: Generate new token if refresh token is valid

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token.Claims.(jwt.MapClaims)["username"].(string)
	}

	return ""
}
