package main

import (
	"alfred/client"
	"alfred/talent"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	clientGroup := router.Group("/client")
	talentGroup := router.Group("/talent")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, world!")
	})

	clientGroup.POST("/register", client.RegisterHandler)
	clientGroup.POST("/login", client.LoginHandler)
	clientGroup.GET("/:id", client.GetClientData)

	talentGroup.POST("/register", talent.RegisterHandler)
	talentGroup.POST("/login", talent.LoginHandler)

	router.Run()
}
