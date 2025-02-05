package handler

import (
	"fmt"
	"net/http"

	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}
	opening := schemas.Opening{}
	//find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
	}

	//delete opening

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to delete opening with id %s", id))
		return
	}
	sendSuccess(ctx, "delete-Opening", opening)
}
