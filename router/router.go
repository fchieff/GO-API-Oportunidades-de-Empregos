package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Create a new gin router
	router := gin.Default()

	// Routes

	InitializeRoutes(router)

	//Run the server
	router.Run(":8080")

}
