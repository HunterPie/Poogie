package crypto_test

import (
	"fmt"
	"testing"

	"github.com/Haato3o/poogie/pkg/crypto"
)

const (
	HASH_SALT = "test"
)

func TestSHA256ServiceHash(t *testing.T) {
	var tests = []struct {
		Input,
		Expected string
	}{
		{"hello world", "4ae2a1e5830c6ce1795ce87a4d287ae15f39784d1409312f74e1d3bb874050b7"},
	}

	service := crypto.NewHashService(HASH_SALT)

	for i, test := range tests {
		name := fmt.Sprintf("SHA256ServiceHash %d", i)

		t.Run(name, func(t *testing.T) {
			actual := service.Hash(test.Input)

			if actual != test.Expected {
				t.Errorf("got %s, expected %s", actual, test.Expected)
			}
		})

	}
}
