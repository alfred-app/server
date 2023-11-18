package talent

import "github.com/golang-jwt/jwt/v5"

type RegisterBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	AboutMe     string `json:"about_me"`
	ImageURL    string `json:"image_url"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Payload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type RegisterResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}

type PayloadResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

type LoginResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}