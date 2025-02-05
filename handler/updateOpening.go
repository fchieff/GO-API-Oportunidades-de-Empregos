package handler

import (
	"fmt"
	"net/http"

	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("failed to validate request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return

	}
	opening := schemas.Opening{}
	//find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	// update opening

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	//save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("failed to save opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to update opening with id %s", id))
		return
	}
	sendSuccess(ctx, "update-Opening", opening)
}
