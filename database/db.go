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

func (c *Cache) Set(key string, v interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()
	var expirationDate time.Time
	if ttl.Milliseconds() > 0 {
		expirationDate = time.Now().Add(ttl)
	}
	c.data[key] = value{
		value:          v,
		expirationDate: expirationDate,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if !v.expirationDate.IsZero() && time.Now().Compare(v.expirationDate) == 1 {
		return nil, false
	}
	return v.value, true
}
