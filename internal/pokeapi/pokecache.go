package pokeapi

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		make(map[string]cacheEntry),
		&sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{time.Now(), val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, found := c.entries[key]
	if !found {
		return []byte{}, found
	}

	return entry.val, found
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.entries {
		if time.Since(value.createdAt) >= interval {
			delete(c.entries, key)
		}
	}
}
