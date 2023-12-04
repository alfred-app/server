package main

import (
	"alfred/bidlist"
	"alfred/client"
	"alfred/job"
	"alfred/middleware"
	"alfred/talent"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORS)

	clientGroup := router.Group("/client")
	talentGroup := router.Group("/talent")
	jobGroup := router.Group("/job")
	bidlistGroup := router.Group("/bidlist")

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

	jobGroup.GET("/all", job.GetAllJobHandler)
	jobGroup.GET("/client/:clientID", job.GetJobByClientIDHandler)
	jobGroup.GET("/talent/:talentID", job.GetJobByTalentIDHandler)
	jobGroup.POST("/create-job/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, middleware.ClientGuard, job.CreateJobHandler)
	jobGroup.GET("/:jobID", middleware.AuthenticationMiddleware, job.GetJobByIDHandler)

	bidlistGroup.GET("/", middleware.AuthenticationMiddleware, bidlist.GetAllBidListHandler)
	bidlistGroup.GET("/:bidListID", middleware.AuthenticationMiddleware, bidlist.GetBidListByIDHandler)
	bidlistGroup.POST("/:create-bidlist/:userID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, middleware.ClientGuard, bidlist.CreateBidListHandler)
	bidlistGroup.DELETE("/:bidListID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, middleware.ClientGuard, bidlist.DeleteBidListHandler)

	router.Run()
}
