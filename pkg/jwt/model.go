package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUserClaims struct {
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	jwt.RegisteredClaims
}
