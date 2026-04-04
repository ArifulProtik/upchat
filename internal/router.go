package internal

import (
	"ArifulProtik/UpChat/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup, controller *controller.Controller) {
	r.GET("/health", controller.Health)
	// Public
	auth := r.Group("/auth")
	auth.POST("/signup", controller.SignUp)
	auth.POST("/login", controller.LogIn)
	auth.POST("/logout", controller.LogOut)

	// Protected
	protected := r.Group("/")
	protected.Use(controller.Middleware.IsAuth())
	protected.GET("/auth/get-session", controller.GetSession)

	// Docs
	r.GET("/docs/openapi.json", controller.OpenAPISpec)
	r.GET("/docs", controller.ScalarDocs)
}
