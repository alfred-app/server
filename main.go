package main

import (
	"alfred/client"
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

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, world!")
	})

	clientGroup.POST("/register", client.RegisterHandler)
	clientGroup.POST("/login", client.LoginHandler)
	clientGroup.GET("/:clientID", middleware.AuthenticationMiddleware, middleware.AuthorizationMiddleware, client.GetClientData)
	clientGroup.PATCH("/:clientID", client.UpdateHandler)
	clientGroup.PATCH("/change-password/:clientID", client.ChangePasswordHandler)

	talentGroup.POST("/register", talent.RegisterHandler)
	talentGroup.POST("/login", talent.LoginHandler)

	router.Run()
}
