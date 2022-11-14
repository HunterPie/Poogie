package crypto

import (
	"crypto/sha256"
	"fmt"

	"github.com/Haato3o/poogie/core/crypto"
)

type SHA256Service struct {
	salt string
}

func (s *SHA256Service) Checksum(content string) string {
	checksum := sha256.Sum256([]byte(content))

	return fmt.Sprintf("%x", checksum)
}

func NewHashService(salt string) crypto.IHashService {
	return &SHA256Service{salt}
}

func (s *SHA256Service) Hash(content string) string {
	checksum := sha256.Sum256([]byte(content + s.salt))

	return fmt.Sprintf("%x", checksum)
}
