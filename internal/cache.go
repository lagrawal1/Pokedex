package pokecache

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
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache

}

func (c *Cache) Add(key string, new_val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		c.entries[key] = cacheEntry{time.Now(), new_val}
	} else {
		entry.val = new_val
		entry.createdAt = time.Now()
		c.entries[key] = entry

	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true

}
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.entries {
			if time.Since(value.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
