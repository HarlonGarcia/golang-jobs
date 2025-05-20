package handler

import (
	"net/http"

	"github.com/HarlonGarcia/golang-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		logger.Errorf("Error getting opening: %v", "id is required")
		sendError(ctx, http.StatusBadRequest, "opening id is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error getting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "an error occurred while getting the opening from database")
		return
	}

	sendSuccess(ctx, "get_opening", opening)
}
