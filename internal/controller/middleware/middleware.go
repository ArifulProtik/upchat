package middleware

import (
	"log/slog"

	"ArifulProtik/UpChat/internal/service"
)

const UserIDKey = "user_id"

type Middleware struct {
	service *service.Service
	logger  *slog.Logger
}

func New(s *service.Service, l *slog.Logger) *Middleware {
	return &Middleware{
		service: s,
		logger:  l,
	}
}
