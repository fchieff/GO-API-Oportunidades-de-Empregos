package handler

import (
	"net/http"

	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error fetching openings")
		return
	}

	sendSuccess(ctx, "list-Openings", openings)
}
