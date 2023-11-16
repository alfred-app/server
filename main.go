package main

import (
	"alfred/database"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	db := database.InitDB()
	database.MigrateDB(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, world!")
	})

	router.Run()
}
