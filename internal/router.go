package internal

import (
	"ArifulProtik/UpChat/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup, controller *controller.Controller) {
	r.GET("/health", controller.Health)
}
