package services

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

type TokenService struct{}

func NewTokenService() TokenService {
	return TokenService{}
}

func (s *TokenService) Generate() string {
	random := uuid.NewString()
	hash := sha256.Sum256([]byte(random))

	return fmt.Sprintf("%x", hash)
}
