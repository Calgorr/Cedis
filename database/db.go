package database

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	id   int
	Data map[string]value
}

type value struct {
	value          string
	expirationDate time.Time
}

func NewCache(id int) *Cache {
	cache := &Cache{Data: make(map[string]value), id: id}
	go cache.startCleanup()
	return cache
}

func (c *Cache) startCleanup() {
	for {
		time.Sleep(time.Minute)
		c.Lock()
		now := time.Now()
		for k, v := range c.Data {
			if v.expirationDate.IsZero() && now.Compare(v.expirationDate) == 1 {
				delete(c.Data, k)
			}
		}
		c.Unlock()
	}
}

func (c *Cache) Set(key, v string, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()
	var expirationDate time.Time
	if ttl.Milliseconds() > 0 {
		expirationDate = time.Now().Add(ttl)
	}
	c.Data[key] = value{
		value:          v,
		expirationDate: expirationDate,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.Data[key]
	fmt.Println(key, c.Data[key])
	if !ok {
		return "", false
	}
	if !v.expirationDate.IsZero() && time.Now().Compare(v.expirationDate) == 1 {
		return "", false
	}
	return v.value, true
}

func (c *Cache) Delete(key string) int {
	c.Lock()
	defer c.Unlock()
	if _, exists := c.Data[key]; !exists {
		return 0
	}
	delete(c.Data, key)
	return 1
}

func (c *Cache) KeysMatchesPatern(pattern string) ([]string, error) {
	var keys []string
	for key := range c.Data {
		ok, err := regexp.MatchString(pattern, key)
		if err != nil {
			return nil, err
		} else if ok {
			keys = append(keys, key)
		}
	}
	return keys, nil
}
