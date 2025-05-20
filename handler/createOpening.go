package handler

import (
	"net/http"

	"github.com/HarlonGarcia/golang-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())

		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Location: request.Location,
		Company:  request.Company,
		Salary:   request.Salary,
		IsRemote: *request.IsRemote,
		Link:     request.Link,
		Role:     request.Role,
	}

	if err := db.Create((&opening)).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "an error occurred while creating the opening on database")
		return
	}

	sendSuccess(ctx, "create_opening", opening)
}
