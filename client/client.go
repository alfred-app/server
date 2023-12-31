package client

import (
	"alfred/database"
	"alfred/middleware"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterClient(data *RegisterBody) Response {
	var client database.Client
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	saltStr := os.Getenv("HASH_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error converting salt"}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), salt)
	data.Password = string(hashedPassword)
	new, err := uuid.NewUUID()
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error generating ID"}
	}
	client = database.Client{
		ID:          new,
		Email:       data.Email,
		Name:        data.Name,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		ImageURL:    data.ImageURL,
	}
	response := db.Create(&client)
	client.Password = ""
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error creating user"}
	}
	return Response{Code: http.StatusCreated, Response: client}
}

func LoginClient(data *LoginBody) Response {
	var client database.Client
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	jwtKey := os.Getenv("JWT_KEY")
	jwtByte := []byte(jwtKey)

	isExist := db.First(&client, "email=?", data.Email)
	if isExist.Error != nil {
		return Response{Code: http.StatusNotFound, Response: "Client not Found"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(data.Password))
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

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Model(&database.Client{}).Preload("Jobs").Find(&client, "ID=?", clientID)
	if response.Error != nil {
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

func GetValueOrDefault(value string, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}

func EditClientData(c *gin.Context, clientID string, data *EditClientBody) Response {
	var client database.Client
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.First(&client, "ID=?", clientID)
	middleware.AuthorizationMiddleware(c, client.ID.String())
	if response.Error != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Client not found",
		}
	}
	client.Email = GetValueOrDefault(data.Email, client.Email)
	client.Name = GetValueOrDefault(data.Name, client.Name)
	client.Address = GetValueOrDefault(data.Address, client.Address)
	client.PhoneNumber = GetValueOrDefault(data.PhoneNumber, client.PhoneNumber)
	client.ImageURL = GetValueOrDefault(data.ImageURL, client.ImageURL)
	edited := db.Save(&client)
	if edited.Error != nil {
		return Response{
			Code:     http.StatusNotImplemented,
			Response: "Failed to update data",
		}
	}
	client.Password = ""
	return Response{
		Code:     http.StatusOK,
		Response: client,
	}
}

func ChangePassword(c *gin.Context, clientID string, data *ChangePasswordBody) Response {
	var client database.Client
	saltStr := os.Getenv("HASH_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		return Response{
			Code:     http.StatusInternalServerError,
			Response: "Error converting salt",
		}
	}
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.First(&client, "ID=?", clientID)
	middleware.AuthorizationMiddleware(c, client.ID.String())

	if response.Error != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Client does not exist",
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(data.OldPassword))
	if err != nil {
		return Response{
			Code:     http.StatusUnauthorized,
			Response: "Password mismatch",
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), salt)
	client.Password = string(hashedPassword)
	changed := db.Save(&client)
	if changed.Error != nil {
		return Response{
			Code:     http.StatusInternalServerError,
			Response: "Error update password",
		}
	}
	client.Password = ""
	return Response{
		Code:     http.StatusAccepted,
		Response: client,
	}
}
