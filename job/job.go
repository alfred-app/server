package job

import (
	"alfred/database"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func CreateJob(data *CreateJobBody, clientID string) Response {
	var job database.Jobs
	db := database.InitDB()

	parsedID, err := uuid.Parse(clientID)

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing client ID"}
	}
	new, err := uuid.NewUUID()
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error generating ID"}
	}

	job = database.Jobs{
		ID:           new,
		ClientID:     parsedID,
		Name:         data.Name,
		Descriptions: data.Descriptions,
		Address:      data.Address,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		ImageURL:     data.ImageURL,
	}

	err = db.Create(&job).Error

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error creating job"}
	}
	return Response{Code: http.StatusCreated, Response: job}
}

func GetJobByID(jobID string) Response {
	var job database.Jobs
	db := database.InitDB()
	err := db.First(&job, "ID=?", jobID).Error
	fmt.Println(job)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get the job"}
	}
	return Response{Code: http.StatusOK, Response: job}
}

func GetAllJobs() Response {
	var job []database.Jobs

	db := database.InitDB()
	err := db.Find(&job).Error
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error gett all jobs"}
	}
	fmt.Println(job)
	return Response{Code: http.StatusOK, Response: job}
}
