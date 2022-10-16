package crypto

import (
	"crypto/rand"
	"encoding/binary"

	"github.com/Haato3o/poogie/core/crypto"
)

type CryptoRandomServiceImpl struct {
}

// UInt64 implements crypto.ICryptoRandomService
func (*CryptoRandomServiceImpl) UInt64() (uint64, error) {
	var bytes [8]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}

func (s *CryptoRandomServiceImpl) Int64() (int64, error) {
	result, err := s.UInt64()
	return int64(result), err
}

func NewCryptoRandomService() crypto.ICryptoRandomService {
	return &CryptoRandomServiceImpl{}
}
