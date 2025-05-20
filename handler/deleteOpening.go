package handler

import (
	"fmt"
	"net/http"

	"github.com/HarlonGarcia/golang-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		logger.Errorf("Error getting opening: %v", "id is required")
		sendError(ctx, http.StatusBadRequest, "opening id is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("opening with id %v was not found", id))
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		logger.Errorf("Error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "an error occurred while deleting the opening on database")
		return
	}

	sendSuccess(ctx, "delete_opening", opening)
}
