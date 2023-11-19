package client

import (
	"alfred/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterClient(data *RegisterBody) Response {
	var client database.Client
	db := database.InitDB()
	saltStr := os.Getenv("HASH_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error converting salt"}
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
		return Response{Code: http.StatusInternalServerError, Response: "Error creating user"}
	}
	return Response{Code: http.StatusCreated, Response: client}
}

func LoginClient(data *LoginBody) Response {
	var client database.Client
	db := database.InitDB()
	jwtKey := os.Getenv("JWT_KEY")
	jwtByte := []byte(jwtKey)

	err := db.First(&client, "email=?", data.Email).Error
	if err != nil {
		return Response{Code: http.StatusNotFound, Response: "Client not Found"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(data.Password))
	if err != nil {
		return Response{Code: http.StatusUnauthorized, Response: "Password mismatch"}
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
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error signing token"}
	}
	response := PayloadResponse{
		ID:    client.ID.String(),
		Token: tokenString,
		Role:  "client",
	}
	return Response{Code: http.StatusOK, Response: response}
}

func GetClientByID(clientID string) Response {
	var client database.Client
	db := database.InitDB()
	err := db.First(&client, "ID=?", clientID).Error
	if err != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Client not found",
		}
	}
	client.Password = ""
	return Response{
		Code:     http.StatusOK,
		Response: client,
	}
}

func EditClientData(clientID string) {}
