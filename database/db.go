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
