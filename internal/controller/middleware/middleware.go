package middleware

import (
	"ArifulProtik/UpChat/internal/service"
	"log/slog"
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
