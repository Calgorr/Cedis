package database_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Calgorr/Cedis/database"
)

func TestCache_SetAndGet(t *testing.T) {
	c := database.NewCache(0)

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
	c := database.NewCache(0)

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
	c := database.NewCache(0)

	c.Set("foo", "bar", time.Minute)

	// wait for cleanup goroutine to run
	time.Sleep(time.Minute + time.Millisecond*100)

	_, exists := c.Get("foo")

	if exists {
		t.Errorf("Expected 'foo' to not exist in cache after cleanup")
	}
}

func TestCache_ConcurrentAccess(t *testing.T) {
	c := database.NewCache(0)
	var wg sync.WaitGroup

	// set 1000 keys with a TTL of 100ms concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set(fmt.Sprint(i), "value", time.Millisecond*100)
		}(i)
	}

	// wait for all keys to be set
	wg.Wait()

	// access all keys concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, exists := c.Get(fmt.Sprint(i))
			if !exists {
				t.Errorf("Expected key %v to exist in cache", i)
			}
		}(i)
	}

	// wait for all keys to be accessed
	wg.Wait()

	// wait for all keys to expire
	time.Sleep(time.Millisecond * 200)

	// ensure all keys have been removed
	for i := 0; i < 1000; i++ {
		_, exists := c.Get(fmt.Sprint(i))
		if exists {
			t.Errorf("Expected key %v to not exist in cache", i)
		}
	}
}
