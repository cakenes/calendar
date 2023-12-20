package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"main/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	Access   []byte `json:"access"`
}

func generateGcm(key [32]byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

func createToken(key []byte) ([]byte, error) {
	k := sha256.Sum256(key)

	gcm, err := generateGcm(k)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	uuid := []byte(uuid.New().String())

	println(string(uuid))

	return gcm.Seal(nonce, nonce, uuid, nil), nil
}

func openToken(ciphertext []byte, key []byte) ([]byte, error) {
	k := sha256.Sum256(key)

	gcm, err := generateGcm(k)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("error: Malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
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

		// current password checking missing

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
	row := models.DB.QueryRow(query, input.Username).Scan(&user.Id, &user.Username, &user.Password, &user.Access)

	if row != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Username or Password"})
		return
	}

	if !matchPassword(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Username or Password"})
		return
	}

	access, err := openToken(user.Access, []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := generateJwt(user.Username, string(access))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": jwt[0]})
}

func CheckValidation(c *gin.Context) {
	if validateUser(c) {
		c.JSON(http.StatusOK, gin.H{"success": "Valid token."})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token."})
}
