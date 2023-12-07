package job

import (
	"alfred/database"
	"alfred/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateJob(data *CreateJobBody, clientID string) Response {
	var job database.Jobs
	db := database.InitDB()

	sqlDB, _ := db.DB()
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

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error creating job"}
	}
	return Response{Code: http.StatusCreated, Response: job}
}

func GetJobByID(jobID string) Response {
	var job database.Jobs
	db := database.InitDB()
	sqlDB, _ := db.DB()
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

	sqlDB, _ := db.DB()
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

	sqlDB, _ := db.DB()
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

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Find(&job)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get all jobs"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Job not found"}
	}
	return Response{Code: http.StatusOK, Response: job}
}

func GetValueOrDefault(value interface{}, defaultValue interface{}) interface{} {
	switch value.(type) {
	case string:
		if value != "" {
			return value
		}
	case float64:
		if value != 0.0 || value != nil {
			return value
		}
	}
	return defaultValue
}

func EditJobById(c *gin.Context, jobID string, data EditJobBody) Response {
	var job database.Jobs

	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.First(&job, "ID=?", jobID)
	middleware.AuthorizationMiddleware(c, job.ClientID.String())
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get job"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Job not found"}
	}

	job.Name = GetValueOrDefault(data.Name, job.Name).(string)
	job.Descriptions = GetValueOrDefault(data.Descriptions, job.Descriptions).(string)
	job.Address = GetValueOrDefault(data.Address, job.Address).(string)
	job.Latitude = GetValueOrDefault(data.Latitude, job.Latitude).(float64)
	job.Longitude = GetValueOrDefault(data.Longitude, job.Longitude).(float64)
	job.ImageURL = GetValueOrDefault(data.ImageURL, job.ImageURL).(string)
	edited := db.Save(&job)
	if edited.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: edited.Error.Error()}
	}
	if edited.RowsAffected == 0 {
		return Response{Code: http.StatusBadRequest, Response: "Error Update data"}
	}
	return Response{
		Code:     http.StatusOK,
		Response: job,
	}
}

func SetTalent(c *gin.Context, data SetTalentBody) Response {
	var job database.Jobs

	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.First(&job, "ID=?", data.JobID)
	middleware.AuthorizationMiddleware(c, job.ClientID.String())
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error get job data"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "Job not found"}
	}

	parsedTalentID, err := uuid.Parse(data.TalentID)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing talent ID"}
	}

	job.TalentID = &parsedTalentID
	job.FixedPrice = data.FixedPrice
	edited := db.Save(&job)
	if edited.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: edited.Error.Error()}
	}
	if edited.RowsAffected == 0 {
		return Response{Code: http.StatusBadRequest, Response: "Error Edit job data"}
	}
	return Response{Code: http.StatusOK, Response: job}
}
