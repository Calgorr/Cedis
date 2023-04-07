package database

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	data map[string]value
}

type value struct {
	value          interface{}
	expirationDate time.Time
}

func NewCache() *Cache {
	cache := &Cache{data: make(map[string]value)}
	go cache.startCleanup()
	return cache
}

func (c *Cache) startCleanup() {
	for {
		time.Sleep(time.Minute)
		c.Lock()
		now := time.Now()
		for k, v := range c.data {
			if v.expirationDate.IsZero() && now.Compare(v.expirationDate) == 1 {
				delete(c.data, k)
			}
		}
		c.Unlock()
	}
}
