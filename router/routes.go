package router

import (
	docs "github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/docs"
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitHandler()
	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Routes
	v1 := router.Group(basePath)
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

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Docs

}
