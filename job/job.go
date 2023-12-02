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

	sqlDB, err := db.DB()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: err.Error()}
	}

	defer sqlDB.Close()

	parsedID, err := uuid.Parse(clientID)

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing client ID"}
	}

	new, err := uuid.NewUUID()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error generating id"}
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
	fmt.Println(job)

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error creating job"}
	}
	return Response{Code: http.StatusCreated, Response: job}
}

func GetJobByID(jobID string) Response {
	var job database.Jobs
	db := database.InitDB()
	sqlDB, err := db.DB()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: err.Error()}
	}

	defer sqlDB.Close()
	response := db.First(&job, "ID=?", jobID)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get the job"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Job not found"}
	}
	return Response{Code: http.StatusOK, Response: job}
}

func GetJobByClientID(clientID string) Response {
	var job []database.Jobs
	db := database.InitDB()
	sqlDB, err := db.DB()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: err.Error()}
	}

	defer sqlDB.Close()
	response := db.Model(&database.Jobs{}).Find(&job, "\"clientID\"=?", clientID)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get the job"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "No Jobs Found"}
	}
	return Response{Code: http.StatusOK, Response: job}
}

func GetJobByTalentID(talentID string) Response {
	var job []database.Jobs
	db := database.InitDB()
	sqlDB, err := db.DB()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: err.Error()}
	}

	defer sqlDB.Close()
	response := db.Model(&database.Jobs{}).Find(&job, "\"talentID\"=?", talentID)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get the job"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "No Jobs Found"}
	}
	return Response{
		Code:     http.StatusOK,
		Response: job,
	}
}

func GetAllJobs() Response {
	var job []database.Jobs

	db := database.InitDB()
	sqlDB, err := db.DB()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: err.Error()}
	}

	defer sqlDB.Close()
	response := db.Find(&job)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get all jobs"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Job not found"}
	}
	fmt.Println(job)
	return Response{Code: http.StatusOK, Response: job}
}
