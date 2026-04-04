package service

import (
	"ArifulProtik/UpChat/internal/data"
	"ArifulProtik/UpChat/internal/ent"
	"ArifulProtik/UpChat/internal/ent/session"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GenerateToken() (string, error) {
	tokenByte := 32
	b := make([]byte, tokenByte)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(b), nil
}

func (s *Service) HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *Service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *Service) VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *Service) Login(
	ctx context.Context,
	account *ent.Account,
	body data.LoginBody,
) (*data.LoginResponse, error) {
	token, err := s.GenerateToken()
	if err != nil {
		return nil, err
	}
	hashedToken := s.HashToken(token)

	_, err = s.db.Session.Create().
		SetAccountID(account.ID).
		SetToken(hashedToken).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &data.LoginResponse{
		Token: token,
		User:  account.Edges.User,
	}, nil
}

func (s *Service) DeleteSession(ctx context.Context, token string) error {
	hashedToken := s.HashToken(token)
	_, err := s.db.Session.Delete().Where(session.Token(hashedToken)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetSession(ctx context.Context, token string) (*ent.Session, error) {
	hashedToken := s.HashToken(token)
	return s.db.Session.Query().Where(session.Token(hashedToken)).WithOwner().Only(ctx)
}
