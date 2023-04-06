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
}

func (c *Cache) startCleanup() {

}
