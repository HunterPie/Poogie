package crypto_test

import (
	"fmt"
	"testing"

	"github.com/Haato3o/poogie/pkg/crypto"
)

const (
	KEY  = "testtesttesttesttesttesttesttest"
	SALT = "testing_salt"
)

func TestAESCryptoServiceEncrypt(t *testing.T) {
	var tests = []struct {
		Input,
		Expected string
	}{
		{"This is a test string for testing purposes", "44dc3d9080e6ef13ceea7fcdd3aedf28076ce662767d34dd3b467c8ed722311ee00175a2bc06b6f8a6e3541eb78ea155b1619cadaeded7ffd78a"},
	}

	service := crypto.NewCryptoService(KEY, SALT)

	for i, test := range tests {
		name := fmt.Sprintf("AESCryptoServiceEncrypt %d", i)

		t.Run(name, func(t *testing.T) {
			actual := service.Encrypt(test.Input)

			if actual != test.Expected {
				t.Errorf("got %s, want %s", actual, test.Expected)
			}
		})
	}
}

func TestAESCryptoServiceDecrypt(t *testing.T) {
	var tests = []struct {
		Input,
		Expected string
		ExpectedError error
	}{
		{"44dc3d9080e6ef13ceea7fcdd3aedf28076ce662767d34dd3b467c8ed722311ee00175a2bc06b6f8a6e3541eb78ea155b1619cadaeded7ffd78a", "This is a test string for testing purposes", nil},
		{"123456", "", crypto.ErrInvalidEncryptedValue},
		{"invalid hex", "", crypto.ErrInvalidContent},
	}

	service := crypto.NewCryptoService(KEY, SALT)

	for i, test := range tests {
		name := fmt.Sprintf("AESCryptoServiceDecrypt %d", i)

		t.Run(name, func(t *testing.T) {
			actual, err := service.Decrypt(test.Input)

			if actual != test.Expected {
				t.Errorf("got %s, want %s", actual, test.Expected)
			}

			if err != test.ExpectedError {
				t.Errorf("got %s, want %s", err, test.ExpectedError)
			}
		})
	}
}
