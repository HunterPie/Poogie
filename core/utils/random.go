package utils

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

func NewRandomString() string {
	random := uuid.NewString()
	hash := sha256.Sum256([]byte(random))

	return fmt.Sprintf("%x", hash)
}
