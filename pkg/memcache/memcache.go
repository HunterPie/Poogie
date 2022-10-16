package memcache

import (
	"time"

	"github.com/Haato3o/poogie/core/cache"
	gocache "github.com/patrickmn/go-cache"
)

const DEFAULT_DELETE_INTERVAL = 5 * time.Minute

type MemoryCache struct {
	cache *gocache.Cache
}

func New(timeout time.Duration) cache.ICache {
	cache := gocache.New(timeout, DEFAULT_DELETE_INTERVAL)
	return &MemoryCache{cache}
}

func (m *MemoryCache) Get(key string) (interface{}, bool) {
	value, exists := m.cache.Get(key)
	return value, exists
}

func (m *MemoryCache) Set(key string, value interface{}) {
	m.cache.Add(key, value, DEFAULT_DELETE_INTERVAL)
}
