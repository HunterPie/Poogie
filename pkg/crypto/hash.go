package crypto

import (
	"crypto/sha256"
	"fmt"
)

type SHA256Service struct {
	salt string
}

func (s *SHA256Service) Hash(content string) string {
	checksum := sha256.Sum256([]byte(content + s.salt))

	return fmt.Sprintf("%x", checksum)
}
