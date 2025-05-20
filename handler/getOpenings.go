package handler

import (
	"net/http"

	"github.com/HarlonGarcia/golang-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error getting all openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "an error occurred while getting all openings from database")
		return
	}

	sendSuccess(ctx, "get_openings", openings)
}
