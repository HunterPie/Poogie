package jwt

import (
	"errors"
	"time"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/golang-jwt/jwt/v4"
)

var ErrInvalidToken = errors.New("token was invalid")

type JWTAuthService struct {
	jwtKey []byte
}

// Create implements auth.IAuthService
func (s *JWTAuthService) Create(userId string) (string, error) {
	claims := JWTUserClaims{
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(s.jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// IsValid implements auth.IAuthService
func (s *JWTAuthService) IsValid(token string) bool {
	claims := JWTUserClaims{}

	parsed, err := jwt.ParseWithClaims(token, &claims, func(tkn *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return false
	}

	return parsed.Valid
}

// Parse implements auth.IAuthService
func (s *JWTAuthService) Parse(token string) (auth.AuthPayload, error) {
	claims := JWTUserClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(tkn *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return auth.AuthPayload{}, err
	}

	return auth.AuthPayload{
		UserId: claims.UserId,
	}, nil
}

func New(key string) auth.IAuthService {
	return &JWTAuthService{
		jwtKey: []byte(key),
	}
}
