package service

import (
	"context"
	"errors"
	"fmt"

	"ArifulProtik/UpChat/internal/data"
	"ArifulProtik/UpChat/internal/ent"
	"ArifulProtik/UpChat/internal/ent/account"
	"ArifulProtik/UpChat/internal/ent/user"
)

func (s *Service) CreateUser(
	ctx context.Context,
	body data.UserCreateBody,
) (*ent.User, error) {
	exists, err := s.FIndUserByEmail(ctx, body.Email)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("user already exists with this email")
	}

	password, err := s.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}

	tx, err := s.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Ensure rollback safety
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			panic(v)
		}
	}()

	user, err := tx.User.Create().
		SetName(body.Name).
		SetEmail(body.Email).
		Save(ctx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("tx err: %w, rollback err: %w", err, rbErr)
		}
		return nil, err
	}

	_, err = tx.Account.Create().
		SetEmail(body.Email).
		SetPassword(password).
		SetUser(user).
		Save(ctx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("tx err: %w, rollback err: %w", err, rbErr)
		}
		return nil, err
	}

	// IMPORTANT: check commit error
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindAccountByEmail(ctx context.Context, email string) (*ent.Account, error) {
	return s.db.Account.Query().Where(account.Email(email)).WithUser().Only(ctx)
}

func (s *Service) FindAccountByID(ctx context.Context, id string) (*ent.Account, error) {
	return s.db.Account.Query().Where(account.ID(id)).WithUser().Only(ctx)
}

func (s *Service) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	return s.db.User.Query().Where(user.ID(id)).WithAccount().Only(ctx)
}

func (s *Service) FIndUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.db.User.Query().Where(user.Email(email)).WithAccount().Only(ctx)
}
