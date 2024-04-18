package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	SignupService(context.Context, SignupRequest) (User, error)
	LoginService(context.Context, LoginRequest) (string, error)
}

func NewService(r Repository, ts string) Service {
	return &service{
		repository:     r,
		tokenSignature: ts,
	}
}

type service struct {
	repository     Repository
	tokenSignature string
}

func (s *service) SignupService(ctx context.Context, su SignupRequest) (User, error) {
	pwdHash := sha256.Sum256([]byte(su.Password))
	u := User{
		Name:     su.Name,
		Email:    su.Email,
		Password: hex.EncodeToString(pwdHash[:]),
	}
	id, err := s.repository.RegisterUser(ctx, u)
	fmt.Println(err)
	if err != nil {
		return User{}, err
	}

	u.ID, u.Password = id, ""
	return u, nil
}

func (s service) LoginService(ctx context.Context, l LoginRequest) (string, error) {
	us, err := s.repository.FindUser(ctx, l.Email)

	if err != nil {
		return "", err
	}

	pwdHash := sha256.Sum256([]byte(l.Password))
	if hex.EncodeToString(pwdHash[:]) != us.Password {
		return "", err
	}
	token, err := generateToken(strconv.Itoa(us.ID), s.tokenSignature)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(id, tokenSignature string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	newToken, err := token.SignedString([]byte(tokenSignature))
	if err != nil {
		return "", err
	}

	return newToken, nil
}
