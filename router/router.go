package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	initRoutes(r)

	// Access the server at http://localhost:8080/
	r.Run(":8080")
}
