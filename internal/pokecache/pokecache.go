package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct { //is one saved item
	createdAt time.Time //when it was saved
	val       []byte    //stores the raw api response
}

type Cache struct {
	entries  map[string]cacheEntry //stores items by URL
	mu       sync.Mutex            //prevents two goroutines from changing map simultaneously
	interval time.Duration         //how long entries must remain cached
}

func NewCache(interval time.Duration) *Cache { //creates initialized cache
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop() //runs the cleanup loop in the background; without "go" NewCache would enter reapLoop() and never return the cache
	return cache
}
