package controller

import (
	"ArifulProtik/UpChat/internal/controller/middleware"
	"ArifulProtik/UpChat/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Controller struct {
	logger     *slog.Logger
	service    *service.Service
	Middleware *middleware.Middleware
}

func New(l *slog.Logger, s *service.Service) *Controller {
	return &Controller{
		logger:     l,
		service:    s,
		Middleware: middleware.New(s, l),
	}
}

type ErrorResponse struct {
	Status int `json:"status"`
	Error  any `json:"error"`
}

// Health checks API availability.
//
//	@Summary	Health check
//	@Tags		health
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Router		/health [get]
func (c *Controller) Health(ctx *gin.Context) {
	c.logger.Info("Health check")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Copy[T any](fromt any) *T {
	var to T
	if err := copier.Copy(&to, fromt); err != nil {
		return nil
	}
	return &to
}
