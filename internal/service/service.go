package service

import (
	"log/slog"

	"ArifulProtik/UpChat/internal/ent"
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
