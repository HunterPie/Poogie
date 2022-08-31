package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

type AESCryptoService struct {
	key  []byte
	salt []byte
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
		return "", errors.New("invalid content data to be decrypted")
	}

	block, _ := aes.NewCipher(s.key)

	gcm, _ := cipher.NewGCM(block)
	decrypted, err := gcm.Open(nil, s.salt, bytes, nil)

	if err != nil {
		return "", errors.New("failed to decrypt value")
	}

	return string(decrypted), nil
}
