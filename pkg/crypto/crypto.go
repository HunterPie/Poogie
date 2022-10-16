package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/Haato3o/poogie/core/crypto"
)

var (
	ErrInvalidContent        = errors.New("invalid content data to be decrypted")
	ErrInvalidEncryptedValue = errors.New("failed to decrypt value")
)

type AESCryptoService struct {
	key  []byte
	salt []byte
}

func NewCryptoService(key, salt string) crypto.ICryptographyService {
	return &AESCryptoService{[]byte(key), []byte(salt)}
}

func (s *AESCryptoService) Encrypt(content string) string {
	bytes := []byte(content)
	block, _ := aes.NewCipher(s.key)

	gcm, _ := cipher.NewGCM(block)
	encrypted := gcm.Seal(nil, s.salt, bytes, nil)

	return fmt.Sprintf("%x", encrypted)
}

func (s *AESCryptoService) Decrypt(content string) (string, error) {
	bytes, err := hex.DecodeString(content)

	if err != nil {
		return "", ErrInvalidContent
	}

	block, _ := aes.NewCipher(s.key)

	gcm, _ := cipher.NewGCM(block)
	decrypted, err := gcm.Open(nil, s.salt, bytes, nil)

	if err != nil {
		return "", ErrInvalidEncryptedValue
	}

	return string(decrypted), nil
}
