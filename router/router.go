package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Create a new gin router
	router := gin.Default()

	// Routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")

}
