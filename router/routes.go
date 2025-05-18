package router

import (
	"github.com/HarlonGarcia/golang-jobs/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.GetOpeningsHandler)
	}
}
