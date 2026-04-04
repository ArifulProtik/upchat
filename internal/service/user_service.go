package service

import (
	"ArifulProtik/UpChat/internal/data"
	"ArifulProtik/UpChat/internal/ent"
	"ArifulProtik/UpChat/internal/ent/account"
	"ArifulProtik/UpChat/internal/ent/user"
	"context"
	"errors"
)

func (s *Service) CreateUser(
	ctx context.Context,
	body data.UserCreateBody,
) (*ent.User, error) {
	exists, _ := s.FIndUserByEmail(ctx, body.Email)
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
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	user, err := tx.User.Create().
		SetName(body.Name).
		SetEmail(body.Email).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.Account.Create().
		SetEmail(body.Email).
		SetPassword(password).
		SetUser(user).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
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
