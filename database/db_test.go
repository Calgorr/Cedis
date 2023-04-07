package database_test

import (
	"testing"
	"time"

	"github.com/Calgorr/Cedis/database"
)

func TestCache_SetAndGet(t *testing.T) {
	c := database.NewCache()

	c.Set("foo", "bar", time.Minute)
	value, exists := c.Get("foo")

	if !exists {
		t.Errorf("Expected 'foo' to exist in cache")
	}

	if value != "bar" {
		t.Errorf("Expected value of 'foo' to be 'bar', but got '%v'", value)
	}

	_, exists = c.Get("nonexistent")

	if exists {
		t.Errorf("Expected 'nonexistent' to not exist in cache")
	}
}

func TestCache_SetWithExpiration(t *testing.T) {
	c := database.NewCache()

	c.Set("foo", "bar", time.Millisecond*100)
	_, exists := c.Get("foo")

	if !exists {
		t.Errorf("Expected 'foo' to exist in cache")
	}

	time.Sleep(time.Millisecond * 200)

	_, exists = c.Get("foo")

	if exists {
		t.Errorf("Expected 'foo' to not exist in cache")
	}
}

func TestCache_Cleanup(t *testing.T) {
	c := database.NewCache()

	c.Set("foo", "bar", time.Minute)

	// wait for cleanup goroutine to run
	time.Sleep(time.Minute + time.Millisecond*100)

	_, exists := c.Get("foo")

	if exists {
		t.Errorf("Expected 'foo' to not exist in cache after cleanup")
	}
}
