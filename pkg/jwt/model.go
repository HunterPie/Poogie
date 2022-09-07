package jwt

import "github.com/golang-jwt/jwt/v4"

type JWTUserClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
