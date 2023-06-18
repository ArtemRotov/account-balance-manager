package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	userRepo repository.UserRepository
	hasher   PasswordHasher
	signKey  string
	tokenTTL time.Duration
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

func NewAuthService(r repository.UserRepository, h PasswordHasher, signKey string, ttl time.Duration) *AuthService {
	return &AuthService{
		userRepo: r,
		hasher:   h,
		signKey:  signKey,
		tokenTTL: ttl,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (int, error) {
	u := &model.User{
		Username: username,
		Password: s.hasher.Hash(password),
	}

	id, err := s.userRepo.CreateUser(ctx, u)
	if err != nil {
		if errors.Is(err, repoerrors.ErrAlreadyExists) {
			return 0, ErrUserAlreadyExists
		}
		logrus.Errorf("AuthService.CreateUser - cannot create user %v", err)
		return 0, err
	}

	return id, nil
}

func (s *AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	u := &model.User{
		Username: username,
		Password: s.hasher.Hash(password),
	}

	user, err := s.userRepo.UserByUsernameAndPassword(ctx, u.Username, u.Password)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			return "", ErrUserNotFound
		}
		logrus.Errorf("AuthService.GenerateToken - cannot get user %v", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	// sign token
	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		logrus.Errorf("AuthService.GenerateToken: cannot sign token: %v", err)
		return "", ErrCannotSignToken
	}

	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signKey), nil
	})

	if err != nil {
		return 0, ErrCannotParseToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, ErrCannotParseToken
	}

	return claims.UserId, nil
}
