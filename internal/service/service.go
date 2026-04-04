package service

import (
	"ArifulProtik/UpChat/internal/ent"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
	db     *ent.Client
}

func New(l *slog.Logger, db *ent.Client) *Service {
	return &Service{
		logger: l,
		db:     db,
	}
}
