package client

import (
	"alfred/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}

func RegisterClient(data *RegisterBody) (RegisterResponse, error) {
	var client database.Client
	db := database.InitDB()
	saltStr, isExist := os.LookupEnv("HASH_SALT")
	if !isExist {
		log.Fatal("Environment variable HASH_SALT is not set")
	}
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		log.Fatal("Error converting salt string")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), salt)
	data.Password = string(hashedPassword)
	client = database.Client{
		Email:       data.Email,
		Name:        data.Name,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		ImageURL:    data.ImageURL,
	}
	err = db.Create(&client).Error
	client.Password = ""
	if err != nil {
		return RegisterResponse{Code: 500, Response: "Error creating user"}, err
	}
	return RegisterResponse{Code: http.StatusCreated, Response: client}, err
}

func LoginClient(data *LoginBody) (LoginResponse, error) {
	var client database.Client
	db := database.InitDB()
	jwtKey, isExist := os.LookupEnv("JWT_KEY")
	jwtByte := []byte(jwtKey)
	if !isExist {
		log.Fatal("JWT_KEY not found")
	}

	err := db.First(&client, "email=?", data.Email).Error
	if err != nil {
		log.Fatal("Client not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(data.Password))
	if err != nil {
		log.Fatal("Password mismatch")
	}
	expirationsTime := time.Now().Add(24 * time.Hour)
	payload := &Payload{
		ID:    client.ID.String(),
		Email: client.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationsTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(jwtByte)
	response := PayloadResponse{
		ID:    client.ID.String(),
		Token: tokenString,
		Role:  "client",
	}
	return LoginResponse{Code: http.StatusOK, Response: response}, err
}