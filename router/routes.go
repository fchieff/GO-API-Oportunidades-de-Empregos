package router

import (
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	// Routes
	v1 := router.Group("/api/v1")

	{
		// Show Opening
		v1.GET("/opening", handler.CreateOpeningHandler)

		// Post Opening
		v1.POST("/opening", handler.ShowOpeningHandler)

		// Delete Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// Update Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// Show Opening
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
