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
	cacheEntry map[string]cacheEntry
	mu         *sync.RWMutex
	interval   time.Duration
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheEntry: make(map[string]cacheEntry),
		mu:         &sync.RWMutex{},
		interval:   interval,
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}

	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for k := range c.cacheEntry {
			if time.Since(c.cacheEntry[k].createdAt) > interval {
				delete(c.cacheEntry, k)
			}
		}
		c.mu.Unlock()
	}
}
