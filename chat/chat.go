package chat

import (
	"alfred/database"

	"net/http"
)

func GetAllChat(userId string, parterId string) Response {
	var chat *[]database.Chat

	db := database.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Where("(\"targetId\" = ? AND \"senderId\" = ?) OR (\"targetId\" = ? AND \"senderId\" = ?)", userId, parterId, parterId, userId).Find(&chat)

	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: response.Error.Error()}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Chat Not Found"}
	}

	return Response{Code: http.StatusOK, Response: chat}
}
