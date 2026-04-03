package service

import "log/slog"

type Service struct {
	logger *slog.Logger
}

func New(l *slog.Logger) *Service {
	return &Service{
		logger: l,
	}
}
