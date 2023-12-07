package notification

import (
	"alfred/database"
	"net/http"
)

func GetAllNotifications(userId string) Response {
	var notification *database.Notification

	db := database.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Find(&notification, "\"targetId\"", userId)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: response.Error.Error()}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "No Notifications"}
	}

	return Response{Code: http.StatusOK, Response: notification}
}
