package controllers

import (
	"main/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateInput struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password" binding:"required"`
	Confirm     string `json:"confirm" binding:"required"`
}

type UserLoginInput struct {
	Id       int
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	return string(hashedPasswordBytes), err
}

func matchPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(c *gin.Context) {
	input := UserCreateInput{}
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match, err := regexp.MatchString("^[a-zA-Z0-9_]*$", input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username can only contain alphanumeric characters and underscores."})
		return
	}

	passwordlen := len(input.Password)
	if passwordlen < 8 || passwordlen > 64 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be between 8 and 64 characters."})
		return
	}

	if input.Password != input.Confirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The 2 passwords do not match."})
		return
	}

	if input.OldPassword == input.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot be the same as old password."})
		return
	}

	hash, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.OldPassword == "" {
		query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id`

		id := 0
		row := models.DB.QueryRow(query, input.Username, hash).Scan(&id)

		if row != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Account created successfully."})
	} else {
		query := `
		UPDATE users
		SET password = ($2)
		WHERE username = ($1)
		RETURNING id`

		id := 0
		row := models.DB.QueryRow(query, input.Username, hash).Scan(&id)

		if row != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Password changed successfully."})
	}
}

func LoginUser(c *gin.Context) {
	input := UserLoginInput{}
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usernamelen := len(input.Username)
	if usernamelen < 4 || usernamelen > 24 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be between 4 and 24 characters."})
		return
	}

	match, err := regexp.MatchString("^[a-zA-Z0-9_]*$", input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username can only contain alphanumeric characters and underscores."})
		return
	}

	query := `SELECT * FROM users WHERE username = ($1)`

	var user models.User
	row := models.DB.QueryRow(query, input.Username).Scan(&user.Id, &user.Username, &user.Password)

	if row != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Username or Password"})
		return
	}

	if !matchPassword(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Username or Password"})
		return
	}

	token, err := generateJwt(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token[0], "refresh": token[1]})
}

func CheckValidation(c *gin.Context) {
	username := validateUser(c)
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": username})
}
