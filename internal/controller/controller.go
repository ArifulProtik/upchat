package controller

import (
	"ArifulProtik/UpChat/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	logger  *slog.Logger
	service *service.Service
}

func New(l *slog.Logger, s *service.Service) *Controller {
	return &Controller{
		logger:  l,
		service: s,
	}
}

func (c *Controller) Health(ctx *gin.Context) {
	c.logger.Info("Health check")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
