package client

import (
	"github.com/golang-jwt/jwt/v5"
)

type RegisterBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	ImageURL    string `json:"image_url"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditClientBody struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	ImageURL    string `json:"image_url"`
}

type ChangePasswordBody struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type Payload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}

type PayloadResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
