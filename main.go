package main

import (
	"alfred/client"
	"alfred/job"
	"alfred/middleware"
	"alfred/talent"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	clientGroup := router.Group("/client")
	talentGroup := router.Group("/talent")
	jobGroup := router.Group("/job")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, world!")
	})

	clientGroup.POST("/register", client.RegisterHandler)
	clientGroup.POST("/login", client.LoginHandler)
	clientGroup.GET("/:userID", client.GetClientData)
	clientGroup.PATCH("/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, client.UpdateHandler)
	clientGroup.PATCH("/change-password/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, client.ChangePasswordHandler)

	talentGroup.POST("/register", talent.RegisterHandler)
	talentGroup.POST("/login", talent.LoginHandler)
	talentGroup.GET("/:userID", talent.GetTalentData)
	talentGroup.PATCH("/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, talent.UpdateHandler)
	talentGroup.PATCH("/change-password/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, talent.ChangePasswordHandler)
	talentGroup.DELETE("/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, talent.DeleteHandler)

	jobGroup.GET("/", job.GetAllJobHandler)
	jobGroup.POST("/create-job/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, middleware.ClientGuard, job.CreateJobHandler)
	jobGroup.GET("/:jobID", middleware.AuthenticationMiddleware, job.GetJobByIDHandler)
	router.Run()
}
