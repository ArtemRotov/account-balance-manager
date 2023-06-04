package service

import (
	"context"
	"errors"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	userRepo repository.UserRepository
	hasher   PasswordHasher
	signKey  string
	tokenTTL time.Duration
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int64
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
		log.Errorf("AuthService.CreateUser - cannot create user %v", err)
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
		log.Errorf("AuthService.GenerateToken - cannot get user %v", err)
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
		log.Errorf("AuthService.GenerateToken: cannot sign token: %v", err)
		return "", ErrCannotSignToken
	}

	return tokenString, nil
}
