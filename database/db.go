package database

import (
	"regexp"
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

func (c *Cache) Delete(key string) int {
	c.Lock()
	defer c.Unlock()
	if _, exists := c.data[key]; !exists {
		return 0
	}
	delete(c.data, key)
	return 1
}

func (c *Cache) KeysMatchesPatern(pattern string) ([]string, error) {
	var keys []string
	for key, _ := range c.data {
		ok, err := regexp.MatchString(pattern, key)
		if err != nil {
			return nil, err
		} else if ok {
			keys = append(keys, key)
		}
	}
	return keys, nil
}
