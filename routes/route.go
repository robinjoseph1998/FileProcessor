package routes

import (
	"FileProcessor/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/submit", controllers.FileReader)
}
