package router

import (
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitHandler()

	// Routes
	v1 := router.Group("/api/v1")
	{
		// Create Opening (POST)
		v1.POST("/opening", handler.CreateOpeningHandler)

		// Show Opening (GET)
		v1.GET("/opening", handler.ShowOpeningHandler)

		// Delete Opening (DELETE)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// Update Opening (PUT)
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// List Openings (GET)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
