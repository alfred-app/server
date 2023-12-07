package main

import (
	"alfred/bidlist"
	"alfred/chat"
	"alfred/client"
	"alfred/job"
	"alfred/middleware"
	"alfred/notification"
	"alfred/talent"

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
	notificationGroup := router.Group("/notification")
	chatGroup := router.Group("/chat")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, world!")
	})

	clientGroup.POST("/register", client.RegisterHandler)
	clientGroup.POST("/login", client.LoginHandler)
	clientGroup.GET("/:userID", client.GetClientData)
	clientGroup.PATCH("/:userID", middleware.AuthenticationMiddleware, client.UpdateHandler)
	clientGroup.PATCH("/change-password/:userID", middleware.AuthenticationMiddleware, client.ChangePasswordHandler)

	talentGroup.POST("/register", talent.RegisterHandler)
	talentGroup.POST("/login", talent.LoginHandler)
	talentGroup.GET("/:userID", talent.GetTalentData)
	talentGroup.PATCH("/:userID", middleware.AuthenticationMiddleware, talent.UpdateHandler)
	talentGroup.PATCH("/change-password/:userID", middleware.AuthenticationMiddleware, talent.ChangePasswordHandler)
	talentGroup.DELETE("/:userID", middleware.AuthenticationMiddleware, talent.DeleteHandler)

	jobGroup.GET("/all", job.GetAllJobHandler)
	jobGroup.GET("/client/:clientID", middleware.AuthenticationMiddleware, job.GetJobByClientIDHandler)
	jobGroup.GET("/talent/:talentID", middleware.AuthenticationMiddleware, job.GetJobByTalentIDHandler)
	jobGroup.POST("/create-job/:userID", middleware.AuthenticationMiddleware, middleware.ClientGuard, job.CreateJobHandler)
	jobGroup.GET("/:jobID", middleware.AuthenticationMiddleware, job.GetJobByIDHandler)
	jobGroup.PATCH("/:jobID", middleware.AuthenticationMiddleware, job.EditJobByJobIDHandler)
	jobGroup.POST("/set-talent", middleware.AuthenticationMiddleware, job.SetTalentHandler)

	bidlistGroup.GET("/job/:jobID", middleware.AuthenticationMiddleware, bidlist.GetBidListByJobIDHandler)
	bidlistGroup.GET("/:bidListID", middleware.AuthenticationMiddleware, bidlist.GetBidListByIDHandler)
	bidlistGroup.POST("/create/:jobID", middleware.AuthenticationMiddleware, middleware.TalentGuard, bidlist.CreateBidListHandler)
	bidlistGroup.DELETE("/:bidListID", middleware.AuthenticationMiddleware, middleware.TalentGuard, bidlist.DeleteBidListHandler)

	notificationGroup.GET("/:userId", middleware.AuthenticationMiddleware, notification.GetAllNotificationHandler)

	chatGroup.GET("/:userId/:targetId", chat.GetAllChatHandler)

	router.Run()
}
