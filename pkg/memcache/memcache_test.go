package memcache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Haato3o/poogie/pkg/memcache"
)

func TestMemoryCacheGet(t *testing.T) {
	var tests = []struct {
		InputKey string
		InputValue,
		Expected interface{}
		Exists bool
	}{
		{"test1", 1, 1, true},
		{"test2", nil, nil, false},
		{"test3", "hello world", "hello world", true},
	}

	service := memcache.New(10 * time.Hour)

	for _, test := range tests {
		if !test.Exists {
			continue
		}

		service.Set(test.InputKey, test.InputValue)
	}

	for i, test := range tests {
		name := fmt.Sprintf("MemoryCacheGet %d", i)

		t.Run(name, func(t *testing.T) {
			actual, exists := service.Get(test.InputKey)

			if actual != test.Expected {
				t.Errorf("got %s, expected %s", actual, test.Expected)
			}

			if exists != test.Exists {
				t.Errorf("got %t, expected %t", exists, test.Exists)
			}
		})
	}
}
