package talent

import (
	"alfred/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterTalent(data *RegisterBody) Response {
	var talent database.Talent
	db := database.InitDB()
	saltStr := os.Getenv("HASH_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error converting salt"}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), salt)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error generating hash password"}
	}
	data.Password = string(hashedPassword)
	talent = database.Talent{
		Email:       data.Email,
		Name:        data.Name,
		Password:    data.Password,
		AboutMe:     data.AboutMe,
		ImageURL:    data.ImageURL,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
	}
	err = db.Create(&talent).Error
	talent.Password = ""
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error creating user"}
	}
	return Response{Code: http.StatusCreated, Response: talent}
}

func LoginTalent(data *LoginBody) Response {
	var talent database.Talent
	db := database.InitDB()
	jwtKey := os.Getenv("JWT_KEY")
	jwtByte := []byte(jwtKey)

	err := db.First(&talent, "email=?", data.Email).Error
	if err != nil {
		return Response{Code: http.StatusNotFound, Response: "Talent not found"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(talent.Password), []byte(data.Password))
	if err != nil {
		return Response{Code: http.StatusUnauthorized, Response: "Password mismatch"}
	}
	expirationsTime := time.Now().Add(24 * time.Hour)
	payload := &Payload{
		ID:    talent.ID.String(),
		Email: talent.Email,
		Role:  "talent",
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
		ID:    talent.ID.String(),
		Token: tokenString,
		Role:  "talent",
	}
	return Response{Code: http.StatusOK, Response: response}
}

func GetValueOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func GetTalentByID(talentID string) Response {
	var talent database.Talent
	db := database.InitDB()
	err := db.First(&talent, "ID=?", talentID).Error
	if err != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Talent not found",
		}
	}
	talent.Password = ""
	return Response{
		Code:     http.StatusOK,
		Response: talent,
	}
}

func GetTalents() Response {
	var talents []database.Talent
	db := database.InitDB()
	err := db.Find(&talents).Error
	if err != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Talents not found",
		}
	}
	return Response{
		Code:     http.StatusOK,
		Response: talents,
	}
}

func EditTalentData(talentID string, data *EditTalentBody) Response {
	var talent database.Talent
	db := database.InitDB()
	err := db.First(&talent, "ID=?", talentID).Error
	if err != nil {
		return Response{
			Code:     http.StatusNotFound,
			Response: "Talent not found",
		}
	}
	talent.Email = GetValueOrDefault(data.Email, talent.Email)
	talent.Name = GetValueOrDefault(data.Name, talent.Name)
	talent.AboutMe = GetValueOrDefault(data.AboutMe, talent.AboutMe)
	talent.ImageURL = GetValueOrDefault(data.ImageURL, talent.ImageURL)
	talent.Address = GetValueOrDefault(data.Address, talent.Address)
	talent.PhoneNumber = GetValueOrDefault(data.PhoneNumber, talent.PhoneNumber)
	err = db.Save(&talent).Error
	if err != nil {
		return Response{
			Code:     http.StatusNotImplemented,
			Response: "Failed to update data",
		}
	}
	talent.Password = ""
	return Response{
		Code:     http.StatusOK,
		Response: talent,
	}
}
